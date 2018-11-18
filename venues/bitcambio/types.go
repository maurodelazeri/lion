package bitcambio

// MessageChannel ...
type MessageChannel struct {
	MsgType                 string   `json:"MsgType,omitempty"`
	MDReqID                 int64    `json:"MDReqID,omitempty"`
	SubscriptionRequestType string   `json:"SubscriptionRequestType,omitempty"`
	MarketDepth             string   `json:"MarketDepth,omitempty"`
	MDUpdateType            string   `json:"MDUpdateType,omitempty"`
	MDEntryTypes            []string `json:"MDEntryTypes,omitempty"`
	Instruments             []string `json:"Instruments,omitempty"`
}

// Message ...
type Message struct {
	MDReqID   int    `json:"MDReqID,omitempty"`
	Symbol    string `json:"Symbol,omitempty"`
	MDBkTyp   string `json:"MDBkTyp,omitempty"`
	MsgType   string `json:"MsgType,omitempty"`
	MDFullGrp []struct {
		MDUpdateAction    string  `json:"MDUpdateAction,omitempty"`
		MDEntryPositionNo int     `json:"MDEntryPositionNo,omitempty"`
		MDEntrySize       float64 `json:"MDEntrySize,omitempty"`
		MDEntryPx         float64 `json:"MDEntryPx,omitempty"`
		MDEntryID         int     `json:"MDEntryID,omitempty"`
		MDEntryTime       string  `json:"MDEntryTime,omitempty"`
		MDEntryDate       string  `json:"MDEntryDate,omitempty"`
		UserID            int     `json:"UserID,omitempty"`
		OrderID           int     `json:"OrderID,omitempty"`
		MDEntryType       string  `json:"MDEntryType,omitempty"`
		Broker            string  `json:"Broker,omitempty"`
		MDEntryBuyerID    int     `json:"MDEntryBuyerID,omitempty"`
		Side              string  `json:"Side,omitempty"`
		TradeID           int     `json:"TradeID,omitempty"`
		Symbol            string  `json:"Symbol,omitempty"`
		MDEntrySellerID   int     `json:"MDEntrySellerID,omitempty"`
		SecondaryOrderID  int     `json:"SecondaryOrderID,omitempty"`
		BRL               int64   `json:"BRL,omitempty"`
		BTC               int64   `json:"BTC,omitempty"`
	} `json:"MDFullGrp,omitempty"`
	MarketDepth int `json:"MarketDepth,omitempty"`
}
