package huobi

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
	"github.com/maurodelazeri/lion/common"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	"github.com/maurodelazeri/lion/streaming/kafka/producer"
	"github.com/maurodelazeri/lion/venues/config"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/sirupsen/logrus"
)

// Message ...
type Message struct {
	//Time          time.Time        `json:"time,string,omitempty"`
	HiveTable     string           `json:"hivetable,omitempty"`
	Type          string           `json:"type,omitempty"`
	ProductID     string           `json:"product_id,omitempty"`
	ProductIDs    []string         `json:"product_ids,omitempty"`
	TradeID       int64            `json:"trade_id,number,omitempty"`
	OrderID       string           `json:"order_id,omitempty"`
	Sequence      int64            `json:"sequence,number,omitempty"`
	MakerOrderID  string           `json:"maker_order_id,omitempty"`
	TakerOrderID  string           `json:"taker_order_id,omitempty"`
	RemainingSize float64          `json:"remaining_size,string,omitempty"`
	NewSize       float64          `json:"new_size,string,omitempty"`
	OldSize       float64          `json:"old_size,string,omitempty"`
	Size          float64          `json:"size,string,omitempty"`
	Price         float64          `json:"price,string,omitempty"`
	Side          string           `json:"side,omitempty"`
	Reason        string           `json:"reason,omitempty"`
	OrderType     string           `json:"order_type,omitempty"`
	Funds         float64          `json:"funds,string,omitempty"`
	NewFunds      float64          `json:"new_funds,string,omitempty"`
	OldFunds      float64          `json:"old_funds,string,omitempty"`
	Message       string           `json:"message,omitempty"`
	Bids          [][]string       `json:"bids,omitempty"`
	Asks          [][]string       `json:"asks,omitempty"`
	Changes       [][]string       `json:"changes,omitempty"`
	LastSize      float64          `json:"last_size,string,omitempty"`
	BestBid       float64          `json:"best_bid,string,omitempty"`
	BestAsk       float64          `json:"best_ask,string,omitempty"`
	Channels      []MessageChannel `json:"channels,omitempty"`
	UserID        string           `json:"user_id,omitempty"`
	ProfileID     string           `json:"profile_id,omitempty"`
	Open24H       float64          `json:"open_24h,string,omitempty"`
	Volume24H     float64          `json:"volume_24h,string,omitempty"`
	Low24H        float64          `json:"low_24h,string,omitempty"`
	High24H       float64          `json:"high_24h,string,omitempty"`
	Volume30D     float64          `json:"volume_30d,string,omitempty"`
}

// MessageChannel ...
type MessageChannel struct {
	Name       string   `json:"name"`
	ProductIDs []string `json:"product_ids"`
}

// Subscribe subsribe public and private endpoints
func (r *Websocket) Subscribe(products []string) error {
	subscribe := Message{}
	if r.base.Streaming {
		subscribe = Message{
			Type: "subscribe",
			Channels: []MessageChannel{
				MessageChannel{
					Name:       "full",
					ProductIDs: products,
				},
				MessageChannel{
					Name:       "level2",
					ProductIDs: products,
				},
			},
		}
	} else {
		subscribe = Message{
			Type: "subscribe",
			Channels: []MessageChannel{
				MessageChannel{
					Name:       "level2",
					ProductIDs: products,
				},
			},
		}
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

	r.OrderBookMAP = make(map[string]map[float64]float64)
	r.base.LiveOrderBook = utils.NewConcurrentMap()
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
						data := Message{}
						err = ffjson.Unmarshal(resp, &data)
						if err != nil {
							logrus.Error(err)
							continue
						}
						if data.Type == "error" {
							logrus.Warn("Error ", string(resp))
						}
						value, exist := r.pairsMapping.Get(data.ProductID)
						if !exist {
							continue
						}
						product := value.(string)

						if data.Type == "l2update" {
							//start := time.Now()
							refBook, ok := r.base.LiveOrderBook.Get(product)
							if !ok {
								continue
							}
							refLiveBook := refBook.(*pbAPI.Orderbook)

							var wg sync.WaitGroup
							updated := false

							for _, data := range data.Changes {
								switch data[0] {
								case "buy":
									if data[2] == "0" {
										price := r.base.Strfloat(data[1])
										if _, ok := r.OrderBookMAP[product+"bids"][price]; ok {
											delete(r.OrderBookMAP[product+"bids"], price)
											updated = true
										}
									} else {
										price := r.base.Strfloat(data[1])
										amount := r.base.Strfloat(data[2])
										totalLevels := len(refLiveBook.GetBids())
										if totalLevels == r.base.MaxLevelsOrderBook {
											if price < refLiveBook.Bids[totalLevels-1].Price {
												continue
											}
										}
										updated = true
										r.OrderBookMAP[product+"bids"][price] = amount
									}
								case "sell":
									if data[2] == "0" {
										price := r.base.Strfloat(data[1])
										if _, ok := r.OrderBookMAP[product+"asks"][price]; ok {
											delete(r.OrderBookMAP[product+"asks"], price)
											updated = true
										}
									} else {
										price := r.base.Strfloat(data[1])
										amount := r.base.Strfloat(data[2])
										totalLevels := len(refLiveBook.GetAsks())
										if totalLevels == r.base.MaxLevelsOrderBook {
											if price > refLiveBook.Asks[totalLevels-1].Price {
												continue
											}
										}
										updated = true
										r.OrderBookMAP[product+"asks"][price] = amount
									}
								default:
									continue
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

						if data.Type == "match" {
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
								Price:     data.Price,
								OrderSide: side,
								Volume:    data.Size,
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

						if data.Type == "snapshot" {
							refBook, ok := r.base.LiveOrderBook.Get(product)
							if !ok {
								continue
							}
							refLiveBook := refBook.(*pbAPI.Orderbook)

							var wg sync.WaitGroup

							wg.Add(1)
							go func(arr [][]string) {
								total := 0
								for _, line := range arr {
									price := r.base.Strfloat(line[0])
									amount := r.base.Strfloat(line[1])
									if total > r.base.MaxLevelsOrderBook {
										continue
									}
									refLiveBook.Bids = append(refLiveBook.Bids, &pbAPI.Item{Price: price, Volume: amount})
									r.OrderBookMAP[product+"bids"][price] = amount
									total++
								}
								wg.Done()
							}(data.Bids)

							wg.Add(1)
							go func(arr [][]string) {
								total := 0
								for _, line := range arr {
									price := r.base.Strfloat(line[0])
									amount := r.base.Strfloat(line[1])
									if total > r.base.MaxLevelsOrderBook {
										continue
									}
									refLiveBook.Asks = append(refLiveBook.Asks, &pbAPI.Item{Price: price, Volume: amount})
									r.OrderBookMAP[product+"asks"][price] = amount
									total++
								}
								wg.Done()
							}(data.Asks)
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
	}()
}

func (r *Websocket) setReadTimeOut(timeout int) {
	if r.Conn != nil && timeout != 0 {
		readTimeout := time.Duration(timeout) * time.Second
		r.Conn.SetReadDeadline(time.Now().Add(readTimeout))
	}

}
