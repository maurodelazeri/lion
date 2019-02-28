package kraken

import (

	//"encoding/json"

	"errors"
	"math"
	"math/rand"
	"net/http"
	"reflect"
	"sort"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/jpillora/backoff"
	utils "github.com/maurodelazeri/concurrency-map-slice"
	"github.com/maurodelazeri/go-number"
	"github.com/maurodelazeri/lion/common"
	"github.com/maurodelazeri/lion/marketdata"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	"github.com/maurodelazeri/lion/venues/config"
	"github.com/sirupsen/logrus"
)

// Subscribe subsribe public and private endpoints
func (r *Websocket) Subscribe(products []string) error {
	subscribeTrades := SubscriptionTrade{
		Event: "subscribe",
		Pair:  products,
	}
	subscribeTrades.Subscription.Name = "trade"
	json, err := common.JSONEncode(subscribeTrades)
	if err != nil {
		return err
	}
	err = r.Conn.WriteMessage(websocket.TextMessage, json)
	if err != nil {
		return err
	}

	subscribeBooks := SubscriptionBook{
		Event: "subscribe",
		Pair:  products,
	}
	subscribeBooks.Subscription.Name = "book"
	subscribeBooks.Subscription.Depth = 25
	json, err = common.JSONEncode(subscribeBooks)
	if err != nil {
		return err
	}
	err = r.Conn.WriteMessage(websocket.TextMessage, json)
	if err != nil {
		return err
	}
	return nil
}

// Heartbeat ...
func (r *Websocket) Heartbeat() {
	go func() {
		for {
			if r.IsConnected() {
				json, err := common.JSONEncode(PingPong{Event: "ping", Reqid: time.Now().Unix()})
				if err != nil {
					logrus.Error("Subscription ", err)
					continue
				}
				err = r.Conn.WriteMessage(websocket.TextMessage, json)
				if err != nil {
					logrus.Error("Subscription ", err)
					continue
				}
			}
			time.Sleep(time.Second * 60)
		}
	}()
}

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
	r.pairsMapping = utils.NewConcurrentMap()
	r.OrderbookTimestamps = utils.NewConcurrentMap()
	r.StreamingChannels = make(map[int64]string)

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

		wsConn, httpResp, err := r.dialer.Dial(websocketURL, r.reqHeader)

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
			err = r.Subscribe(venueArrayPairs)
			if err != nil {
				logrus.Printf("Websocket subscription error: %s\n", err)
			}
			break
		} else {
			logrus.Println(err)
			logrus.Println("Dial: will try again in", nextItvl, "seconds.")
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

						var result interface{}
						err = common.JSONDecode(resp, &result)
						if err != nil {
							logrus.Error("JSONDecode ", err)
							continue
						}
						switch reflect.TypeOf(result).String() {
						case "map[string]interface {}":
							data := result.(map[string]interface{})
							if _, ok := data["channelID"]; ok {
								r.StreamingChannels[int64(data["channelID"].(float64))] = data["pair"].(string)
							}
						case "[]interface {}":
							data := result.([]interface{})
							switch reflect.TypeOf(data[1]).String() {
							case "map[string]interface {}":
								if val, ok := r.StreamingChannels[int64(data[0].(float64))]; ok {
									value, exist := r.pairsMapping.Get(val)
									if !exist {
										logrus.Warn("Product does not exist ", val)
										continue
									}
									product := value.(string)
									refBook, ok := r.base.LiveOrderBook.Get(product)
									if !ok {
										continue
									}
									refLiveBook := refBook.(*pbAPI.Orderbook)
									var wg sync.WaitGroup
									updated := false

									subData := data[1].(interface{})
									data := subData.(map[string]interface{})

									if val, ok := data["as"]; ok {
										asks := val.([]interface{})
										wg.Add(1)
										go func(arr []interface{}) {
											total := 0
											for _, line := range arr {
												row := line.([]interface{})
												price := number.FromString(row[0].(string)).Float64()
												amount := number.FromString(row[1].(string)).Float64()
												refLiveBook.Asks = append(refLiveBook.Asks, &pbAPI.Item{Price: price, Volume: amount})
												r.OrderBookMAP[product+"asks"][price] = amount
												total++
											}
											wg.Done()
										}(asks)
									}

									if val, ok := data["bs"]; ok {
										bids := val.([]interface{})
										wg.Add(1)
										go func(arr []interface{}) {
											total := 0
											for _, line := range arr {
												row := line.([]interface{})
												price := number.FromString(row[0].(string)).Float64()
												amount := number.FromString(row[1].(string)).Float64()
												refLiveBook.Bids = append(refLiveBook.Bids, &pbAPI.Item{Price: price, Volume: amount})
												r.OrderBookMAP[product+"bids"][price] = amount
												total++
											}
											wg.Done()
										}(bids)
									}

									if val, ok := data["a"]; ok {
										asks := val.([]interface{})
										wg.Add(1)
										go func(arr []interface{}) {
											for _, line := range arr {
												row := line.([]interface{})
												price := number.FromString(row[0].(string)).Float64()
												amount := number.FromString(row[1].(string)).Float64()
												totalLevels := len(refLiveBook.GetAsks())
												if amount == 0 {
													if _, ok := r.OrderBookMAP[product+"asks"][price]; ok {
														delete(r.OrderBookMAP[product+"asks"], price)
														updated = true
													}
												} else {
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
										}(asks)
									}

									if val, ok := data["b"]; ok {
										bids := val.([]interface{})
										wg.Add(1)
										go func(arr []interface{}) {
											for _, line := range arr {
												row := line.([]interface{})
												price := number.FromString(row[0].(string)).Float64()
												amount := number.FromString(row[1].(string)).Float64()
												if amount == 0 {
													if _, ok := r.OrderBookMAP[product+"bids"][price]; ok {
														delete(r.OrderBookMAP[product+"bids"], price)
														updated = true
													}
												} else {
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
										}(bids)
									}

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
										Levels:          int64(r.base.MaxLevelsOrderBook),
										SystemTimestamp: time.Now().UTC().Format(time.RFC3339Nano),
										//VenueTimestamp:  dateTimeRef.UTC().Format(time.RFC3339Nano),
										Asks: refLiveBook.Asks,
										Bids: refLiveBook.Bids,
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
									value, exist = r.OrderbookTimestamps.Get(book.GetVenue() + book.GetProduct())
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
							case "[]interface {}":
								if val, ok := r.StreamingChannels[int64(data[0].(float64))]; ok {
									value, exist := r.pairsMapping.Get(val)
									if !exist {
										logrus.Warn("Product does not exist ", val)
										continue
									}
									product := value.(string)
									subData := data[1].([]interface{})
									for _, tradeData := range subData {
										data := tradeData.([]interface{})
										var side string
										if data[3].(string) == "b" {
											side = "buy"
										} else {
											side = "sell"
										}
										sec, dec := math.Modf(number.FromString(data[2].(string)).Float64())
										trades := &pbAPI.Trade{
											Product:         product,
											VenueTradeId:    data[4].(string),
											Venue:           r.base.GetName(),
											SystemTimestamp: time.Now().UTC().Format(time.RFC3339Nano),
											VenueTimestamp:  time.Unix(int64(sec), int64(dec*(1e9))).UTC().Format(time.RFC3339Nano),
											Price:           number.FromString(data[0].(string)).Float64(),
											OrderSide:       side,
											Volume:          number.FromString(data[1].(string)).Float64(),
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
							default:
								logrus.Warn("Datatype not found ", data)
							}
						default:
							logrus.Warn("Datatype not found")
						}
					}
				}
			}
		}
	}()
}
