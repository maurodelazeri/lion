package gemini

import (

	//"encoding/json"

	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/jpillora/backoff"
	utils "github.com/maurodelazeri/concurrency-map-slice"
	number "github.com/maurodelazeri/go-number"
	event "github.com/maurodelazeri/lion/events"
	"github.com/maurodelazeri/lion/marketdata"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	pbEvent "github.com/maurodelazeri/lion/protobuf/heraldsquareAPI"
	"github.com/maurodelazeri/lion/venues/config"
	"github.com/pquerna/ffjson/ffjson"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

// Close closes the underlying network connection without
// sending or waiting for a close frame.
func (r *Websocket) Close() {
	r.mu.Lock()
	if r.Conn != nil {
		r.Conn.Close()
	}
	r.isConnected = false
	r.mu.Unlock()
}

// WebsocketClient ...
func (r *Websocket) WebsocketClient() {
	if r.RecIntvlMin == 0 {
		r.RecIntvlMin = 2 * time.Second
	}

	if r.RecIntvlMax == 0 {
		r.RecIntvlMax = 30 * time.Second
	}

	if r.RecIntvlFactor == 0 {
		r.RecIntvlFactor = 1.5
	}

	if r.HandshakeTimeout == 0 {
		r.HandshakeTimeout = 2 * time.Second
	}

	r.dialer = websocket.DefaultDialer
	r.dialer.HandshakeTimeout = r.HandshakeTimeout

	// Start reading from the socket
	r.startReading()

	go func() {
		r.connect()
	}()

	// wait on first attempt
	time.Sleep(r.HandshakeTimeout)
}

// IsConnected returns the WebSocket connection state
func (r *Websocket) IsConnected() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.isConnected
}

// CloseAndRecconect will try to reconnect.
func (r *Websocket) closeAndRecconect() {
	eventID, _ := uuid.NewV4()
	eventData := event.CreateBaseEvent(eventID.String(), "closeAndRecconect", nil, time.Now().UTC().Format(time.RFC3339Nano), r.base.GetName(), true, 0, pbEvent.System_WINTER)
	event.PublishEvent(eventData, "events", int64(1), false)
	r.Close()
	go func() {
		r.connect()
	}()
}

// GetHTTPResponse returns the http response from the handshake.
// Useful when WebSocket handshake fails,
// so that callers can handle redirects, authentication, etc.
func (r *Websocket) GetHTTPResponse() *http.Response {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.httpResp
}

// GetDialError returns the last dialer error.
// nil on successful connection.
func (r *Websocket) GetDialError() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.dialErr
}

func (r *Websocket) connect() {

	bb := &backoff.Backoff{
		Min:    r.RecIntvlMin,
		Max:    r.RecIntvlMax,
		Factor: r.RecIntvlFactor,
		Jitter: true,
	}

	rand.Seed(time.Now().UTC().UnixNano())

	r.OrderBookMAP = make(map[string]map[float64]float64)
	r.pairsMapping = utils.NewConcurrentMap()
	r.OrderbookTimestamps = utils.NewConcurrentMap()

	venueArrayPairs := []string{}
	for _, sym := range r.subscribedPairs {
		r.base.LiveOrderBook.Set(sym, &pbAPI.Orderbook{})
		r.OrderBookMAP[sym+"bids"] = make(map[float64]float64)
		r.OrderBookMAP[sym+"asks"] = make(map[float64]float64)
		venueConf, ok := r.base.VenueConfig.Get(r.base.GetName())
		if ok {
			venueArrayPairs = append(venueArrayPairs, venueConf.(config.VenueConfig).Products[sym].VenueSymbolIdentifier)
			r.pairsMapping.Set(venueConf.(config.VenueConfig).Products[sym].VenueSymbolIdentifier, sym)
			r.OrderbookTimestamps.Set(r.base.GetName()+sym, time.Now())
		}

		// We have an idenvidual connection per product, so we can specify this here
	}

	for {
		nextItvl := bb.Duration()

		currencies := []string{}
		for _, x := range venueArrayPairs {
			if r.base.Streaming {
				currencies = append(currencies, x)
			} else {
				currencies = append(currencies, x)
			}
		}
		wsConn, httpResp, err := r.dialer.Dial(websocketURL+strings.Join(currencies, "/"), r.reqHeader)

		r.mu.Lock()
		r.Conn = wsConn
		r.dialErr = err
		r.isConnected = err == nil
		r.httpResp = httpResp
		r.mu.Unlock()

		if err == nil {
			if r.base.Verbose {
				logrus.Printf("Dial: connection was successfully established with %s\n", websocketURL)
			}
			if err != nil {
				logrus.Printf("Websocket subscription error: %s\n", err)
			}
			break
		} else {
			eventID, _ := uuid.NewV4()
			eventData := event.CreateBaseEvent(eventID.String(), "connect", nil, time.Now().UTC().Format(time.RFC3339Nano), r.base.GetName(), true, 0, pbEvent.System_WINTER)
			event.PublishEvent(eventData, "events", int64(1), false)
			if r.base.Verbose {
				logrus.Println(err)
				logrus.Println("Dial: will try again in", nextItvl, "seconds.")
			}
		}

		time.Sleep(nextItvl)
	}
}

