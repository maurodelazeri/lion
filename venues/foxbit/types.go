package foxbit

// Message ...
type Message struct {
	//Time          time.Time        `json:"time,string,omitempty"`
	HiveTable     string           `json:"hivetable,omitempty"`
	Type          string           `json:"type,omitempty"`
	ProductID     string           `json:"product_id,omitempty"`
	ProductIDs    []string         `json:"product_ids,omitempty"`
	TradeID       int64            `json:"trade_id,number,omitempty"`
	OrderID       string           `json:"order_id,omitempty"`
	Sequence      int64            `json:"sequence,number,omitempty"`
	MakerOrderID  string           `json:"maker_order_id,omitempty"`
	TakerOrderID  string           `json:"taker_order_id,omitempty"`
	RemainingSize float64          `json:"remaining_size,string,omitempty"`
	NewSize       float64          `json:"new_size,string,omitempty"`
	OldSize       float64          `json:"old_size,string,omitempty"`
	Size          float64          `json:"size,string,omitempty"`
	Price         float64          `json:"price,string,omitempty"`
	Side          string           `json:"side,omitempty"`
	Reason        string           `json:"reason,omitempty"`
	OrderType     string           `json:"order_type,omitempty"`
	Funds         float64          `json:"funds,string,omitempty"`
	NewFunds      float64          `json:"new_funds,string,omitempty"`
	OldFunds      float64          `json:"old_funds,string,omitempty"`
	Message       string           `json:"message,omitempty"`
	Bids          [][]string       `json:"bids,omitempty"`
	Asks          [][]string       `json:"asks,omitempty"`
	Changes       [][]string       `json:"changes,omitempty"`
	LastSize      float64          `json:"last_size,string,omitempty"`
	BestBid       float64          `json:"best_bid,string,omitempty"`
	BestAsk       float64          `json:"best_ask,string,omitempty"`
	Channels      []MessageChannel `json:"channels,omitempty"`
	UserID        string           `json:"user_id,omitempty"`
	ProfileID     string           `json:"profile_id,omitempty"`
	Open24H       float64          `json:"open_24h,string,omitempty"`
	Volume24H     float64          `json:"volume_24h,string,omitempty"`
	Low24H        float64          `json:"low_24h,string,omitempty"`
	High24H       float64          `json:"high_24h,string,omitempty"`
	Volume30D     float64          `json:"volume_30d,string,omitempty"`
}

// MessageChannel ...
// type MessageChannel struct {
// 	OMSID            int `json:"OMSId"`
// 	InstrumentID     int `json:"InstrumentId"`
// 	IncludeLastCount int `json:"IncludeLastCount"`
// }

// MessageChannel ...
type MessageChannel struct {
	M int    `json:"m"`
	I int    `json:"i"`
	N string `json:"n"`
	O string `json:"o"`
}

// Payload ...
type Payload struct {
	OMSID            int    `json:"OMSId,omitempty"`
	Symbol           string `json:"Symbol,omitempty"`
	Depth            int    `json:"Depth,omitempty"`
	IncludeLastCount int    `json:"IncludeLastCount,omitempty"`
	InstrumentId     int    `json:"InstrumentId,omitempty"`
}