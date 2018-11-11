package bitmex

import (

	//"encoding/json"

	"encoding/json"
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
	"github.com/maurodelazeri/lion/common"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	"github.com/maurodelazeri/lion/streaming/kafka/producer"
	"github.com/maurodelazeri/lion/venues/config"
	"github.com/sirupsen/logrus"
)

// MessageChannel takes in subscription information
type MessageChannel struct {
	OP   string   `json:"op"`
	Args []string `json:"args"`
}

// Message ...
type Message struct {
	Table  string   `json:"table,omitempty"`
	Action string   `json:"action,omitempty"`
	Keys   []string `json:"keys,omitempty"`
	Types  struct {
		Symbol string `json:"symbol,omitempty"`
		ID     string `json:"id,omitempty"`
		Side   string `json:"side,omitempty"`
		Size   string `json:"size,omitempty"`
		Price  string `json:"price,omitempty"`
	} `json:"types,omitempty"`
	ForeignKeys struct {
		Symbol string `json:"symbol,omitempty"`
		Side   string `json:"side,omitempty"`
	} `json:"foreignKeys"`
	Attributes struct {
		Symbol string `json:"symbol,omitempty"`
		ID     string `json:"id,omitempty"`
	} `json:"attributes,omitempty"`
	Filter struct {
		Symbol string `json:"symbol,omitempty"`
	} `json:"filter,omitempty"`
	Data []struct {
		Symbol string  `json:"symbol,omitempty"`
		ID     int64   `json:"id,omitempty"`
		Side   string  `json:"side,omitempty"`
		Size   int     `json:"size,omitempty"`
		Price  float64 `json:"price,omitempty"`
	} `json:"data,omitempty"`
}

// BookItem ...
type BookItem struct {
	Price  float64
	Volume float64
}

