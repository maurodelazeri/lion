package gateio

import "encoding/json"

// Message ...
type Message struct {
	Time    int64          `json:"time"`
	Channel string         `json:"channel"`
	Event   string         `json:""`
	Error   WebsocketError `json:"error"`
	Result  struct {
		Status string `json:"status"`
	} `json:"result"`
	Method string            `json:"method"`
	Params []json.RawMessage `json:"params"`
}

// MessageChannel ...
type MessageChannel struct {
	ID     int64         `json:"id"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

// WebsocketError defines a websocket error type
type WebsocketError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// WebsocketTrade defines trade data
type WebsocketTrade struct {
	ID     int64   `json:"id"`
	Time   float64 `json:"time"`
	Price  float64 `json:"price,string"`
	Amount float64 `json:"amount,string"`
	Type   string  `json:"type"`
}
