package deribit

// Params ...
type Params struct {
	Channels []string `json:"channels"`
}

// MessageChannel ...
type MessageChannel struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int64  `json:"id"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
}

// Message ...
type Message struct {
	Jsonrpc string `json:"jsonrpc,omitempty"`
	Method  string `json:"method,omitempty"`
	Params  struct {
		Channel string `json:"channel,omitempty"`
		Data    struct {
			Instrument     string          `json:"instrument,omitempty"`
			InstrumentName string          `json:"instrument_name,omitempty"`
			Bids           [][]interface{} `json:"bids,omitempty"`
			Asks           [][]interface{} `json:"asks,omitempty"`
			PrevChangeID   int64           `json:"prev_change_id,omitempty"`
			Date           int64           `json:"date,omitempty"`
			ChangeID       int64           `json:"change_id,omitempty"`
			TradeID        int             `json:"tradeId,omitempty"`
			TimeStamp      int64           `json:"timeStamp,omitempty"`
			TradeSeq       int64           `json:"tradeSeq,omitempty"`
			Quantity       int             `json:"quantity,omitempty"`
			Amount         float64         `json:"amount,omitempty"`
			Price          float64         `json:"price,omitempty"`
			Direction      string          `json:"direction,omitempty"`
			TickDirection  int             `json:"tickDirection,omitempty"`
			IndexPrice     float64         `json:"indexPrice,omitempty"`
		} `json:"data,omitempty"`
	} `json:"params,omitempty"`
}

// PingPong ...
type PingPong struct {
	Action string `json:"action"`
}
