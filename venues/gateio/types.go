package gateio

// Message ...
type Message struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
	ID     interface{}   `json:"id"`
}

// MessageChannel ...
type MessageChannel struct {
	ID     int64         `json:"id"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}
