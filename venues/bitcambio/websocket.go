package bitcambio

import (

	//"encoding/json"

	"errors"
	"log"
	"math/rand"
	"net/http"
	"sort"
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

// Subscribe subsribe public and private endpoints
func (r *Websocket) Subscribe(products []string) error {
	// book contain all data, including trades
	subscribe := []MessageChannel{}
	if r.base.Streaming {
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
	} else {
		book := MessageChannel{
			MsgType:                 "V",
			MDReqID:                 time.Now().Unix(),
			SubscriptionRequestType: "1",
			MarketDepth:             "0",
			MDUpdateType:            "1",
			MDEntryTypes:            []string{"2"},
			Instruments:             products,
		}
		subscribe = append(subscribe, book)
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

	r.setReadTimeOut(1)

	bb := &backoff.Backoff{
		Min:    r.RecIntvlMin,
		Max:    r.RecIntvlMax,
		Factor: r.RecIntvlFactor,
		Jitter: true,
	}

	rand.Seed(time.Now().UTC().UnixNano())

	r.OrderBookMAP = make(map[string]map[int64]BookItem)
	r.base.LiveOrderBook = utils.NewConcurrentMap()
	r.pairsMapping = utils.NewConcurrentMap()

	venueArrayPairs := []string{}
	for _, sym := range r.subscribedPairs {
		r.base.LiveOrderBook.Set(sym, &pbAPI.Orderbook{})
		r.OrderBookMAP[sym+"bids"] = make(map[int64]BookItem)
		r.OrderBookMAP[sym+"asks"] = make(map[int64]BookItem)
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

//https://blinktrade.com/docs/#subscribe-to-orderbook
// https://github.com/blinktrade/BlinkTradeJS/blob/3bba1a4154a5f1d69638938e2846ba2a6b77c58e/test/websocket.spec.js

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
						if len(message.MDIncGrp) > 0 {
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

							for _, data := range message.MDIncGrp {
								switch data.MDEntryType { // “0” = Bid, “1” = Offer, “2” = Trade
								case "0":
									switch data.MDUpdateAction { // “0” = New, “1” = Update, “2” = Delete, “3” = Delete Thru
									case "0":
										price := number.NewDecimal(data.MDEntryPx, 8).Div(number.NewDecimal(1e8, 8)).Float64()
										amount := number.NewDecimal(data.MDEntrySize, 8).Div(number.NewDecimal(1e8, 8)).Float64()
										totalLevels := len(refLiveBook.GetBids())
										if totalLevels == r.base.MaxLevelsOrderBook {
											if price < refLiveBook.Bids[totalLevels-1].Price {
												continue
											}
										}
										updated = true
										r.OrderBookMAP[product+"bids"][data.MDEntryID] = BookItem{Price: price, Volume: amount}
									case "1":
										logrus.Warn("NOT SURE WHAT EXPECT HERE ", string(resp))
									case "2":
										ID := data.MDEntryID
										if _, ok := r.OrderBookMAP[product+"bids"][ID]; ok {
											delete(r.OrderBookMAP[product+"bids"], ID)
											updated = true
										}
									case "3":
										ID := data.MDEntryID
										if _, ok := r.OrderBookMAP[product+"asks"][ID]; ok {
											delete(r.OrderBookMAP[product+"asks"], ID)
											updated = true
										}
									}
								case "1":
									switch data.MDUpdateAction { // “0” = New, “1” = Update, “2” = Delete, “3” = Delete Thru
									case "0":
										price := number.NewDecimal(data.MDEntryPx, 8).Div(number.NewDecimal(1e8, 8)).Float64()
										amount := number.NewDecimal(data.MDEntrySize, 8).Div(number.NewDecimal(1e8, 8)).Float64()
										totalLevels := len(refLiveBook.GetAsks())
										if totalLevels == r.base.MaxLevelsOrderBook {
											if price > refLiveBook.Asks[totalLevels-1].Price {
												continue
											}
										}
										updated = true
										r.OrderBookMAP[product+"asks"][data.MDEntryID] = BookItem{Price: price, Volume: amount}
									case "1":
										logrus.Warn("NOT SURE WHAT EXPECT HERE ", string(resp))
									case "2":
										ID := data.MDEntryID
										if _, ok := r.OrderBookMAP[product+"asks"][ID]; ok {
											delete(r.OrderBookMAP[product+"asks"], ID)
											updated = true
										}
									case "3":
										ID := data.MDEntryID
										if _, ok := r.OrderBookMAP[product+"asks"][ID]; ok {
											delete(r.OrderBookMAP[product+"asks"], ID)
											updated = true
										}
									}
								case "2":
									var side pbAPI.Side
									if data.Side == "buy" {
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
										Price:     number.NewDecimal(data.MDEntryPx, 8).Div(number.NewDecimal(1e8, 8)).Float64(),
										OrderSide: side,
										Volume:    number.NewDecimal(data.MDEntrySize, 8).Div(number.NewDecimal(1e8, 8)).Float64(),
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

								// we dont need to update the book if any level we care was changed
								if !updated {
									continue
								}

								wg.Add(1)
								go func() {
									refLiveBook.Bids = []*pbAPI.Item{}
									for _, values := range r.OrderBookMAP[product+"bids"] {
										refLiveBook.Bids = append(refLiveBook.Bids, &pbAPI.Item{Price: values.Price, Volume: values.Volume})
									}
									sort.Slice(refLiveBook.Bids, func(i, j int) bool {
										return refLiveBook.Bids[i].Price > refLiveBook.Bids[j].Price
									})
									wg.Done()
								}()

								wg.Add(1)
								go func() {
									refLiveBook.Asks = []*pbAPI.Item{}
									for _, values := range r.OrderBookMAP[product+"asks"] {
										refLiveBook.Asks = append(refLiveBook.Asks, &pbAPI.Item{Price: values.Price, Volume: values.Volume})
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
						} else if len(message.MDFullGrp) > 0 {

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

							for _, values := range message.MDFullGrp {
								if values.MDEntryType == "0" {
									refLiveBook.Bids = append(refLiveBook.Bids, &pbAPI.Item{Id: values.MDEntryID, Price: number.NewDecimal(values.MDEntryPx, 8).Div(number.NewDecimal(1e8, 8)).Float64(), Volume: number.NewDecimal(values.MDEntrySize, 8).Div(number.NewDecimal(1e8, 8)).Float64()})
								} else if values.MDEntryType == "1" {
									refLiveBook.Asks = append(refLiveBook.Asks, &pbAPI.Item{Id: values.MDEntryID, Price: number.NewDecimal(values.MDEntryPx, 8).Div(number.NewDecimal(1e8, 8)).Float64(), Volume: number.NewDecimal(values.MDEntrySize, 8).Div(number.NewDecimal(1e8, 8)).Float64()})
								}
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
									r.OrderBookMAP[product+"asks"][value.Id] = BookItem{Price: value.Price, Volume: value.Volume}
								}
								wg.Done()
							}()

							wg.Add(1)
							go func() {
								for _, value := range refLiveBook.Bids {
									r.OrderBookMAP[product+"bids"][value.Id] = BookItem{Price: value.Price, Volume: value.Volume}
								}
								wg.Done()
							}()
							wg.Wait()
						}
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
