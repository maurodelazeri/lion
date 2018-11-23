package gateio

import (

	//"encoding/json"

	"errors"
	"log"
	"math/rand"
	"net/http"
	"strings"
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
						//						logrus.Warn(string(resp))
						data := Message{}
						err = ffjson.Unmarshal(resp, &data)
						if err != nil {
							logrus.Error(err)
							continue
						}
						if len(data.Params) > 0 {
							if data.Method == "trades.update" {
								value, exist := r.pairsMapping.Get(data.Params[0].(string))
								if !exist {
									continue
								}
								product := value.(string)
								refBook, ok := r.base.LiveOrderBook.Get(product)
								if !ok {
									continue
								}
								refLiveBook := refBook.(*pbAPI.Orderbook)
								params := data.Params[1].([]interface{})
								for _, p := range params {
									mapData := p.(map[string]interface{})
									var side pbAPI.Side
									if mapData["type"].(string) == "buy" {
										side = pbAPI.Side_BUY
									} else {
										side = pbAPI.Side_SELL
									}
									trades := &pbAPI.Trade{
										Product:   pbAPI.Product((pbAPI.Product_value[product])),
										Venue:     pbAPI.Venue((pbAPI.Venue_value[r.base.GetName()])),
										Timestamp: common.MakeTimestamp(),
										Price:     number.FromString(mapData["price"].(string)).Float64(),
										OrderSide: side,
										Volume:    number.FromString(mapData["amount"].(string)).Float64(),
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
							} else if data.Method == "depth.update" {
								value, exist := r.pairsMapping.Get(data.Params[2].(string))
								if !exist {
									continue
								}
								product := value.(string)
								refBook, ok := r.base.LiveOrderBook.Get(product)
								if !ok {
									continue
								}
								refLiveBook := refBook.(*pbAPI.Orderbook)
								params := data.Params[1].(map[string]interface{})
								for symbol, values := range params {
									if symbol == "bids" {
										for _, bids := range values.([]interface{}) {
											values := bids.([]interface{})
											refLiveBook.Bids = append(refLiveBook.Bids, &pbAPI.Item{Price: number.FromString(values[0].(string)).Float64(), Volume: number.FromString(values[1].(string)).Float64()})
										}
									} else {
										for _, asks := range values.([]interface{}) {
											values := asks.([]interface{})
											refLiveBook.Asks = append(refLiveBook.Asks, &pbAPI.Item{Price: number.FromString(values[0].(string)).Float64(), Volume: number.FromString(values[1].(string)).Float64()})
										}
									}
								}

								var wg sync.WaitGroup
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
					}
				}
			}
		}
	}()
}
