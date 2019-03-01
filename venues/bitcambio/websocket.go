package bitcambio

import (

	//"encoding/json"

	"errors"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/jpillora/backoff"
	utils "github.com/maurodelazeri/concurrency-map-slice"
	number "github.com/maurodelazeri/go-number"
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
	// book contain all data, including trades
	subscribe := []MessageChannel{}
	book := MessageChannel{
		MsgType:                 "V",
		MDReqID:                 time.Now().Unix(),
		SubscriptionRequestType: "1",
		MarketDepth:             "0",
		MDUpdateType:            "1",
		MDEntryTypes:            []string{"0", "1", "2"},
		Instruments:             products,
	}
	subscribe = append(subscribe, book)
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

	r.OrderBookMAP = make(map[string]map[int64]BookItem)
	r.pairsMapping = utils.NewConcurrentMap()
	r.OrderbookTimestamps = utils.NewConcurrentMap()

	venueArrayPairs := []string{}
	for _, sym := range r.subscribedPairs {
		r.base.LiveOrderBook.Set(sym, &pbAPI.Orderbook{})
		r.OrderBookMAP[sym+"bids"] = make(map[int64]BookItem)
		r.OrderBookMAP[sym+"asks"] = make(map[int64]BookItem)

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

func (r *Websocket) insert(currentSlice []*pbAPI.Item, at int64, val *pbAPI.Item) []*pbAPI.Item {
	logrus.Warn(currentSlice, " ---- ", at, " varl ", val, " len ", len(currentSlice))
	if len(currentSlice) > int(at) {
		// Move all elements of s up one slot
		copy(currentSlice[at+1:], currentSlice[at:])
		// Insert the new element at the now free position
		currentSlice[at] = val
		return currentSlice
	}
	return currentSlice
}

func (r *Websocket) delete(currentSlice []*pbAPI.Item, at int) []*pbAPI.Item {
	logrus.Info(currentSlice, " ---- ", at, " len ", len(currentSlice))
	if len(currentSlice) > at {
		newarr := currentSlice[:at+copy(currentSlice[at:], currentSlice[at+1:])]
		return newarr
	}
	return currentSlice
}

//https://blinktrade.com/docs/#subscribe-to-orderbook
// https://github.com/blinktrade/BlinkTradeJS/blob/3bba1a4154a5f1d69638938e2846ba2a6b77c58e/test/websocket.spec.js

//CODE EXA
// https://github.com/Judahh/cryptobot/blob/09780a8fb33abc8762ac236c37cdc300d7eb9410/api/apiConnection/foxbit/webSocket/foxbit.ts

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

						message := Message{}
						err = ffjson.Unmarshal(resp, &message)
						if err != nil {
							logrus.Error("Problem Unmarshal ", err)
							continue
						}
						var wg sync.WaitGroup
						updated := false
						switch message.MsgType {
						case "f":
						case "W":
							value, exist := r.pairsMapping.Get(message.Symbol)
							if !exist {
								continue
							}
							product := value.(string)
							refBook, ok := r.base.LiveOrderBook.Get(product)
							if !ok {
								continue
							}
							refLiveBook := refBook.(*pbAPI.Orderbook)
							for _, data := range message.MDFullGrp {
								switch data.MDEntryType {
								case "0":
									refLiveBook.Bids = append(refLiveBook.Bids, &pbAPI.Item{Price: number.NewDecimal(data.MDEntryPx, 8).Div(number.NewDecimal(1e8, 8)).Float64(), Volume: number.NewDecimal(data.MDEntrySize, 8).Div(number.NewDecimal(1e8, 8)).Float64()})
								case "1":
									refLiveBook.Asks = append(refLiveBook.Asks, &pbAPI.Item{Price: number.NewDecimal(data.MDEntryPx, 8).Div(number.NewDecimal(1e8, 8)).Float64(), Volume: number.NewDecimal(data.MDEntrySize, 8).Div(number.NewDecimal(1e8, 8)).Float64()})
								case "2":
								}
							}
						case "X":
							value, exist := r.pairsMapping.Get(message.MDIncGrp[0].Symbol)
							if !exist {
								continue
							}
							product := value.(string)
							refBook, ok := r.base.LiveOrderBook.Get(product)
							if !ok {
								continue
							}
							refLiveBook := refBook.(*pbAPI.Orderbook)
							updated = true

							// https://github.com/seanoflynn/circus/blob/722835469aa220da413e9a56b50dd1900f806192/src/Cme/Enums/MDUpdateAction.cs
							// INFO[0621] DELETE THU ask {"MDUpdateAction":"3","Symbol":"BTCBRL","UserID":0,"MDEntryType":"1","MDEntryPositionNo":2}
							// INFO[0670] DELETE THU ask {"MDUpdateAction":"3","Symbol":"BTCBRL","UserID":0,"MDEntryType":"1","MDEntryPositionNo":1}
							// INFO[1132] DELETE THU bid {"MDUpdateAction":"3","Symbol":"BTCBRL","UserID":0,"MDEntryType":"0","MDEntryPositionNo":1}

							for _, data := range message.MDIncGrp {
								switch data.MDEntryType {
								case "0": // Bid
									switch data.MDUpdateAction {
									case "0": //onMDNewOrder_
										refLiveBook.Bids = r.insert(refLiveBook.Bids, data.MDEntryPositionNo-1, &pbAPI.Item{Price: number.NewDecimal(data.MDEntryPx, 8).Div(number.NewDecimal(1e8, 8)).Float64(), Volume: number.NewDecimal(data.MDEntrySize, 8).Div(number.NewDecimal(1e8, 8)).Float64()})
									case "1": //onMDUpdateOrder_
										refLiveBook.Bids[data.MDEntryPositionNo-1] = &pbAPI.Item{Price: number.NewDecimal(data.MDEntryPx, 8).Div(number.NewDecimal(1e8, 8)).Float64(), Volume: number.NewDecimal(data.MDEntrySize, 8).Div(number.NewDecimal(1e8, 8)).Float64()}
									case "2": //onMDDeleteOrder_
										r.delete(refLiveBook.Bids, int(data.MDEntryPositionNo-1))
									case "3": //onMDDeleteOrderThru_
										// aaa, _ := ffjson.Marshal(data)
										// logrus.Info("DELETE THU bid ", string(aaa))
										refLiveBook.Bids = refLiveBook.Bids[data.MDEntryPositionNo:]
									}
								case "1": // Ask
									switch data.MDUpdateAction {
									case "0": //onMDNewOrder_
										refLiveBook.Asks = r.insert(refLiveBook.Asks, data.MDEntryPositionNo-1, &pbAPI.Item{Price: number.NewDecimal(data.MDEntryPx, 8).Div(number.NewDecimal(1e8, 8)).Float64(), Volume: number.NewDecimal(data.MDEntrySize, 8).Div(number.NewDecimal(1e8, 8)).Float64()})
									case "1": //onMDUpdateOrder_
										refLiveBook.Asks[data.MDEntryPositionNo-1] = &pbAPI.Item{Price: number.NewDecimal(data.MDEntryPx, 8).Div(number.NewDecimal(1e8, 8)).Float64(), Volume: number.NewDecimal(data.MDEntrySize, 8).Div(number.NewDecimal(1e8, 8)).Float64()}
									case "2": //onMDDeleteOrder_
										r.delete(refLiveBook.Asks, int(data.MDEntryPositionNo-1))
									case "3": //onMDDeleteOrderThru_
										// aaa, _ := ffjson.Marshal(data)
										// logrus.Info("DELETE THU ask ", string(aaa))
										refLiveBook.Asks = refLiveBook.Asks[data.MDEntryPositionNo:]
									}

								case "2": // Trade
									var side string
									if data.Side == "1" {
										side = "buy"
									} else {
										side = "sell"
									}

									layout := "2006-01-02T15:04:05.000Z"
									t, err := time.Parse(layout, data.MDEntryDate+"T"+data.MDEntryTime+".000Z")
									if err != nil {
										logrus.Error("problem converting date ", err)
										continue
									}

									trades := &pbAPI.Trade{
										Product:         product,
										VenueTradeId:    strconv.FormatInt(data.OrderID, 10),
										Venue:           r.base.GetName(),
										SystemTimestamp: time.Now().UTC().Format(time.RFC3339Nano),
										VenueTimestamp:  t.UTC().Format(time.RFC3339Nano),
										Price:           number.NewDecimal(data.MDEntryPx, 8).Div(number.NewDecimal(1e8, 8)).Float64(),
										OrderSide:       side,
										Volume:          number.NewDecimal(data.MDEntrySize, 8).Div(number.NewDecimal(1e8, 8)).Float64(),
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
									continue
								}
							}

							// we dont need to update the book if any level we care was changed
							if !updated {
								continue
							}

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
