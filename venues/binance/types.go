package binance

// OrderBook ...
type OrderBook struct {
	LastUpdateID int             `json:"lastUpdateId"`
	Bids         [][]interface{} `json:"bids"`
	Asks         [][]interface{} `json:"asks"`
}

// Message ...
type Message struct {
	Stream string `json:"stream"`
}

// MessageTrade ...
type MessageTrade struct {
	Stream string `json:"stream"`
	Data   struct {
		EventType     string  `json:"e"`
		EventTime     int64   `json:"E"`
		Symbol        string  `json:"s"`
		TradeID       int64   `json:"t"`
		Price         float64 `json:"p,string"`
		Quantity      float64 `json:"q,string"`
		TradeTime     int64   `json:"T"`
		Buyer         bool    `json:"m"`
		Igrore        bool    `json:"M"`
		FirstUpdateID int64   `json:"U"`
		FinalUpdateID int64   `json:"u"`
	} `json:"data"`
}

// MessageDepht ...
type MessageDepht struct {
	Stream string `json:"stream"`
	Data   struct {
		EventType     string          `json:"e"`
		EventTime     int64           `json:"E"`
		Symbol        string          `json:"s"`
		FirstUpdateID int             `json:"U"`
		FinalUpdateID int             `json:"u"`
		Bids          [][]interface{} `json:"b"`
		Asks          [][]interface{} `json:"a"`
	} `json:"data"`
}
