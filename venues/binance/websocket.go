package binance

import (

	//"encoding/json"

	"errors"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jpillora/backoff"
	"github.com/maurodelazeri/concurrency-map-slice"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	"github.com/maurodelazeri/lion/venues/config"
	"github.com/sirupsen/logrus"
)

// Message ...
type Message struct {
	//Time          time.Time        `json:"time,string,omitempty"`
	HiveTable     string     `json:"hivetable,omitempty"`
	Type          string     `json:"type,omitempty"`
	ProductID     string     `json:"product_id,omitempty"`
	ProductIDs    []string   `json:"product_ids,omitempty"`
	TradeID       int64      `json:"trade_id,number,omitempty"`
	OrderID       string     `json:"order_id,omitempty"`
	Sequence      int64      `json:"sequence,number,omitempty"`
	MakerOrderID  string     `json:"maker_order_id,omitempty"`
	TakerOrderID  string     `json:"taker_order_id,omitempty"`
	RemainingSize float64    `json:"remaining_size,string,omitempty"`
	NewSize       float64    `json:"new_size,string,omitempty"`
	OldSize       float64    `json:"old_size,string,omitempty"`
	Size          float64    `json:"size,string,omitempty"`
	Price         float64    `json:"price,string,omitempty"`
	Side          string     `json:"side,omitempty"`
	Reason        string     `json:"reason,omitempty"`
	OrderType     string     `json:"order_type,omitempty"`
	Funds         float64    `json:"funds,string,omitempty"`
	NewFunds      float64    `json:"new_funds,string,omitempty"`
	OldFunds      float64    `json:"old_funds,string,omitempty"`
	Message       string     `json:"message,omitempty"`
	Bids          [][]string `json:"bids,omitempty"`
	Asks          [][]string `json:"asks,omitempty"`
	Changes       [][]string `json:"changes,omitempty"`
	LastSize      float64    `json:"last_size,string,omitempty"`
	BestBid       float64    `json:"best_bid,string,omitempty"`
	BestAsk       float64    `json:"best_ask,string,omitempty"`
	UserID        string     `json:"user_id,omitempty"`
	ProfileID     string     `json:"profile_id,omitempty"`
	Open24H       float64    `json:"open_24h,string,omitempty"`
	Volume24H     float64    `json:"volume_24h,string,omitempty"`
	Low24H        float64    `json:"low_24h,string,omitempty"`
	High24H       float64    `json:"high_24h,string,omitempty"`
	Volume30D     float64    `json:"volume_30d,string,omitempty"`
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

// FetchEnabledOrderBooks gets the initial orderbook
// func (b *Websocket) FetchEnabledOrderBooks() {
// 	go func() {
// 		for _, p := range b.ExchangeEnabledPairs {
// 			time.Sleep(1 * time.Second)

// 			orderbookNew, err := b.GetOrderBook(p, 50)
// 			position, exist := b.StringInSlice(p, b.ExchangeEnabledPairs)
// 			if err != nil {
// 				logrus.Error(b.Name, " ", p, " Problem to Fetch orderbook ", err)
// 				time.Sleep(5 * time.Second)
// 				continue
// 			}
// 			if exist {
// 				symbol := strings.Split(b.APIEnabledPairs[position], "/")
// 				for _, bids := range orderbookNew.Bids {
// 					b.redisSession.HSet("exchange:"+b.Name+":book:"+symbol[0]+symbol[1]+":bids", b.FloatToString(bids.Price), b.FloatToString(bids.Quantity))
// 				}
// 				for _, asks := range orderbookNew.Asks {
// 					b.redisSession.HSet("exchange:"+b.Name+":book:"+symbol[0]+symbol[1]+":asks", b.FloatToString(asks.Price), b.FloatToString(asks.Quantity))
// 				}

// 				b.redisSession.Expire("exchange:"+b.Name+":book:"+symbol[0]+symbol[1]+":asks", 1*time.Minute)
// 				b.redisSession.Expire("exchange:"+b.Name+":book:"+symbol[0]+symbol[1]+":bids", 1*time.Minute)

// 				b.LockTillBookFetchToFinish[strings.ToLower(p)+"@depth"] = symbol[0] + symbol[1]
// 			}
// 		}
// 	}()
// }

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

		currencies := []string{}
		for _, x := range venueArrayPairs {
			currencies = append(currencies, strings.ToLower(x)+"@trade")
			currencies = append(currencies, strings.ToLower(x)+"@depth")
		}

		wsConn, httpResp, err := r.dialer.Dial(websocketURL+strings.Join(currencies, "_"), r.reqHeader)

		r.mu.Lock()
		r.Conn = wsConn
		r.dialErr = err
		r.isConnected = err == nil
		r.httpResp = httpResp
		r.mu.Unlock()

		if err != nil {
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
						logrus.Warn(string(resp))
						// data := Message{}
						// err = ffjson.Unmarshal(resp, &data)
						// if err != nil {
						// 	logrus.Error(err)
						// 	continue
						// }

						// value, exist := r.pairsMapping.Get(data.ProductID)
						// if !exist {
						// 	continue
						// }
						// product := value.(string)

						// we dont need to update the book if any level we care was changed
						// if !updated {
						// 	continue
						// }

						// wg.Add(1)
						// go func() {
						// 	refLiveBook.Bids = []*pbAPI.Item{}
						// 	for price, amount := range r.OrderBookMAP[product+"bids"] {
						// 		refLiveBook.Bids = append(refLiveBook.Bids, &pbAPI.Item{Price: price, Volume: amount})
						// 	}
						// 	sort.Slice(refLiveBook.Bids, func(i, j int) bool {
						// 		return refLiveBook.Bids[i].Price > refLiveBook.Bids[j].Price
						// 	})
						// 	wg.Done()
						// }()

						// wg.Add(1)
						// go func() {
						// 	refLiveBook.Asks = []*pbAPI.Item{}
						// 	for price, amount := range r.OrderBookMAP[product+"asks"] {
						// 		refLiveBook.Asks = append(refLiveBook.Asks, &pbAPI.Item{Price: price, Volume: amount})
						// 	}
						// 	sort.Slice(refLiveBook.Asks, func(i, j int) bool {
						// 		return refLiveBook.Asks[i].Price < refLiveBook.Asks[j].Price
						// 	})
						// 	wg.Done()
						// }()

						// wg.Wait()

						// wg.Add(1)
						// go func() {
						// 	totalBids := len(refLiveBook.Bids)
						// 	if totalBids > r.base.MaxLevelsOrderBook {
						// 		refLiveBook.Bids = refLiveBook.Bids[0:r.base.MaxLevelsOrderBook]
						// 	}
						// 	wg.Done()
						// }()
						// wg.Add(1)
						// go func() {
						// 	totalAsks := len(refLiveBook.Asks)
						// 	if totalAsks > r.base.MaxLevelsOrderBook {
						// 		refLiveBook.Asks = refLiveBook.Asks[0:r.base.MaxLevelsOrderBook]
						// 	}
						// 	wg.Done()
						// }()

						// wg.Wait()
						// book := &pbAPI.Orderbook{
						// 	Product:   pbAPI.Product((pbAPI.Product_value[product])),
						// 	Venue:     pbAPI.Venue((pbAPI.Venue_value[r.base.GetName()])),
						// 	Levels:    int32(r.base.MaxLevelsOrderBook),
						// 	Timestamp: common.MakeTimestamp(),
						// 	Asks:      refLiveBook.Asks,
						// 	Bids:      refLiveBook.Bids,
						// 	VenueType: pbAPI.VenueType_SPOT,
						// }
						// refLiveBook = book

						// if r.base.Streaming {
						// 	serialized, err := proto.Marshal(book)
						// 	if err != nil {
						// 		log.Fatal("proto.Marshal error: ", err)
						// 	}
						// 	r.MessageType[0] = 1
						// 	serialized = append(r.MessageType, serialized[:]...)
						// 	kafkaproducer.PublishMessageAsync(product+"."+r.base.Name+".orderbook", serialized, 1, false)
						// }
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
