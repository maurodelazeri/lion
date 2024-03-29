package binance

import (

	//"encoding/json"

	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/jpillora/backoff"
	utils "github.com/maurodelazeri/concurrency-map-slice"
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

	r.FetchEnabledOrderBooks()

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

// FetchEnabledOrderBooks gets the initial orderbook
func (r *Websocket) FetchEnabledOrderBooks() {
	logrus.Info("Feching orderbooks")
	go func() {
		for _, sym := range r.subscribedPairs {
			venueConf, _ := r.base.VenueConfig.Get(r.base.GetName())
			resp := OrderBook{}
			params := url.Values{}
			params.Set("symbol", venueConf.(config.VenueConfig).Products[sym].VenueSymbolIdentifier)
			params.Set("limit", "50")
			path := fmt.Sprintf("%s%s?%s", mainWebsite, orderBookDepth, params.Encode())
			if err := r.base.SendHTTPRequest(path, &resp); err != nil {
				logrus.Error("problem to fech the orderbook ", venueConf.(config.VenueConfig).Products[sym].VenueSymbolIdentifier, " ", err)
				continue
			}
			for _, book := range resp.Bids {
				amount, _ := strconv.ParseFloat(book[1].(string), 64)
				price, _ := strconv.ParseFloat(book[0].(string), 64)
				r.OrderBookMAP[sym+"bids"][price] = amount
			}
			for _, book := range resp.Asks {
				amount, _ := strconv.ParseFloat(book[1].(string), 64)
				price, _ := strconv.ParseFloat(book[0].(string), 64)
				r.OrderBookMAP[sym+"asks"][price] = amount
			}
			r.LockTillBookFetchToFinish[venueConf.(config.VenueConfig).Products[sym].VenueSymbolIdentifier+"@depth"] = sym
		}
	}()
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
	}

	for {
		nextItvl := bb.Duration()

		currencies := []string{}
		for _, x := range venueArrayPairs {
			currencies = append(currencies, strings.ToLower(x)+"@trade")
			if r.base.Streaming {
				currencies = append(currencies, strings.ToLower(x)+"@depth")
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
						logrus.Error(r.base.Name, " problem to read: ", err)
						r.closeAndRecconect()
						continue
					}
					switch msgType {
					case websocket.TextMessage:
						data := Message{}
						err = ffjson.Unmarshal(resp, &data)
						if err != nil {
							logrus.Error(err)
							continue
						}
						if strings.Contains(data.Stream, "@depth") {
							var wg sync.WaitGroup
							message := MessageDepht{}
							err = ffjson.Unmarshal(resp, &message)
							if err != nil {
								logrus.Error(err)
								continue
							}
							value, exist := r.pairsMapping.Get(message.Data.Symbol)
							if !exist {
								continue
							}
							product := value.(string)
							if _, ok := r.LockTillBookFetchToFinish[message.Data.Symbol+"@depth"]; ok {

								updated := false
								refBook, ok := r.base.LiveOrderBook.Get(product)
								if !ok {
									continue
								}
								refLiveBook := refBook.(*pbAPI.Orderbook)

								wg.Add(1)
								go func() {
									for _, book := range message.Data.Bids {
										amount, _ := strconv.ParseFloat(book[1].(string), 64)
										price, _ := strconv.ParseFloat(book[0].(string), 64)
										if amount == 0 {
											if _, ok := r.OrderBookMAP[product+"bids"][price]; ok {
												delete(r.OrderBookMAP[product+"bids"], price)
												updated = true
											}
										} else {
											amount, _ := strconv.ParseFloat(book[1].(string), 64)
											price, _ := strconv.ParseFloat(book[0].(string), 64)
											totalLevels := len(refLiveBook.GetBids())
											if totalLevels == r.base.MaxLevelsOrderBook {
												if price < refLiveBook.Bids[totalLevels-1].Price {
													continue
												}
											}
											updated = true
											r.OrderBookMAP[product+"bids"][price] = amount
										}
									}
									wg.Done()
								}()

								wg.Add(1)
								go func() {
									for _, book := range message.Data.Asks {
										amount, _ := strconv.ParseFloat(book[1].(string), 64)
										price, _ := strconv.ParseFloat(book[0].(string), 64)
										if amount == 0 {
											if _, ok := r.OrderBookMAP[product+"asks"][price]; ok {
												delete(r.OrderBookMAP[product+"asks"], price)
												updated = true
											}
										} else {
											amount, _ := strconv.ParseFloat(book[1].(string), 64)
											price, _ := strconv.ParseFloat(book[0].(string), 64)
											totalLevels := len(refLiveBook.GetAsks())
											if totalLevels == r.base.MaxLevelsOrderBook {
												if price > refLiveBook.Asks[totalLevels-1].Price {
													continue
												}
											}
											updated = true
											r.OrderBookMAP[product+"asks"][price] = amount
										}
									}
									wg.Done()
								}()

								wg.Wait()
								// we dont need to update the book if any level we care was changed
								if !updated {
									continue
								}

								wg.Add(1)
								go func() {
									refLiveBook.Bids = []*pbAPI.Item{}
									for price, amount := range r.OrderBookMAP[product+"bids"] {
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
									for price, amount := range r.OrderBookMAP[product+"asks"] {
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
									Product:         product,
									Venue:           r.base.GetName(),
									SystemTimestamp: time.Now().UTC().Format(time.RFC3339Nano),
									VenueTimestamp:  time.Unix(0, message.Data.EventTime*int64(time.Millisecond)).UTC().Format(time.RFC3339Nano),
									Asks:            refLiveBook.Asks,
									Bids:            refLiveBook.Bids,
								}
								r.base.LiveOrderBook.Set(product, book)

								serialized, err := proto.Marshal(book)
								if err != nil {
									logrus.Error("Marshal ", err)
								}
								err = r.base.SocketClient.Publish("orderbooks:"+r.base.GetName()+"."+product, serialized)
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
								marketdata.PublishMarketData(serialized, "orderbooks."+r.base.GetName()+"."+product, 1, false)
							}
						} else {
							message := MessageTrade{}
							err = ffjson.Unmarshal(resp, &message)
							if err != nil {
								logrus.Error(err)
								continue
							}
							value, exist := r.pairsMapping.Get(message.Data.Symbol)
							if !exist {
								continue
							}
							product := value.(string)

							var side string
							if message.Data.Buyer {
								side = "buy"
							} else {
								side = "sell"
							}
							logrus.Info(message.Data.EventTime)
							trades := &pbAPI.Trade{
								Product:         product,
								VenueTradeId:    strconv.FormatInt(message.Data.TradeID, 10),
								Venue:           r.base.GetName(),
								SystemTimestamp: time.Now().UTC().Format(time.RFC3339Nano),
								VenueTimestamp:  time.Unix(0, message.Data.EventTime*int64(time.Millisecond)).UTC().Format(time.RFC3339Nano),
								Price:           message.Data.Price,
								OrderSide:       side,
								Volume:          message.Data.Quantity,
							}
							serialized, err := proto.Marshal(trades)
							if err != nil {
								logrus.Error("Marshal ", err)
							}
							err = r.base.SocketClient.Publish("trades:"+r.base.GetName()+"."+product, serialized)
							if err != nil {
								logrus.Error("Socket sent ", err)
							}
							marketdata.PublishMarketData(serialized, "trades."+r.base.GetName()+"."+product, 1, false)
						}
					}
				}
			}
		}
	}()
}
