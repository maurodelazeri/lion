package poloniex

import (

	//"encoding/json"

	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jpillora/backoff"
	utils "github.com/maurodelazeri/concurrency-map-slice"
	"github.com/maurodelazeri/lion/common"
	event "github.com/maurodelazeri/lion/events"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	pbEvent "github.com/maurodelazeri/lion/protobuf/heraldsquareAPI"
	"github.com/maurodelazeri/lion/venues/config"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

// https://poloniex.com/support/api/#reference_currencypairs

// Subscribe subsribe public and private endpoints
func (r *Websocket) Subscribe(products []string) error {
	for _, product := range products {
		orderbookJSON, err := common.JSONEncode(MessageChannel{
			Command: "subscribe",
			Channel: product,
		})
		if err != nil {
			return err
		}
		err = r.Conn.WriteMessage(websocket.TextMessage, orderbookJSON)
		if err != nil {
			return err
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
			if r.base.Verbose {
				logrus.Println(err)
				logrus.Println("Dial: will try again in", nextItvl, "seconds.")
			}
		}

		time.Sleep(nextItvl)
	}
}

func checkSubscriptionSuccess(data []interface{}) bool {
	return data[1].(float64) == 1
}

func getWSDataType(data interface{}) string {
	subData := data.([]interface{})
	dataType := subData[0].(string)
	return dataType
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
						err = common.JSONDecode(resp, &result)
						if err != nil {
							logrus.Error("JSONDecode ", err)
							continue
						}

						data := result.([]interface{})
						chanID := int(data[0].(float64))

						if len(data) == 2 && chanID != wsHeartbeat {
							if checkSubscriptionSuccess(data) {
								if r.base.Verbose {
									logrus.Debugf("poloniex websocket subscribed to channel successfully. %d", chanID)
								}
							} else {
								if r.base.Verbose {
									logrus.Debugf("poloniex websocket subscription to channel failed. %d", chanID)
								}
							}
							continue
						}

						switch chanID {
						case wsAccountNotificationID:
						case wsTickerDataID:
						case ws24HourExchangeVolumeID:
						case wsHeartbeat:
						default:
							if len(data) > 2 {
								subData := data[2].([]interface{})

								for x := range subData {
									dataL2 := subData[x]
									dataL3 := dataL2.([]interface{})

									switch getWSDataType(dataL2) {
									case "i":
										dataL3map := dataL3[1].(map[string]interface{})
										currencyPair, ok := dataL3map["currencyPair"].(string)
										if !ok {
											logrus.Error("poloniex.go error - could not find currency pair in map")
											continue
										}

										orderbookData, ok := dataL3map["orderBook"].([]interface{})
										if !ok {
											logrus.Error("poloniex.go error - could not find orderbook data in map")
											continue
										}

										// err := p.WsProcessOrderbookSnapshot(orderbookData, currencyPair)
										// if err != nil {
										// 	logrus.Error("WsProcessOrderbookSnapshot ", err)
										// 	continue
										// }

										// p.Websocket.DataHandler <- exchange.WebsocketOrderbookUpdate{
										// 	Exchange: p.GetName(),
										// 	Asset:    "SPOT",
										// 	Pair:     pair.NewCurrencyPairFromString(currencyPair),
										// }

									case "o":
										// currencyPair := CurrencyPairID[chanID]
										// err := p.WsProcessOrderbookUpdate(dataL3, currencyPair)
										// if err != nil {
										// 	p.Websocket.DataHandler <- err
										// 	continue
										// }

										// p.Websocket.DataHandler <- exchange.WebsocketOrderbookUpdate{
										// 	Exchange: p.GetName(),
										// 	Asset:    "SPOT",
										// 	Pair:     pair.NewCurrencyPairFromString(currencyPair),
										// }
									case "t":
										// currencyPair := CurrencyPairID[chanID]
										// var trade WsTrade
										// trade.Symbol = CurrencyPairID[chanID]
										// trade.TradeID, _ = strconv.ParseInt(dataL3[1].(string), 10, 64)
										// // 1 for buy 0 for sell
										// side := "buy"
										// if dataL3[2].(float64) != 1 {
										// 	side = "sell"
										// }
										// trade.Side = side
										// trade.Volume, _ = strconv.ParseFloat(dataL3[3].(string), 64)
										// trade.Price, _ = strconv.ParseFloat(dataL3[4].(string), 64)
										// trade.Timestamp = int64(dataL3[5].(float64))

										// p.Websocket.DataHandler <- exchange.TradeData{
										// 	Timestamp:    time.Unix(trade.Timestamp, 0),
										// 	CurrencyPair: pair.NewCurrencyPairFromString(currencyPair),
										// 	Side:         trade.Side,
										// 	Amount:       trade.Volume,
										// 	Price:        trade.Price,
										// }
									}
								}
							}
						}

					}
				}
			}
		}
	}()
}
