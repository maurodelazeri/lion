package foxbit

import (

	//"encoding/json"

	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/jpillora/backoff"
	utils "github.com/maurodelazeri/concurrency-map-slice"
	"github.com/maurodelazeri/lion/common"
	event "github.com/maurodelazeri/lion/events"
	"github.com/maurodelazeri/lion/marketdata"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	pbEvent "github.com/maurodelazeri/lion/protobuf/heraldsquareAPI"
	"github.com/maurodelazeri/lion/venues/config"
	"github.com/pquerna/ffjson/ffjson"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

// Subscribe subsribe public and private endpoints
func (r *Websocket) Subscribe(products []string) error {
	subscribe := []MessageChannel{}
	count := 0
	if r.base.Streaming {
		for _, product := range products {
			i, _ := strconv.Atoi(product)
			payload, _ := ffjson.Marshal(Payload{InstrumentID: i, OMSID: 1, IncludeLastCount: 1})
			count++
			trade := MessageChannel{
				M: 0,
				I: count,
				N: "SubscribeTrades",
				O: string(payload),
			}
			subscribe = append(subscribe, trade)

			payload, _ = ffjson.Marshal(Payload{InstrumentID: i, OMSID: 1, Depth: 20})
			count++
			book := MessageChannel{
				M: 0,
				I: count,
				N: "SubscribeLevel2",
				O: string(payload),
			}
			subscribe = append(subscribe, book)
		}
	} else {
		for _, product := range products {
			i, _ := strconv.Atoi(product)
			payload, _ := ffjson.Marshal(Payload{InstrumentID: i, OMSID: 1, IncludeLastCount: 1})
			count++
			trade := MessageChannel{
				M: 2,
				I: count,
				N: "SubscribeTrades",
				O: string(payload),
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

//https://docs.foxbit.com.br/EN/websocket_intro.html#

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
						message := MessageChannel{}
						err = ffjson.Unmarshal(resp, &message)
						if err != nil {
							logrus.Error("Problem Unmarshal ", err)
							continue
						}
						stream := Stream{}
						err := ffjson.Unmarshal([]byte("{\"streaming\":"+message.O+"}"), &stream)
						if err != nil {
							logrus.Error("Problem Unmarshal ", err)
							continue
						}
						var wg sync.WaitGroup
						updated := false
						var product string
						if len(stream.Streaming) > 0 {
							var index int
							switch message.N {
							case "Level2UpdateEvent":
								index = 7
							case "TradeDataUpdateEvent":
								index = 1
							case "SubscribeLevel2":
								index = 1
							case "SubscribeTrades":
								continue
							}
							value, exist := r.pairsMapping.Get(strconv.Itoa(int(stream.Streaming[0][index].(float64))))
							if !exist {
								logrus.Error("shit could not find the product")
								continue
							}
							product = value.(string)
						}
						refBook, ok := r.base.LiveOrderBook.Get(product)
						if !ok {
							continue
						}
						refLiveBook := refBook.(*pbAPI.Orderbook)
						// 							{
						//0     "MDUpdateID": 0,
						//1     "Accounts": 0,
						//2     "ActionDateTime": 635872032000000000,
						//3     "ActionType": 0,
						//4     "LastTradePrice": 0,
						//5     "Orders": 0,
						//6     "Price": 0,
						//7     "ProductPairCode": 0,
						//8     "Quantity": 0,
						//9     "Side": 0,
						// }

						// 1 - ActionType tem values 0,1,2 a que isso se refere?
						// -> 0=New, 1=Update, 2=Delete
						// 2 - quando o size vem 0 eu poss remover o nivel de preco?"
						// -> Pode sim, significa que a ordem não possue mais amount
						switch message.N {
						case "Level2UpdateEvent":
							for _, values := range stream.Streaming {
								if values[8].(float64) == 0 {
									price := values[6].(float64)
									switch values[9].(float64) {
									case 0:
										if _, ok := r.OrderBookMAP[product+"bids"][price]; ok {
											delete(r.OrderBookMAP[product+"bids"], price)
											updated = true
										}
									case 1:
										if _, ok := r.OrderBookMAP[product+"asks"][price]; ok {
											delete(r.OrderBookMAP[product+"asks"], price)
											updated = true
										}
									}
								} else {
									switch values[9].(float64) {
									case 0:
										price := values[6].(float64)
										amount := values[8].(float64)
										totalLevels := len(refLiveBook.GetBids())
										if totalLevels == r.base.MaxLevelsOrderBook {
											if price < refLiveBook.Bids[totalLevels-1].Price {
												continue
											}
										}
										updated = true
										r.OrderBookMAP[product+"bids"][price] = amount
									case 1:
										price := values[6].(float64)
										amount := values[8].(float64)
										totalLevels := len(refLiveBook.GetAsks())
										if totalLevels == r.base.MaxLevelsOrderBook {
											if price > refLiveBook.Asks[totalLevels-1].Price {
												continue
											}
										}
										updated = true
										r.OrderBookMAP[product+"asks"][price] = amount

									case 2:
										logrus.Warn("orderbook case 2 ", string(resp))
									default:
										logrus.Warn("orderbook default case ", string(resp))
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

						case "TradeDataUpdateEvent":
							for _, values := range stream.Streaming {
								var side string
								if values[8].(float64) == 0 {
									side = "buy"
								} else {
									side = "sell"
								}
								trades := &pbAPI.Trade{
									Product:         product,
									VenueTradeId:    fmt.Sprintf("%f", values[0].(float64)),
									Venue:           r.base.GetName(),
									SystemTimestamp: time.Now().UTC().Format(time.RFC3339Nano),
									//VenueTimestamp:  values[6].(float64),
									Price:     values[3].(float64),
									OrderSide: side,
									Volume:    values[2].(float64),
								}
								logrus.Info(string(resp))
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
						case "SubscribeLevel2":
							for _, values := range stream.Streaming {
								if values[9].(float64) == 0 {
									refLiveBook.Bids = append(refLiveBook.Bids, &pbAPI.Item{Price: values[6].(float64), Volume: values[8].(float64)})
								} else {
									refLiveBook.Asks = append(refLiveBook.Asks, &pbAPI.Item{Price: values[6].(float64), Volume: values[8].(float64)})
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
							}
						}
					}
				}
			}
		}
	}()
}
