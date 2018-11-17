package zb

// Message ...
type Message struct {
	Asks      [][]float64 `json:"asks,omitempty"`
	Bids      [][]float64 `json:"bid,omitemptys"`
	Channel   string      `json:"channel,omitempty"`
	DataType  string      `json:"dataType,omitempty"`
	Timestamp int         `json:"timestamp,omitempty"`
	Data      []struct {
		Amount    string `json:"amount,omitempty"`
		Date      int    `json:"date,omitempty"`
		Price     string `json:"price,omitempty"`
		Tid       int    `json:"tid,omitempty"`
		TradeType string `json:"trade_type,omitempty"`
		Type      string `json:"type,omitempty"`
	} `json:"data,omitempty"`
}

// MessageChannel ...
type MessageChannel struct {
	Event   string `json:"event"`
	Channel string `json:"channel"`
}
