package okex

import (

	//"encoding/json"

	"bytes"
	"compress/flate"
	"errors"
	"io/ioutil"
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
	"github.com/maurodelazeri/go-number"
	"github.com/maurodelazeri/lion/common"
	event "github.com/maurodelazeri/lion/events"
	"github.com/maurodelazeri/lion/marketdata"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	pbEvent "github.com/maurodelazeri/lion/protobuf/heraldsquareAPI"
	"github.com/maurodelazeri/lion/venues/config"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

// Subscribe subsribe public and private endpoints
func (r *Websocket) Subscribe() error {
	// Users may subscribe to one or more channels，The total length of multiple channels should not exceed 4,096 bytes. ，
	// https://www.okex.com/docs/en/#ws_swap-format
	// {"op": "subscribe", "args": ["spot/ticker:ETH-USDT","spot/candle60s:ETH-USDT"]}
	subscribe := MessageChannel{}
	subscribe.Op = "subscribe"
	for _, sym := range r.subscribedPairs {
		venueConf, ok := r.base.VenueConfig.Get(r.base.GetName())
		if ok {
			switch venueConf.(config.VenueConfig).Products[sym].GetKind() {
			case "spot":
				subscribe.Args = append(subscribe.Args, "spot/trade:"+venueConf.(config.VenueConfig).Products[sym].GetVenueSymbolIdentifier())
				subscribe.Args = append(subscribe.Args, "spot/depth:"+venueConf.(config.VenueConfig).Products[sym].GetVenueSymbolIdentifier())
			case "futures":
				subscribe.Args = append(subscribe.Args, "futures/trade:"+venueConf.(config.VenueConfig).Products[sym].GetVenueSymbolIdentifier())
				subscribe.Args = append(subscribe.Args, "futures/depth:"+venueConf.(config.VenueConfig).Products[sym].GetVenueSymbolIdentifier())
			case "options":
			case "swaps":
				subscribe.Args = append(subscribe.Args, "swap/trade:"+venueConf.(config.VenueConfig).Products[sym].GetVenueSymbolIdentifier())
				subscribe.Args = append(subscribe.Args, "swap/depth:"+venueConf.(config.VenueConfig).Products[sym].GetVenueSymbolIdentifier())
			default:
				logrus.Warn("Subscription mode not found")
			}
		}
	}

	if len(subscribe.Args) > 0 {
		json, err := common.JSONEncode(subscribe)
		if err != nil {
			logrus.Error("Subscription ", err)
		}
		err = r.Conn.WriteMessage(websocket.TextMessage, json)
		if err != nil {
			logrus.Error("Subscription ", err)
		}
	} else {
		logrus.Info("Nothing to subscribe")
	}
	return nil
}

// Heartbeat ...
func (r *Websocket) Heartbeat() {
	ticker := time.NewTicker(time.Second * 27)
	for {
		select {
		case <-ticker.C:
			err := r.Conn.WriteJSON("{'event':'ping'}")
			if err != nil {
				logrus.Error(err)
			}
		}
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

	for _, sym := range r.subscribedPairs {
		r.base.LiveOrderBook.Set(sym, &pbAPI.Orderbook{})
		r.OrderBookMAP[sym+"bids"] = make(map[float64]float64)
		r.OrderBookMAP[sym+"asks"] = make(map[float64]float64)
		venueConf, ok := r.base.VenueConfig.Get(r.base.GetName())
		if ok {
			r.pairsMapping.Set(venueConf.(config.VenueConfig).Products[sym].VenueSymbolIdentifier, sym)
			r.OrderbookTimestamps.Set(r.base.GetName()+sym, time.Now())
		}
	}

	for {
		nextItvl := bb.Duration()
		//u := url.URL{Scheme: "wss", Host: websocketURL, Path: "/websocket", RawQuery: "compress=true"}
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
			err = r.Subscribe()
			if err != nil {
				logrus.Printf("Websocket subscription error: %s\n", err)
			}
			break
		} else {
			eventID, _ := uuid.NewV4()
			eventData := event.CreateBaseEvent(eventID.String(), "connect", nil, time.Now().UTC().Format(time.RFC3339Nano), r.base.GetName(), true, 0, pbEvent.System_WINTER)
			event.PublishEvent(eventData, "events", int64(1), false)

			logrus.Println(err)
			logrus.Println("Dial: will try again in", nextItvl, "seconds.")
		}

		time.Sleep(nextItvl)
	}
}

// WsReadData reads data from the websocket connection
func (r *Websocket) WsReadData() (WebsocketResponse, error) {
	msgType, resp, err := r.Conn.ReadMessage()
	if err != nil {
		logrus.Error(r.base.Name, " problem to read: ", err)
		r.closeAndRecconect()
		return WebsocketResponse{}, errors.New("websocket: not connected")
	}

	var standardMessage []byte

	switch msgType {
	case websocket.TextMessage:
		standardMessage = resp

	case websocket.BinaryMessage:
		reader := flate.NewReader(bytes.NewReader(resp))
		standardMessage, err = ioutil.ReadAll(reader)
		reader.Close()
		if err != nil {
			return WebsocketResponse{}, err
		}
	}
	return WebsocketResponse{Raw: standardMessage}, nil
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
					resp, err := r.WsReadData()
					if err != nil {
						logrus.Error("Problem reading data ", err)
						continue
					}
					data := Message{}
					err = common.JSONDecode(resp.Raw, &data)
					if err != nil {
						if strings.Contains(string(resp.Raw), "pong") {
							continue
						} else {
							logrus.Error("Problem to JSONDecode ", err)
							continue
						}
					}
					switch {
					case common.StringContains(data.Table, "trade"):
						for _, tradeData := range data.Data {
							var side string
							if tradeData.Side == "buy" {
								side = "buy"
							} else {
								side = "sell"
							}
							value, exist := r.pairsMapping.Get(tradeData.InstrumentID)
							if !exist {
								continue
							}
							product := value.(string)
							trades := &pbAPI.Trade{
								Product:         product,
								VenueTradeId:    tradeData.TradeID,
								Venue:           r.base.GetName(),
								SystemTimestamp: time.Now().UTC().Format(time.RFC3339Nano),
								VenueTimestamp:  tradeData.Timestamp,
								Price:           number.FromString(tradeData.Price).Float64(),
								OrderSide:       side,
								Volume:          number.FromString(tradeData.Size).Float64(),
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
					case common.StringContains(data.Table, "depth"):
						if data.Action == "update" {

							var wg sync.WaitGroup
							updated := false
							product := ""
							for _, fullBook := range data.Data {
								product = fullBook.InstrumentID
								refBook, ok := r.base.LiveOrderBook.Get(product)
								if !ok {
									continue
								}
								refLiveBook := refBook.(*pbAPI.Orderbook)

								for _, bids := range fullBook.Bids {
									price := number.FromString(bids[0].(string)).Float64()
									amount := number.FromString(bids[1].(string)).Float64()
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

								for _, asks := range fullBook.Asks {
									price := number.FromString(asks[0].(string)).Float64()
									amount := number.FromString(asks[1].(string)).Float64()
									if asks[1].(string) == "0" {
										if _, ok := r.OrderBookMAP[product+"asks"][price]; ok {
											delete(r.OrderBookMAP[product+"asks"], price)
											updated = true
										}
									} else {
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
									VenueTimestamp:  fullBook.Timestamp,
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
						} else if data.Action == "partial" {
							for _, bookData := range data.Data {
								product := bookData.InstrumentID
								refBook, ok := r.base.LiveOrderBook.Get(product)
								if !ok {
									continue
								}
								refLiveBook := refBook.(*pbAPI.Orderbook)
								var wg sync.WaitGroup
								wg.Add(1)
								go func(arr [][]interface{}) {
									total := 0
									for _, line := range arr {
										price := number.FromString(line[0].(string)).Float64()
										amount := number.FromString(line[1].(string)).Float64()
										if total > r.base.MaxLevelsOrderBook {
											continue
										}
										refLiveBook.Bids = append(refLiveBook.Bids, &pbAPI.Item{Price: price, Volume: amount})
										r.OrderBookMAP[product+"bids"][price] = amount
										total++
									}
									wg.Done()
								}(bookData.Bids)

								wg.Add(1)
								go func(arr [][]interface{}) {
									total := 0
									for _, line := range arr {
										price := number.FromString(line[0].(string)).Float64()
										amount := number.FromString(line[1].(string)).Float64()
										if total > r.base.MaxLevelsOrderBook {
											continue
										}
										refLiveBook.Asks = append(refLiveBook.Asks, &pbAPI.Item{Price: price, Volume: amount})
										r.OrderBookMAP[product+"asks"][price] = amount
										total++
									}
									wg.Done()
								}(bookData.Asks)
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
