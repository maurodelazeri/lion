package bitfinex

import (

	//"encoding/json"

	"errors"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/jpillora/backoff"
	"github.com/maurodelazeri/concurrency-map-slice"
	"github.com/maurodelazeri/lion/common"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	"github.com/maurodelazeri/lion/streaming/kafka/producer"
	"github.com/maurodelazeri/lion/venues/config"
	"github.com/sirupsen/logrus"
)

// MessageChannel ...
type MessageChannel struct {
	Event   string `json:"event,omitempty"`
	Channel string `json:"channel,omitempty"`
	Symbol  string `json:"symbol,omitempty"`
	Prec    string `json:"prec,omitempty"`
	Freq    string `json:"freq,omitempty"`
	Len     int    `json:"len,omitempty"`
}

// Subscribe subsribe public and private endpoints
func (r *Websocket) Subscribe(products []string) error {
	subscribe := []MessageChannel{}
	if r.base.Streaming {
		for _, product := range products {
			book := MessageChannel{
				Event:   "subscribe",
				Channel: "book",
				Symbol:  "t" + product,
				Prec:    "P0",
				Freq:    "F0",
				Len:     25,
			}
			subscribe = append(subscribe, book)
			trade := MessageChannel{
				Event:   "subscribe",
				Channel: "trades",
				Symbol:  "t" + product,
			}
			subscribe = append(subscribe, trade)
		}
	} else {
		for _, product := range products {
			trade := MessageChannel{
				Event:   "subscribe",
				Channel: "trades",
				Symbol:  "t" + product,
			}
			subscribe = append(subscribe, trade)
		}
	}
	for _, channels := range subscribe {
		json, err := common.JSONEncode(channels)
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
	return nil
}

// WebsocketAddSubscriptionChannel adds a new subscription channel to the
// WebsocketSubdChannels map in bitfinex.go (Bitfinex struct)
func (r *Websocket) WebsocketAddSubscriptionChannel(chanID int, channel, pair string) {
	chanInfo := WebsocketChanInfo{Pair: pair, Channel: channel}
	r.base.WebsocketSubdChannels[chanID] = chanInfo
	if r.base.Verbose {
		log.Printf("Subscribed to Channel: %s Pair: %s ChannelID: %d\n", channel, pair, chanID)
	}
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

	r.setReadTimeOut(1)

	bb := &backoff.Backoff{
		Min:    r.RecIntvlMin,
		Max:    r.RecIntvlMax,
		Factor: r.RecIntvlFactor,
		Jitter: true,
	}

	rand.Seed(time.Now().UTC().UnixNano())

	r.OrderBookMAP = make(map[string]map[float64]float64)
	r.pairsMapping = utils.NewConcurrentMap()

	venueArrayPairs := []string{}
	for _, sym := range r.subscribedPairs {
		r.base.LiveOrderBook.Set(sym, &pbAPI.Orderbook{})
		r.OrderBookMAP[sym+"bids"] = make(map[float64]float64)
		r.OrderBookMAP[sym+"asks"] = make(map[float64]float64)
		venueConf, ok := r.base.VenueConfig.Get(r.base.GetName())
		if ok {
			venueArrayPairs = append(venueArrayPairs, venueConf.(config.VenueConfig).Products[sym].VenueName)
			r.pairsMapping.Set(venueConf.(config.VenueConfig).Products[sym].VenueName, sym)
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

						var result interface{}
						err := common.JSONDecode(resp, &result)
						if err != nil {
							log.Println("Ops ", err)
							continue
						}
						switch reflect.TypeOf(result).String() {
						case "map[string]interface {}":
							eventData := result.(map[string]interface{})
							event := eventData["event"]

							switch event {
							case "subscribed":
								r.WebsocketAddSubscriptionChannel(int(eventData["chanId"].(float64)), eventData["channel"].(string), eventData["pair"].(string))
							case "auth":
								status := eventData["status"].(string)

								if status == "OK" {
									r.WebsocketAddSubscriptionChannel(0, "account", "N/A")
								} else if status == "fail" {
									log.Printf("%s Websocket unable to AUTH", eventData["code"].(string))
								}
							}
						case "[]interface {}":
							chanData := result.([]interface{})
							chanID := int(chanData[0].(float64))
							chanInfo, ok := r.base.WebsocketSubdChannels[chanID]
							if !ok {
								log.Printf("Unable to locate chanID: %d\n", chanID)
							}
							if len(chanData) == 2 {
								if reflect.TypeOf(chanData[1]).String() == "string" {
									if chanData[1].(string) == "hb" {
										continue
									}
								}
							}

							value, exist := r.pairsMapping.Get(strings.ToUpper(chanInfo.Pair))
							if !exist {
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

							switch chanInfo.Channel {
							case "unknown":
								logrus.Error(r.base.GetName(), " Problem subscribing channel, wrong name")
							case "book":
								switch len(chanData) {
								case 2:
									data := chanData[1].([]interface{})
									// Snapshoot
									if len(data) > 10 {
										for _, x := range data {
											y := x.([]interface{})
											if y[2].(float64) > 0 {
												refLiveBook.Bids = append(refLiveBook.Bids, &pbAPI.Item{Price: y[0].(float64), Volume: y[2].(float64)})
												r.OrderBookMAP[product+"bids"][y[0].(float64)] = y[2].(float64)
											} else {
												refLiveBook.Asks = append(refLiveBook.Asks, &pbAPI.Item{Price: y[0].(float64), Volume: y[2].(float64)})
												r.OrderBookMAP[product+"asks"][y[0].(float64)] = y[2].(float64) * -1
											}
										}
									} else {
										// book updates
										y := data
										if y[1].(float64) > 0 {
											if y[2].(float64) > 0 {
												price := y[0].(float64)
												amount := y[2].(float64)
												totalLevels := len(refLiveBook.GetBids())
												if totalLevels == r.base.MaxLevelsOrderBook {
													if price < refLiveBook.Bids[totalLevels-1].Price {
														continue
													}
												}
												updated = true
												r.OrderBookMAP[product+"bids"][price] = amount
											} else {
												price := y[0].(float64)
												amount := y[2].(float64) * -1
												totalLevels := len(refLiveBook.GetAsks())
												if totalLevels == r.base.MaxLevelsOrderBook {
													if price > refLiveBook.Asks[totalLevels-1].Price {
														continue
													}
												}
												updated = true
												r.OrderBookMAP[product+"asks"][price] = amount
											}
										} else if y[1].(float64) == 0 {
											if y[2].(float64) == 1 {
												price := y[0].(float64)
												if _, ok := r.OrderBookMAP[product+"bids"][price]; ok {
													delete(r.OrderBookMAP[product+"bids"], price)
													updated = true
												}
											} else if y[2].(float64) == -1 {
												price := y[0].(float64) * -1
												if _, ok := r.OrderBookMAP[product+"asks"][price]; ok {
													delete(r.OrderBookMAP[product+"asks"], price)
													updated = true
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
											Product:   pbAPI.Product((pbAPI.Product_value[product])),
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
											kafkaproducer.PublishMessageAsync(product+"."+r.base.Name+".orderbook", serialized, 1, false)
										}
									}
								}
							case "trades":
								if len(chanData) == 3 {
									if chanData[1] == "te" {
										data := chanData[2].([]interface{})

										var side pbAPI.Side
										var size float64
										if data[2].(float64) > 0 {
											side = pbAPI.Side_BUY
											size = data[2].(float64)
										} else {
											side = pbAPI.Side_SELL
											size = data[2].(float64) * -1
										}
										refBook, ok := r.base.LiveOrderBook.Get(product)
										if !ok {
											continue
										}
										refLiveBook := refBook.(*pbAPI.Orderbook)
										trades := &pbAPI.Trade{
											Product:   pbAPI.Product((pbAPI.Product_value[product])),
											Venue:     pbAPI.Venue((pbAPI.Venue_value[r.base.GetName()])),
											Timestamp: common.MakeTimestamp(),
											Price:     data[3].(float64),
											OrderSide: side,
											Volume:    size,
											VenueType: pbAPI.VenueType_SPOT,
											Asks:      refLiveBook.Asks,
											Bids:      refLiveBook.Bids,
										}
										serialized, err := proto.Marshal(trades)
										if err != nil {
											log.Fatal("proto.Marshal error: ", err)
										}
										r.MessageType[0] = 0
										serialized = append(r.MessageType, serialized[:]...)
										kafkaproducer.PublishMessageAsync(product+"."+r.base.Name+".trade", serialized, 1, false)
									}
								}
							}
						}

						// we dont need to update the book if any level we care was changed
						// if !updated {
						// 	continue
						// }

						// wg.Add(1)
						// go func() {
						// 	refLiveBook.Bids = []*pbAPI.Item{}
						// 	for price, amount := range r.OrderBookMAP[product+"bids"] {
						// 		refLiveBook.Bids = append(refLiveBook.Bids, &pbAPI.Item{Price: price, Volume: amount})
						// 	}
						// 	sort.Slice(refLiveBook.Bids, func(i, j int) bool {
						// 		return refLiveBook.Bids[i].Price > refLiveBook.Bids[j].Price
						// 	})
						// 	wg.Done()
						// }()

						// wg.Add(1)
						// go func() {
						// 	refLiveBook.Asks = []*pbAPI.Item{}
						// 	for price, amount := range r.OrderBookMAP[product+"asks"] {
						// 		refLiveBook.Asks = append(refLiveBook.Asks, &pbAPI.Item{Price: price, Volume: amount})
						// 	}
						// 	sort.Slice(refLiveBook.Asks, func(i, j int) bool {
						// 		return refLiveBook.Asks[i].Price < refLiveBook.Asks[j].Price
						// 	})
						// 	wg.Done()
						// }()

						// wg.Wait()

						// wg.Add(1)
						// go func() {
						// 	totalBids := len(refLiveBook.Bids)
						// 	if totalBids > r.base.MaxLevelsOrderBook {
						// 		refLiveBook.Bids = refLiveBook.Bids[0:r.base.MaxLevelsOrderBook]
						// 	}
						// 	wg.Done()
						// }()
						// wg.Add(1)
						// go func() {
						// 	totalAsks := len(refLiveBook.Asks)
						// 	if totalAsks > r.base.MaxLevelsOrderBook {
						// 		refLiveBook.Asks = refLiveBook.Asks[0:r.base.MaxLevelsOrderBook]
						// 	}
						// 	wg.Done()
						// }()

						// wg.Wait()
						// book := &pbAPI.Orderbook{
						// 	Product:   pbAPI.Product((pbAPI.Product_value[product])),
						// 	Venue:     pbAPI.Venue((pbAPI.Venue_value[r.base.GetName()])),
						// 	Levels:    int32(r.base.MaxLevelsOrderBook),
						// 	Timestamp: common.MakeTimestamp(),
						// 	Asks:      refLiveBook.Asks,
						// 	Bids:      refLiveBook.Bids,
						// 	VenueType: pbAPI.VenueType_SPOT,
						// }
						// refLiveBook = book

						// if r.base.Streaming {
						// 	serialized, err := proto.Marshal(book)
						// 	if err != nil {
						// 		log.Fatal("proto.Marshal error: ", err)
						// 	}
						// 	r.MessageType[0] = 1
						// 	serialized = append(r.MessageType, serialized[:]...)
						// 	kafkaproducer.PublishMessageAsync(product+"."+r.base.Name+".orderbook", serialized, 1, false)
						// }

					}
				}
			}
		}
	}()
}

func (r *Websocket) setReadTimeOut(timeout int) {
	if r.Conn != nil && timeout != 0 {
		readTimeout := time.Duration(timeout) * time.Second
		r.Conn.SetReadDeadline(time.Now().Add(readTimeout))
	}

}
