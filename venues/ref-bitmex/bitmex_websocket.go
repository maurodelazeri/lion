package bitmex

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/maurodelazeri/zion/influx"

	"github.com/gorilla/websocket"
	"github.com/jpillora/backoff"
	"github.com/maurodelazeri/zion/common"
	"github.com/maurodelazeri/zion/currency/pair"
	"github.com/maurodelazeri/zion/exchanges/ticker"
	"github.com/sirupsen/logrus"
)

const (
	bitmexWebsocketURL = "wss://www.bitmex.com/realtime"
)

// WebsocketSubscribe subsribe public and private endpoints
func (b *Bitmex) WebsocketSubscribe(endpoint []string) error {
	subscribe := WebsocketSubscription{"subscribe", endpoint}
	json, err := common.JSONEncode(subscribe)
	if err != nil {
		return err
	}
	err = b.Conn.WriteMessage(websocket.TextMessage, json)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the underlying network connection without
// sending or waiting for a close frame.
func (b *Bitmex) Close() {
	b.mu.Lock()
	if b.Conn != nil {
		b.Conn.Close()
	}
	b.isConnected = false
	b.mu.Unlock()
}

func (b *Bitmex) Loop() {
	b.WebsocketClient()
}

func (b *Bitmex) connect() {

	bb := &backoff.Backoff{
		Min:    b.RecIntvlMin,
		Max:    b.RecIntvlMax,
		Factor: b.RecIntvlFactor,
		Jitter: true,
	}

	rand.Seed(time.Now().UTC().UnixNano())

	for {
		nextItvl := bb.Duration()

		wsConn, httpResp, err := b.dialer.Dial(b.url, b.reqHeader)

		b.mu.Lock()
		b.Conn = wsConn
		b.dialErr = err
		b.isConnected = err == nil
		b.httpResp = httpResp
		b.mu.Unlock()

		if err == nil {
			if b.Verbose {
				log.Printf("Dial: connection was successfully established with %s\n", b.url)
			}

			// Subscribe trade
			endpoint := []string{}
			for _, x := range b.ExchangeEnabledPairs {
				//	endpoint = append(endpoint, "quote:"+x)
				//	endpoint = append(endpoint, "trade:"+x)
				endpoint = append(endpoint, "instrument:"+x)
				//endpoint = append(endpoint, "settlement:"+x)
			}
			err = b.WebsocketSubscribe(endpoint)
			if err != nil {
				log.Printf("Websocket send data error:%v", err)
			}

			break
		} else {
			if b.Verbose {
				log.Println(err)
				log.Println("Dial: will try again in", nextItvl, "seconds.")
			}
		}

		time.Sleep(nextItvl)
	}
}

func (b *Bitmex) WebsocketClient() {

	if b.RecIntvlMin == 0 {
		b.RecIntvlMin = 2 * time.Second
	}

	if b.RecIntvlMax == 0 {
		b.RecIntvlMax = 30 * time.Second
	}

	if b.RecIntvlFactor == 0 {
		b.RecIntvlFactor = 1.5
	}

	if b.HandshakeTimeout == 0 {
		b.HandshakeTimeout = 2 * time.Second
	}

	b.dialer = websocket.DefaultDialer
	b.dialer.HandshakeTimeout = b.HandshakeTimeout

	// Start reading from the socket
	b.startReading()

	go func() {
		b.connect()
	}()

	// wait on first attempt
	time.Sleep(b.HandshakeTimeout)

}

// IsConnected returns the WebSocket connection state
func (b *Bitmex) IsConnected() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.isConnected
}

// CloseAndRecconect will try to reconnect.
func (b *Bitmex) closeAndRecconect() {
	b.Close()
	go func() {
		b.connect()
	}()
}

// GetHTTPResponse returns the http response from the handshake.
// Useful when WebSocket handshake fails,
// so that callers can handle redirects, authentication, etc.
func (b *Bitmex) GetHTTPResponse() *http.Response {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.httpResp
}

// GetDialError returns the last dialer error.
// nil on successful connection.
func (b *Bitmex) GetDialError() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.dialErr
}

// ReadMessage is a helper method for getting a reader
// using NextReader and reading from that reader to a buffer.
//
// If the connection is closed ErrNotConnected is returned
// ReadMessage initiates a websocket client
func (b *Bitmex) startReading() {
	go func() {
		for {
			select {
			default:
				if b.IsConnected() {
					err := ErrNotConnected
					msgType, resp, err := b.Conn.ReadMessage()
					if err != nil {
						logrus.Error(b.Name, " problem to read: ", err)
						b.closeAndRecconect()
						continue
					}
					switch msgType {
					case websocket.TextMessage:
						data := wsData{}
						err := common.JSONDecode(resp, &data)
						if err != nil {
							log.Println(err)
							continue
						}
						var table wsData
						json.Unmarshal(resp, &table)

						switch table.Table {
						case "settlement":
							logrus.Info(string(resp))
						case "instrument":
							var instrument []WSInstrument
							err := json.Unmarshal(table.Data, &instrument)
							if err != nil {
								logrus.Info("Errr to unmarshal instrumet ", err)
								continue
							}
							for _, one := range instrument {
								if one.LastPrice == 0 {
									continue
								}
								position, exist := b.StringInSlice(one.Symbol, b.ExchangeEnabledPairs)
								if exist {
									symbol := strings.Split(b.APIEnabledPairs[position], "/")

									var tick ticker.Price
									tick.CurrencyPair = strings.ToUpper(symbol[0]) + strings.ToUpper(symbol[1])
									tick.Pair = pair.NewCurrencyPair(strings.ToUpper(symbol[0]), strings.ToUpper(symbol[1]))
									tick.Exchange = b.Name
									tick.Ask = one.AskPrice
									tick.Bid = one.BidPrice
									tick.Low = one.LowPrice
									tick.Last = one.LastPrice
									tick.Volume = float64(one.Volume)
									tick.High = one.HighPrice
									tick.UpdateTime = time.Now()

									serialized, _ := json.Marshal(tick)
									b.redisSession.HSet("exchange:"+b.Name+":tick", strings.ToUpper(symbol[0])+strings.ToUpper(symbol[1]), serialized)
									b.redisSession.Expire("exchange:"+b.Name+":tick", 1*time.Minute)

									if b.InfluxStreaming {
										influxserver.MetricsPointCh <- influxserver.MakeMetricsPoint(tick)
									}

								}
							}
						case "quote":
							// var quotes []WSQuote
							// err := json.Unmarshal(table.Data, &quotes)
							// if err != nil {
							// 	continue
							// }
							// for _, one := range quotes {
							// 	logrus.Info("quote ", one)
							// }
						case "trade":
							// var trades []WSTrade
							// err := json.Unmarshal(table.Data, &trades)
							// if err != nil {
							// 	continue
							// }
							// for _, one := range trades {
							// 	logrus.Info(one)
							// }
						}
					}
				}
			}
		}
	}()
}
