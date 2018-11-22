package gemini

import (

	//"encoding/json"

	"errors"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/jpillora/backoff"
	"github.com/maurodelazeri/concurrency-map-slice"
	number "github.com/maurodelazeri/go-number"
	"github.com/maurodelazeri/lion/common"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	"github.com/maurodelazeri/lion/streaming/kafka/producer"
	"github.com/maurodelazeri/lion/venues/config"
	"github.com/pquerna/ffjson/ffjson"
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
	r.base.LiveOrderBook = utils.NewConcurrentMap()
	r.pairsMapping = utils.NewConcurrentMap()

	venueArrayPairs := []string{}
	for _, sym := range r.subscribedPairs {
		r.base.LiveOrderBook.Set(sym, &pbAPI.Orderbook{})
		r.OrderBookMAP[sym+"bids"] = make(map[float64]float64)
		r.OrderBookMAP[sym+"asks"] = make(map[float64]float64)
		venueConf, ok := r.base.VenueConfig.Get(r.base.GetName())
		if ok {
			venueArrayPairs = append(venueArrayPairs, venueConf.(config.VenueConfig).Products[sym].VenueName)
			r.product = venueConf.(config.VenueConfig).Products[sym].VenueName

			r.pairsMapping.Set(venueConf.(config.VenueConfig).Products[sym].VenueName, sym)
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
			//if r.base.Verbose {
			logrus.Printf("Dial: connection was successfully established with %s\n", websocketURL+strings.Join(currencies, "/"))
			//	}
			break
		} else {
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

						logrus.Info(string(resp))

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
									logrus.Info(string(resp))

									var side pbAPI.Side
									if values.Side == "bid" {
										side = pbAPI.Side_BUY
									} else {
										side = pbAPI.Side_SELL
									}
									refBook, ok := r.base.LiveOrderBook.Get(r.product)
									if !ok {
										continue
									}
									refLiveBook := refBook.(*pbAPI.Orderbook)
									trades := &pbAPI.Trade{
										Product:   pbAPI.Product((pbAPI.Product_value[r.product])),
										Venue:     pbAPI.Venue((pbAPI.Venue_value[r.base.GetName()])),
										Timestamp: common.MakeTimestamp(),
										Price:     number.FromString(values.Price).Float64(),
										OrderSide: side,
										Volume:    number.FromString(values.Amount).Float64(),
										VenueType: pbAPI.VenueType_SPOT,
										Asks:      refLiveBook.Asks,
										Bids:      refLiveBook.Bids,
									}

									logrus.Warn("NEW TRADe ", trades)

									serialized, err := proto.Marshal(trades)
									if err != nil {
										log.Fatal("proto.Marshal error: ", err)
									}
									r.MessageType[0] = 0
									serialized = append(r.MessageType, serialized[:]...)
									kafkaproducer.PublishMessageAsync(r.product+"."+r.base.Name+".trade", serialized, 1, false)
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
							Product:   pbAPI.Product((pbAPI.Product_value[r.product])),
							Venue:     pbAPI.Venue((pbAPI.Venue_value[r.base.GetName()])),
							Levels:    int32(r.base.MaxLevelsOrderBook),
							Timestamp: common.MakeTimestamp(),
							Asks:      refLiveBook.Asks,
							Bids:      refLiveBook.Bids,
							VenueType: pbAPI.VenueType_SPOT,
						}
						refLiveBook = book

						if r.base.Streaming {
							serialized, err := proto.Marshal(book)
							if err != nil {
								log.Fatal("proto.Marshal error: ", err)
							}
							r.MessageType[0] = 1
							serialized = append(r.MessageType, serialized[:]...)
							kafkaproducer.PublishMessageAsync(r.product+"."+r.base.Name+".orderbook", serialized, 1, false)
						}
					}
				}
			}
		}
	}()
}
