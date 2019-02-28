package okex

import "time"

// MessageTrade ...
type MessageTrade struct {
	Channel string          `json:"channel,omitempty"`
	Data    [][]interface{} `json:"data,omitempty"`
	Binary  int             `json:"binary,omitempty"`
}

// MessageBook ...
type MessageBook struct {
	Binary  int    `json:"binary,omitempty"`
	Channel string `json:"channel,omitempty"`
	Data    struct {
		Asks      [][]interface{} `json:"asks,omitempty"`
		Bids      [][]interface{} `json:"bids,omitempty"`
		Timestamp int64           `json:"timestamp,omitempty"`
	} `json:"data"`
}

// MessageChannel ...
type MessageChannel struct {
	Op   string   `json:"op"`
	Args []string `json:"args"`
}

// PingPong ...
type PingPong struct {
	Event int64 `json:"event"`
}

// WebsocketResponse defines generalised data from the websocket connection
type WebsocketResponse struct {
	Type int
	Raw  []byte
}

// Message ....
type Message struct {
	Table string `json:"table"`
	Data  []struct {
		Asks         [][]interface{} `json:"asks"`
		Bids         [][]interface{} `json:"bids"`
		Timestamp    time.Time       `json:"timestamp"`
		Checksum     int             `json:"checksum"`
		InstrumentID string          `json:"instrument_id"`
		Price        string          `json:"price"`
		Side         string          `json:"side"`
		Size         string          `json:"size"`
		TradeID      string          `json:"trade_id"`
	} `json:"data"`
}

// ErrorResponse defines an error response type from the websocket connection
type ErrorResponse struct {
	Result    bool   `json:"result"`
	ErrorMsg  string `json:"error_msg"`
	ErrorCode int64  `json:"error_code"`
}

// Request defines the JSON request structure to the websocket server
type Request struct {
	Event      string `json:"event"`
	Channel    string `json:"channel"`
	Parameters string `json:"parameters,omitempty"`
}

// DealsStreamData defines Deals data
type DealsStreamData = [][]string
