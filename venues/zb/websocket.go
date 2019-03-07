package zb

import (

	//"encoding/json"

	"errors"
	"log"
	"math/rand"
	"net/http"
	"reflect"
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
	for _, product := range products {
		book := MessageChannel{
			Event:   "addChannel",
			Channel: product + "_depth",
		}
		subscribe = append(subscribe, book)
		trade := MessageChannel{
			Event:   "addChannel",
			Channel: product + "_trades",
		}
		subscribe = append(subscribe, trade)
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
	r.historyTrades = make(map[string]bool)

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
						err := common.JSONDecode(resp, &result)
						if err != nil {
							log.Println("Ops ", err)
							continue
						}
						switch reflect.TypeOf(result).String() {
						case "map[string]interface {}":
							data := Message{}
							err = ffjson.Unmarshal(resp, &data)
							if err != nil {
								logrus.Error(err)
								continue
							}
							if strings.Contains(data.Channel, "_depth") {
								symbol := strings.Replace(data.Channel, "_depth", "", -1)
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

								var wg sync.WaitGroup

								wg.Add(1)
								go func() {
									refLiveBook.Bids = []*pbAPI.Item{}
									for _, info := range data.Bids {
										refLiveBook.Bids = append(refLiveBook.Bids, &pbAPI.Item{Price: info[0], Volume: info[1]})
									}
									sort.Slice(refLiveBook.Bids, func(i, j int) bool {
										return refLiveBook.Bids[i].Price > refLiveBook.Bids[j].Price
									})
									wg.Done()
								}()

								wg.Add(1)
								go func() {
									refLiveBook.Asks = []*pbAPI.Item{}
									for _, info := range data.Asks {
										refLiveBook.Asks = append(refLiveBook.Asks, &pbAPI.Item{Price: info[0], Volume: info[1]})
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
									VenueTimestamp:  time.Unix(data.Timestamp, 0).UTC().Format(time.RFC3339Nano),
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

							} else {
								symbol := strings.Replace(data.Channel, "_trades", "", -1)
								value, exist := r.pairsMapping.Get(symbol)
								if !exist {
									continue
								}
								product := value.(string)
								for _, values := range data.Data {
									if r.historyTrades[product] {
										var side string
										if values.Type == "buy" {
											side = "buy"
										} else {
											side = "sell"
										}
										trades := &pbAPI.Trade{
											Product:         product,
											Venue:           r.base.GetName(),
											VenueTradeId:    strconv.FormatInt(values.Tid, 10),
											SystemTimestamp: time.Now().UTC().Format(time.RFC3339Nano),
											VenueTimestamp:  time.Unix(values.Date, 0).UTC().Format(time.RFC3339Nano),
											Price:           values.Price,
											OrderSide:       side,
											Volume:          values.Amount,
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
								r.historyTrades[product] = true
							}
						}
					}
				}
			}
		}
	}()
}