// Subscribe subsribe public and private endpoints
func (r *Websocket) Subscribe(products []string) error {
	endpoint := []string{}
	subscribe := MessageChannel{}
	if r.base.Streaming {
		for _, x := range products {
			endpoint = append(endpoint, "orderBookL2:"+x)
			endpoint = append(endpoint, "trade:"+x)
		}
		subscribe = MessageChannel{"subscribe", endpoint}
	} else {
		for _, x := range products {
			endpoint = append(endpoint, "orderBookL2:"+x)
			endpoint = append(endpoint, "trade:"+x)
		}
		subscribe = MessageChannel{"subscribe", endpoint}
	}
	json, err := common.JSONEncode(subscribe)
	if err != nil {
		return err
	}
	err = r.Conn.WriteMessage(websocket.TextMessage, json)
	if err != nil {
		return err
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
						var message Message
						json.Unmarshal(resp, &message)
						switch message.Table {
						case "trade":
							if message.Action == "insert" {
								for _, data := range message.Data {
									var side pbAPI.Side
									if data.Side == "Buy" {
										side = pbAPI.Side_BUY
									} else {
										side = pbAPI.Side_SELL
									}
									value, exist := r.pairsMapping.Get(data.Symbol)
									if !exist {
										continue
									}
									product := value.(string)
									refBook, ok := r.base.LiveOrderBook.Get(product)
									if !ok {
										continue
									}
									refLiveBook := refBook.(*pbAPI.Orderbook)
									trades := &pbAPI.Trade{
										Product:   pbAPI.Product((pbAPI.Product_value[product])),
										Venue:     pbAPI.Venue((pbAPI.Venue_value[r.base.GetName()])),
										Timestamp: common.MakeTimestamp(),
										Price:     data.Price,
										OrderSide: side,
										Volume:    float64(data.Size),
										VenueType: pbAPI.VenueType_FUTURES,
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
						case "orderBookL2":
							var wg sync.WaitGroup
							value, exist := r.pairsMapping.Get(message.Data[0].Symbol)
							if !exist {
								continue
							}
							product := value.(string)

							refBook, ok := r.base.LiveOrderBook.Get(product)
							if !ok {
								continue
							}
							refLiveBook := refBook.(*pbAPI.Orderbook)

							switch message.Action {
							case "update":
								updated := false
								for _, data := range message.Data {
									if data.Side == "Buy" {
										if val, ok := r.OrderBookMAP[product+"bids"][data.ID]; ok {
											r.OrderBookMAP[product+"bids"][data.ID] = BookItem{Price: val.Price, Volume: float64(data.Size)}
											updated = true
										}
									} else {
										if val, ok := r.OrderBookMAP[product+"asks"][data.ID]; ok {
											r.OrderBookMAP[product+"asks"][data.ID] = BookItem{Price: val.Price, Volume: float64(data.Size)}
											updated = true
										}
									}
								}
								if !updated {
									continue
								}
							case "delete":
								updated := false
								for _, data := range message.Data {
									if data.Side == "Buy" {
										if _, ok := r.OrderBookMAP[product+"bids"][data.ID]; ok {
											delete(r.OrderBookMAP[product+"bids"], data.ID)
											updated = true
										}
									} else {
										if _, ok := r.OrderBookMAP[product+"asks"][data.ID]; ok {
											delete(r.OrderBookMAP[product+"asks"], data.ID)
											updated = true
										}
									}
								}
								if !updated {
									continue
								}
							case "insert":
								updated := false
								for _, data := range message.Data {
									if data.Side == "Buy" {
										totalLevels := len(refLiveBook.GetBids())
										if totalLevels == r.base.MaxLevelsOrderBook {
											if data.Price < refLiveBook.Bids[totalLevels-1].Price {
												continue
											}
										}
										updated = true
										r.OrderBookMAP[product+"bids"][data.ID] = BookItem{Price: data.Price, Volume: float64(data.Size)}
									} else {
										totalLevels := len(refLiveBook.GetAsks())
										if totalLevels == r.base.MaxLevelsOrderBook {
											if data.Price > refLiveBook.Asks[totalLevels-1].Price {
												continue
											}
										}
										updated = true
										r.OrderBookMAP[product+"asks"][data.ID] = BookItem{Price: data.Price, Volume: float64(data.Size)}
									}
								}
								if !updated {
									continue
								}
							case "partial":
								for _, data := range message.Data {
									if data.Side == "Buy" {
										refLiveBook.Bids = append(refLiveBook.Bids, &pbAPI.Item{Id: data.ID, Price: data.Price, Volume: float64(data.Size)})
									} else {
										refLiveBook.Asks = append(refLiveBook.Asks, &pbAPI.Item{Id: data.ID, Price: data.Price, Volume: float64(data.Size)})
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

								// Cut off the waste
								refLiveBook.Bids = refLiveBook.Bids[0:r.base.MaxLevelsOrderBook]
								refLiveBook.Asks = refLiveBook.Asks[0:r.base.MaxLevelsOrderBook]

								for _, bids := range refLiveBook.Bids {
									r.OrderBookMAP[product+"bids"][bids.Id] = BookItem{Price: bids.Price, Volume: bids.Volume}
								}
								for _, asks := range refLiveBook.Asks {
									r.OrderBookMAP[product+"asks"][asks.Id] = BookItem{Price: asks.Price, Volume: asks.Volume}
								}
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

							// if len(book.Asks) > 0 && len(book.Bids) > 0 {
							// 	if book.Bids[0].Price > 5000 {
							// 		log.Print("==")
							// 		logrus.Warn("BIDS: ", book.Bids[0].Price, book.Bids[0].Volume)
							// 		logrus.Warn("ASKS: ", book.Asks[0].Price, book.Asks[0].Volume)
							// 		log.Print("==\n")
							// 	}
							// }

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

func (r *Websocket) setReadTimeOut(timeout int) {
	if r.Conn != nil && timeout != 0 {
		readTimeout := time.Duration(timeout) * time.Second
		r.Conn.SetReadDeadline(time.Now().Add(readTimeout))
	}

}
