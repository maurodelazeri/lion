package zb

// Message ...
type Message struct {
	Timestamp int         `json:"timestamp,omitempty"`
	DataType  string      `json:"dataType,omitempty"`
	Asks      [][]float64 `json:"asks,omitempty"`
	Bids      [][]float64 `json:"bids,omitempty"`
	Data      []struct {
		Amount    string `json:"amount,omitempty"`
		Price     string `json:"price,omitempty"`
		Tid       int    `json:"tid,omitempty"`
		Date      int    `json:"date,omitempty"`
		Type      string `json:"type,omitempty"`
		TradeType string `json:"trade_type,omitempty"`
	} `json:"data,omitempty"`
	Channel string `json:"channel,omitempty"`
}

// MessageChannel ...
type MessageChannel struct {
	Event   string `json:"event"`
	Channel string `json:"channel"`
}
