package poloniex

import (

	//"encoding/json"

	"errors"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"sort"
	"strconv"
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
	"github.com/sirupsen/logrus"
)

// https://poloniex.com/support/api/#reference_currencypairs

// Subscribe subsribe public and private endpoints
func (r *Websocket) Subscribe(products []string) error {
	subscribe := []MessageChannel{}
	if r.base.Streaming {
		for _, product := range products {
			book := MessageChannel{
				Command: "subscribe",
				Channel: product,
			}
			subscribe = append(subscribe, book)
			trade := MessageChannel{
				Command: "subscribe",
				Channel: product,
			}
			subscribe = append(subscribe, trade)
		}
	} else {
		for _, product := range products {
			trade := MessageChannel{
				Command: "subscribe",
				Channel: product,
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
						//logrus.Warn(string(resp))
						var result interface{}
						err := common.JSONDecode(resp, &result)
						if err != nil {
							log.Println("Ops ", err)
							continue
						}
						switch reflect.TypeOf(result).String() {
						case "[]interface {}":
							var wg sync.WaitGroup
							updated := false

							chanData := result.([]interface{})
							symbol := ""
							switch reflect.TypeOf(chanData[0]).String() {
							case "string":
								symbol = chanData[0].(string)
							case "float64":
								symbol = strconv.Itoa(int(chanData[0].(float64)))
							}
							value, exist := r.pairsMapping.Get(symbol)
							if !exist {
								continue
							}
							product := value.(string)
							refBook, ok := r.base.LiveOrderBook.Get(product)
							if !ok {
								continue
							}
							refLiveBook := refBook.(*pbAPI.Orderbook)

							if len(chanData) <= 2 {
								continue
							}
							arrData := chanData[2].([]interface{})
							for _, data := range arrData {
								finalData := data.([]interface{})
								switch reflect.TypeOf(finalData[1]).String() {
								case "map[string]interface {}":
									eventData := finalData[1].(map[string]interface{})
									for key, value := range eventData {
										if key == "orderBook" {
											book := value.([]interface{})
											asks := book[0].(map[string]interface{})
											bids := book[1].(map[string]interface{})

											for price, volume := range asks {
												refLiveBook.Asks = append(refLiveBook.Asks, &pbAPI.Item{Price: number.FromString(price).Float64(), Volume: number.FromString(volume.(string)).Float64()})
											}
											for price, volume := range bids {
												refLiveBook.Bids = append(refLiveBook.Bids, &pbAPI.Item{Price: number.FromString(price).Float64(), Volume: number.FromString(volume.(string)).Float64()})
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
													r.OrderBookMAP[product+"asks"][value.Price] = value.Volume
												}
												wg.Done()
											}()

											wg.Add(1)
											go func() {
												for _, value := range refLiveBook.Bids {
													r.OrderBookMAP[product+"bids"][value.Price] = value.Volume
												}
												wg.Done()
											}()
											wg.Wait()
										}

									}
								case "float64", "string":
									switch finalData[0] {
									case "o":
										//										logrus.Warn(reflect.TypeOf(finalData[1]).String())
										if finalData[1].(float64) == 1 {
											if number.FromString(finalData[3].(string)).Float64() == 0 {
												price := number.FromString(finalData[2].(string)).Float64()
												if _, ok := r.OrderBookMAP[product+"bids"][price]; ok {
													delete(r.OrderBookMAP[product+"bids"], price)
													updated = true
												}
											} else {
												price := number.FromString(finalData[2].(string)).Float64()
												amount := number.FromString(finalData[3].(string)).Float64()
												totalLevels := len(refLiveBook.GetBids())
												if totalLevels == r.base.MaxLevelsOrderBook {
													if price < refLiveBook.Bids[totalLevels-1].Price {
														continue
													}
												}
												updated = true
												r.OrderBookMAP[product+"bids"][price] = amount
											}
										} else {
											if number.FromString(finalData[3].(string)).Float64() == 0 {
												price := number.FromString(finalData[2].(string)).Float64()
												if _, ok := r.OrderBookMAP[product+"asks"][price]; ok {
													delete(r.OrderBookMAP[product+"asks"], price)
													updated = true
												}
											} else {
												price := number.FromString(finalData[2].(string)).Float64()
												amount := number.FromString(finalData[3].(string)).Float64()
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
									case "t":
										var side pbAPI.Side
										if finalData[1] == "1" {
											side = pbAPI.Side_BUY
										} else {
											side = pbAPI.Side_SELL
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
											Price:     number.FromString(finalData[3].(string)).Float64(),
											OrderSide: side,
											Volume:    number.FromString(finalData[4].(string)).Float64(),
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

								default:
									logrus.Warn("DEF ", string(resp))
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
							//logrus.Warn("BIDS: ", book.Bids[0].Price, book.Bids[0].Volume)
							//logrus.Warn("ASKS: ", book.Asks[0].Price, book.Asks[0].Volume)

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
				}
			}
		}
	}()
}
