package okex

import (

	//"encoding/json"

	"bytes"
	"compress/flate"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jpillora/backoff"
	"github.com/maurodelazeri/concurrency-map-slice"
	"github.com/maurodelazeri/lion/common"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	"github.com/maurodelazeri/lion/venues/config"
	"github.com/sirupsen/logrus"
)

// https://support.okcoin.com/hc/en-us/articles/360000754131-WebSocket-API-Reference
// https://github.com/okcoin-okex/API-docs-OKEx.com/blob/master/API-For-Futures-EN/WebSocket%20API%20for%20FUTURES.md

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
	Event   string `json:"event"`
	Channel string `json:"channel"`
}

// Subscribe subsribe public and private endpoints
func (r *Websocket) Subscribe(products []string) error {

	tradesBegin := ""
	tradesEnd := ""
	bookBegin := ""
	bookEnd := ""
	switch r.base.GetName() {
	case "OKEX_INTERNATIONAL_FUT":
		tradesBegin = "ok_sub_future"
		tradesEnd = "_trade_this_week"
		bookBegin = "ok_sub_future"
		bookEnd = "_depth_week"
	case "OKEX_INTERNATIONAL_SPOT":
		tradesBegin = "ok_sub_spot_"
		tradesEnd = "_deals"
		bookBegin = "ok_sub_spot_"
		bookEnd = "_depth_20"
	}

	subscribe := []MessageChannel{}
	if r.base.Streaming {
		for _, product := range products {
			book := MessageChannel{
				Event:   "addChannel",
				Channel: fmt.Sprintf(`%s%s%s`, bookBegin, product, bookEnd),
			}
			subscribe = append(subscribe, book)
			trade := MessageChannel{
				Event:   "addChannel",
				Channel: fmt.Sprintf(`%s%s%s`, tradesBegin, product, tradesEnd),
			}
			subscribe = append(subscribe, trade)
		}
	} else {
		for _, product := range products {
			trade := MessageChannel{
				Event:   "addChannel",
				Channel: fmt.Sprintf(`%s%s%s`, tradesBegin, product, tradesEnd),
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

// GzipDecode ...
func (r *Websocket) GzipDecode(in []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(in))
	defer reader.Close()
	return ioutil.ReadAll(reader)
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
		u := url.URL{Scheme: "wss", Host: websocketURL, Path: "/websocket", RawQuery: "compress=true"}
		wsConn, httpResp, err := r.dialer.Dial(u.String(), r.reqHeader)

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
					case websocket.TextMessage, websocket.BinaryMessage:
						var err error
						msg, err := r.GzipDecode(resp)
						if err != nil {
							log.Println(err)
							return
						}
						logrus.Warn("HERE ", string(msg))
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
