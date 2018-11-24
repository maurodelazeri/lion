package huobi

import (

	//"encoding/json"

	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/maurodelazeri/lion/streaming/kafka/producer"
	uuid "github.com/satori/go.uuid"

	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/jpillora/backoff"
	"github.com/maurodelazeri/concurrency-map-slice"
	"github.com/maurodelazeri/lion/common"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	"github.com/maurodelazeri/lion/venues/config"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/sirupsen/logrus"
)

// Subscribe subsribe public and private endpoints
func (r *Websocket) Subscribe(products []string) error {
	subscribe := []MessageChannel{}

	if r.base.Streaming {
		for _, product := range products {
			idBook, _ := uuid.NewV4()
			book := MessageChannel{
				Sub: fmt.Sprintf("market.%s.depth.%s", product, "step2"),
				ID:  idBook.String(),
			}
			subscribe = append(subscribe, book)
			idTrade, _ := uuid.NewV4()
			trade := MessageChannel{
				Sub: fmt.Sprintf("market.%s.trade.detail", product),
				ID:  idTrade.String(),
			}
			subscribe = append(subscribe, trade)
		}
	} else {
		for _, product := range products {
			idTrade, _ := uuid.NewV4()
			trade := MessageChannel{
				Sub: fmt.Sprintf("market.%s.trade.detail", product),
				ID:  idTrade.String(),
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

// Heartbeat ...
func (r *Websocket) Heartbeat() {
	go func() {
		for {
			if r.IsConnected() {
				json, err := common.JSONEncode(PingPong{Ping: time.Now().Unix()})
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
			time.Sleep(time.Second * 15)
		}
	}()
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
					_, resp, err := r.Conn.ReadMessage()
					if err != nil {
						logrus.Error(r.base.Name, " problem to read: ", err)
						r.closeAndRecconect()
						continue
					}
					reader := bytes.NewReader(resp)
					zr, err := gzip.NewReader(reader)
					if err != nil {
						logrus.Error("zip decompress:", err)
						return
					}
					msg, err := ioutil.ReadAll(zr)
					data := Message{}
					err = ffjson.Unmarshal(msg, &data)
					if err != nil {
						logrus.Error(err)
						continue
					}
					if strings.Contains(data.Ch, "depth") {
						symbol := strings.Replace(data.Ch, "market.", "", -1)
						symbol = strings.Replace(symbol, ".depth.step2", "", -1)
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
							for _, info := range data.Tick.Bids {
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
							for _, info := range data.Tick.Asks {
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

					} else {
						symbol := strings.Replace(data.Ch, "market.", "", -1)
						symbol = strings.Replace(symbol, ".trade.detail", "", -1)
						value, exist := r.pairsMapping.Get(symbol)
						if !exist {
							continue
						}
						product := value.(string)
						for _, values := range data.Tick.Data {
							var side pbAPI.Side
							if values.Direction == "buy" {
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
								Price:     values.Price,
								OrderSide: side,
								Volume:    values.Amount,
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
		}
	}()
}