// startReading is a helper method for getting a reader
// using NextReader and reading from that reader to a buffer.
// If the connection is closed an error is returned
// startReading initiates a websocket client
// https://github.com/json-iterator/go
func (r *Websocket) startReading() {
	go func() {
		for {
			select {
			default:
				if r.IsConnected() {
					err := errors.New("websocket: not connected")
					msgType, resp, err := r.Conn.ReadMessage()
					if err != nil {
						logrus.Error(r.base.Name, " problem to read: ", err, " - ", time.Now())
						r.closeAndRecconect()
						continue
					}
					switch msgType {
					case websocket.TextMessage:
						data := Message{}
						err = ffjson.Unmarshal(resp, &data)
						if err != nil {
							logrus.Error("Unmarshal ", err)
							continue
						}
						var wg sync.WaitGroup
						updated := false

						refBook, ok := r.base.LiveOrderBook.Get(r.product)
						if !ok {
							logrus.Info("product not found  ", r.product)
							continue
						}
						refLiveBook := refBook.(*pbAPI.Orderbook)

						switch data.Type {
						case "update":
							for _, values := range data.Events {
								switch values.Type {
								case "change":
									switch values.Reason {
									case "initial":
										if values.Side == "ask" {
											refLiveBook.Asks = append(refLiveBook.Asks, &pbAPI.Item{Price: number.FromString(values.Price).Float64(), Volume: number.FromString(values.Remaining).Float64()})
										} else {
											refLiveBook.Bids = append(refLiveBook.Bids, &pbAPI.Item{Price: number.FromString(values.Price).Float64(), Volume: number.FromString(values.Remaining).Float64()})
										}
										wg.Add(1)
										go func() {
											sort.Slice(refLiveBook.Bids, func(i, j int) bool {
												return refLiveBook.Bids[i].Price > refLiveBook.Bids[j].Price
											})
											wg.Done()
										}()

										wg.Add(1)
										go func() {
											sort.Slice(refLiveBook.Asks, func(i, j int) bool {
												return refLiveBook.Asks[i].Price < refLiveBook.Asks[j].Price
											})
											wg.Done()
										}()
										wg.Wait()

										if len(refLiveBook.Asks) > 20 && len(refLiveBook.Bids) > 20 {
											refLiveBook.Asks = refLiveBook.Asks[0:20]
											refLiveBook.Bids = refLiveBook.Bids[0:20]
										}

										wg.Add(1)
										go func() {
											for _, value := range refLiveBook.Asks {
												r.OrderBookMAP[r.product+"asks"][value.Price] = value.Volume
											}
											wg.Done()
										}()

										wg.Add(1)
										go func() {
											for _, value := range refLiveBook.Bids {
												r.OrderBookMAP[r.product+"bids"][value.Price] = value.Volume
											}
											wg.Done()
										}()
										wg.Wait()

									case "cancel":
										if values.Side == "ask" {
											if number.FromString(values.Remaining).Float64() == 0 {
												price := number.FromString(values.Price).Float64()
												if _, ok := r.OrderBookMAP[r.product+"asks"][price]; ok {
													delete(r.OrderBookMAP[r.product+"asks"], price)
													updated = true
												}
											} else {
												price := number.FromString(values.Price).Float64()
												if _, ok := r.OrderBookMAP[r.product+"asks"][price]; ok {
													r.OrderBookMAP[r.product+"asks"][price] = number.FromString(values.Remaining).Float64()
													updated = true
												}
											}
										} else {
											if number.FromString(values.Remaining).Float64() == 0 {
												price := number.FromString(values.Price).Float64()
												if _, ok := r.OrderBookMAP[r.product+"bids"][price]; ok {
													delete(r.OrderBookMAP[r.product+"bids"], price)
													updated = true
												}
											} else {
												price := number.FromString(values.Price).Float64()
												if _, ok := r.OrderBookMAP[r.product+"bids"][price]; ok {
													r.OrderBookMAP[r.product+"bids"][price] = number.FromString(values.Remaining).Float64()
													updated = true
												}
											}
										}
									case "place":
										if values.Side == "ask" {
											price := number.FromString(values.Price).Float64()
											amount := number.FromString(values.Remaining).Float64()
											totalLevels := len(refLiveBook.GetAsks())
											if totalLevels == r.base.MaxLevelsOrderBook {
												if price > refLiveBook.Asks[totalLevels-1].Price {
													continue
												}
											}
											updated = true
											r.OrderBookMAP[r.product+"asks"][price] = amount
										} else {
											price := number.FromString(values.Price).Float64()
											amount := number.FromString(values.Remaining).Float64()
											totalLevels := len(refLiveBook.GetBids())
											if totalLevels == r.base.MaxLevelsOrderBook {
												if price < refLiveBook.Bids[totalLevels-1].Price {
													continue
												}
											}
											updated = true
											r.OrderBookMAP[r.product+"bids"][price] = amount
										}
									}
								case "trade":
									var side string
									if values.Side == "bid" {
										side = "buy"
									} else {
										side = "sell"
									}
									trades := &pbAPI.Trade{
										Product:         r.product,
										Venue:           r.base.GetName(),
										VenueTradeId:    fmt.Sprintf("%d", values.Tid),
										SystemTimestamp: time.Now().UTC().Format(time.RFC3339Nano),
										VenueTimestamp:  time.Unix(0, data.Timestampms*int64(time.Millisecond)).UTC().Format(time.RFC3339Nano),
										Price:           number.FromString(values.Price).Float64(),
										OrderSide:       side,
										Volume:          number.FromString(values.Amount).Float64(),
									}

									serialized, err := proto.Marshal(trades)
									if err != nil {
										logrus.Error("Marshal ", err)
									}
									err = r.base.SocketClient.Publish("trades:"+r.base.GetName()+"."+r.product, serialized)
									if err != nil {
										logrus.Error("Socket sent ", err)
									}
									marketdata.PublishMarketData(serialized, "trades."+r.base.GetName()+"."+r.product, 1, false)
								}
							}
						}

						// we dont need to update the book if any level we care was changed
						if !updated {
							continue
						}

						wg.Add(1)
						go func() {
							refLiveBook.Bids = []*pbAPI.Item{}
							for price, amount := range r.OrderBookMAP[r.product+"bids"] {
								refLiveBook.Bids = append(refLiveBook.Bids, &pbAPI.Item{Price: price, Volume: amount})
							}
							sort.Slice(refLiveBook.Bids, func(i, j int) bool {
								return refLiveBook.Bids[i].Price > refLiveBook.Bids[j].Price
							})
							wg.Done()
						}()

						wg.Add(1)
						go func() {
							refLiveBook.Asks = []*pbAPI.Item{}
							for price, amount := range r.OrderBookMAP[r.product+"asks"] {
								refLiveBook.Asks = append(refLiveBook.Asks, &pbAPI.Item{Price: price, Volume: amount})
							}
							sort.Slice(refLiveBook.Asks, func(i, j int) bool {
								return refLiveBook.Asks[i].Price < refLiveBook.Asks[j].Price
							})
							wg.Done()
						}()

						wg.Wait()

						wg.Add(1)
						go func() {
							totalBids := len(refLiveBook.Bids)
							if totalBids > r.base.MaxLevelsOrderBook {
								refLiveBook.Bids = refLiveBook.Bids[0:r.base.MaxLevelsOrderBook]
							}
							wg.Done()
						}()
						wg.Add(1)
						go func() {
							totalAsks := len(refLiveBook.Asks)
							if totalAsks > r.base.MaxLevelsOrderBook {
								refLiveBook.Asks = refLiveBook.Asks[0:r.base.MaxLevelsOrderBook]
							}
							wg.Done()
						}()

						wg.Wait()
						book := &pbAPI.Orderbook{
							Product:         r.product,
							Venue:           r.base.GetName(),
							Levels:          int64(r.base.MaxLevelsOrderBook),
							SystemTimestamp: time.Now().UTC().Format(time.RFC3339Nano),
							VenueTimestamp:  time.Unix(0, data.Timestampms*int64(time.Millisecond)).UTC().Format(time.RFC3339Nano),
							Asks:            refLiveBook.Asks,
							Bids:            refLiveBook.Bids,
						}

						r.base.LiveOrderBook.Set(r.product, book)

						serialized, err := proto.Marshal(book)
						if err != nil {
							logrus.Error("Marshal ", err)
						}
						err = r.base.SocketClient.Publish("orderbooks:"+r.base.GetName()+"."+r.product, serialized)
						if err != nil {
							logrus.Error("Socket sent ", err)
						}
						// publish orderbook within a timeframe at least 1 second
						value, exist := r.OrderbookTimestamps.Get(book.GetVenue() + book.GetProduct())
						if !exist {
							continue
						}
						elapsed := time.Since(value.(time.Time))
						if elapsed.Seconds() <= 1 {
							continue
						}
						r.OrderbookTimestamps.Set(book.GetVenue()+book.GetProduct(), time.Now())
						marketdata.PublishMarketData(serialized, "orderbooks."+r.base.GetName()+"."+r.product, 1, false)
					}
				}
			}
		}
	}()
}
