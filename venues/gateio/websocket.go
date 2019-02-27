package gateio

import (

	//"encoding/json"

	"errors"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"strings"
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
	if r.base.Streaming {
		var paramsBook []interface{}
		for _, prod := range products {
			paramsBook = append(paramsBook, []string{prod, "20", "0"})
		}
		book := MessageChannel{
			ID:     time.Now().Unix(),
			Method: "depth.subscribe",
			Params: paramsBook,
		}
		subscribe = append(subscribe, book)
		var paramsTrade []interface{}
		for _, prod := range products {
			paramsTrade = append(paramsTrade, prod)
		}
		trade := MessageChannel{
			ID:     time.Now().Unix(),
			Method: "trades.subscribe",
			Params: paramsTrade,
		}
		subscribe = append(subscribe, trade)
	} else {
		var params []interface{}
		for _, prod := range products {
			params = append(params, prod)
		}
		trade := MessageChannel{
			ID:     time.Now().Unix(),
			Method: "trades.subscribe",
			Params: params,
		}
		subscribe = append(subscribe, trade)
	}
	for _, channels := range subscribe {
		json, err := common.JSONEncode(channels)
		if err != nil {
			logrus.Error("Subscription ", err)
			continue
		}
		newbook := strings.Replace(string(json), "\"20\"", "20", -1)
		err = r.Conn.WriteMessage(websocket.TextMessage, []byte(newbook))
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

						result := Message{}
						err = ffjson.Unmarshal(resp, &result)
						if err != nil {
							logrus.Error(err)
							continue
						}
						switch {
						case common.StringContains(result.Method, "trades"):

							var trades []WebsocketTrade
							var c string
							err = common.JSONDecode(result.Params[1], &trades)
							if err != nil {
								continue
							}

							err = common.JSONDecode(result.Params[0], &c)
							if err != nil {
								continue
							}

							value, exist := r.pairsMapping.Get(c)
							if !exist {
								continue
							}
							product := value.(string)

							for _, trade := range trades {
								var side string
								if trade.Type == "buy" {
									side = "buy"
								} else {
									side = "sell"
								}
								sec, dec := math.Modf(trade.Time)
								trades := &pbAPI.Trade{
									Product:         product,
									VenueTradeId:    fmt.Sprintf("%d", trade.ID),
									Venue:           r.base.GetName(),
									SystemTimestamp: time.Now().UTC().Format(time.RFC3339Nano),
									VenueTimestamp:  time.Unix(int64(sec), int64(dec*(1e9))).UTC().Format(time.RFC3339Nano),
									Price:           trade.Price,
									OrderSide:       side,
									Volume:          trade.Amount,
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

						case common.StringContains(result.Method, "depth"):
							var wg sync.WaitGroup
							updated := false
							var IsSnapshot bool
							var c string
							var data = make(map[string][][]string)
							err = common.JSONDecode(result.Params[0], &IsSnapshot)
							if err != nil {
								continue
							}

							err = common.JSONDecode(result.Params[2], &c)
							if err != nil {
								continue
							}

							err = common.JSONDecode(result.Params[1], &data)
							if err != nil {
								continue
							}

							value, exist := r.pairsMapping.Get(c)
							if !exist {
								continue
							}
							product := value.(string)

							refBook, ok := r.base.LiveOrderBook.Get(product)
							if !ok {
								logrus.Warn("Skiping product ", product)
								continue
							}

							refLiveBook := refBook.(*pbAPI.Orderbook)
							//updated := false

							if IsSnapshot {
								askData, askOk := data["asks"]
								for _, ask := range askData {
									amount, _ := strconv.ParseFloat(ask[1], 64)
									price, _ := strconv.ParseFloat(ask[0], 64)
									refLiveBook.Asks = append(refLiveBook.Asks, &pbAPI.Item{Price: price, Volume: amount})
									r.OrderBookMAP[product+"asks"][price] = amount

								}

								bidData, bidOk := data["bids"]
								for _, bid := range bidData {
									amount, _ := strconv.ParseFloat(bid[1], 64)
									price, _ := strconv.ParseFloat(bid[0], 64)
									refLiveBook.Bids = append(refLiveBook.Bids, &pbAPI.Item{Price: price, Volume: amount})
									r.OrderBookMAP[product+"bids"][price] = amount

								}

								if !askOk && !bidOk {
									logrus.Error("gatio websocket error - cannot access ask or bid data")
									continue
								}

								continue
							} else {

								refBook, ok := r.base.LiveOrderBook.Get(product)
								if !ok {
									continue
								}
								refLiveBook := refBook.(*pbAPI.Orderbook)

								askData, askOk := data["asks"]
								if askOk {
									for _, ask := range askData {
										amount, _ := strconv.ParseFloat(ask[1], 64)
										price, _ := strconv.ParseFloat(ask[0], 64)
										if amount == 0 {
											if _, ok := r.OrderBookMAP[product+"asks"][price]; ok {
												delete(r.OrderBookMAP[product+"asks"], price)
												updated = true
											}
										} else {
											totalLevels := len(refLiveBook.GetBids())
											if totalLevels == r.base.MaxLevelsOrderBook {
												if len(refLiveBook.Asks) > 0 {
													if price < refLiveBook.Asks[totalLevels-1].Price {
														continue
													}
												}
											}
											updated = true
											r.OrderBookMAP[product+"asks"][price] = amount
										}
									}
								}

								bidData, bidOk := data["bids"]
								if bidOk {
									for _, bid := range bidData {
										amount, _ := strconv.ParseFloat(bid[1], 64)
										price, _ := strconv.ParseFloat(bid[0], 64)
										if amount == 0 {
											if _, ok := r.OrderBookMAP[product+"bids"][price]; ok {
												delete(r.OrderBookMAP[product+"bids"], price)
												updated = true
											}
										} else {
											totalLevels := len(refLiveBook.GetBids())
											if totalLevels == r.base.MaxLevelsOrderBook {
												if len(refLiveBook.Bids) > 0 {
													if price < refLiveBook.Bids[totalLevels-1].Price {
														continue
													}
												}
											}
											updated = true
											r.OrderBookMAP[product+"bids"][price] = amount
										}
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
								VenueTimestamp:  time.Now().UTC().Format(time.RFC3339Nano),
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

					}
				}
			}
		}
	}()
}
