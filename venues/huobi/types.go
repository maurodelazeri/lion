package huobi

import (
	"encoding/json"
)

// MessageChannel defines a request data structure
type MessageChannel struct {
	Topic             string `json:"req,omitempty"`
	Subscribe         string `json:"sub,omitempty"`
	ClientGeneratedID string `json:"id,omitempty"`
}

// PingPong ...
type PingPong struct {
	Ping int64 `json:"ping"`
}

// WsDepth defines market depth websocket response
type WsDepth struct {
	Channel   string `json:"ch"`
	Timestamp int64  `json:"ts"`
	Tick      struct {
		Bids      []interface{} `json:"bids"`
		Asks      []interface{} `json:"asks"`
		Timestamp int64         `json:"ts"`
		Version   int64         `json:"version"`
	} `json:"tick"`
}

// WsTrade defines market trade websocket response
type WsTrade struct {
	Channel   string `json:"ch"`
	Timestamp int64  `json:"ts"`
	Tick      struct {
		ID        int64 `json:"id"`
		Timestamp int64 `json:"ts"`
		Data      []struct {
			Amount    float64     `json:"amount"`
			Timestamp int64       `json:"ts"`
			ID        json.Number `json:"id,number"`
			Price     float64     `json:"price"`
			Direction string      `json:"direction"`
		} `json:"data"`
	}
}

// WsResponse defines a response from the websocket connection when there
// is an error
type WsResponse struct {
	TS           int64  `json:"ts"`
	Status       string `json:"status"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
	Ping         int64  `json:"ping"`
	Channel      string `json:"ch"`
	Subscribed   string `json:"subbed"`
}

// WebsocketResponse defines generalised data from the websocket connection
type WebsocketResponse struct {
	Type int
	Raw  []byte
}
