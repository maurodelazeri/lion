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

// WARN[0005] {"jsonrpc":"2.0","method":"subscription","params":{"channel":"book.BTC-PERPETUAL.raw","data":{"instrument_name":"BTC-PERPETUAL","bids":[["change",5485.0,727.0]],"asks":[],"prev_change_id":8399573118,"date":1542553161096,"change_id":8399573119}}}

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

// ,[7033.0,25.0],[7049.5,10.0],[7050.0,741.0],[7080.0,10.0],[7094.0,20.0],[7099.5,10.0],[7100.0,25.0],[7149.5,10.0],[7167.0,50.0],[7175.0,1.0],[7199.5,10.0],[73
