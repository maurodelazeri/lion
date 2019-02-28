package poloniex

// Message ...
type Message struct {
	HiveTable string `json:"hivetable,omitempty"`
	Type      string `json:"type,omitempty"`
}

// MessageChannel defines the request params after a websocket connection has been established
type MessageChannel struct {
	Command string      `json:"command"`
	Channel interface{} `json:"channel"`
	APIKey  string      `json:"key,omitempty"`
	Payload string      `json:"payload,omitempty"`
	Sign    string      `json:"sign,omitempty"`
}
