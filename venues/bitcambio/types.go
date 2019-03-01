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
	MDReqID  int    `json:"MDReqID,omitempty"`
	Symbol   string `json:"Symbol,omitempty"`
	MDBkTyp  string `json:"MDBkTyp,omitempty"`
	MsgType  string `json:"MsgType,omitempty"`
	MDIncGrp []struct {
		OrderID           int64  `json:"OrderID,omitempty"`
		MDEntryPx         int64  `json:"MDEntryPx,omitempty"`
		MDUpdateAction    string `json:"MDUpdateAction,omitempty"`
		MDEntryTime       string `json:"MDEntryTime,omitempty"`
		Symbol            string `json:"Symbol,omitempty"`
		UserID            int64  `json:"UserID"`
		Broker            string `json:"Broker,omitempty"`
		MDEntryType       string `json:"MDEntryType,omitempty"`
		MDEntryPositionNo int64  `json:"MDEntryPositionNo,omitempty"`
		MDEntrySize       int64  `json:"MDEntrySize,omitempty"`
		Side              string `json:"Side,omitempty"`
		MDEntryID         int64  `json:"MDEntryID,omitempty"`
		MDEntryDate       string `json:"MDEntryDate,omitempty"`
	} `json:"MDIncGrp,omitempty"`
	MDFullGrp []struct {
		MDUpdateAction    string `json:"MDUpdateAction,omitempty"`
		MDEntryPositionNo int64  `json:"MDEntryPositionNo,omitempty"`
		MDEntrySize       int64  `json:"MDEntrySize,omitempty"`
		MDEntryPx         int64  `json:"MDEntryPx,omitempty"`
		MDEntryID         int64  `json:"MDEntryID,omitempty"`
		MDEntryTime       string `json:"MDEntryTime,omitempty"`
		MDEntryDate       string `json:"MDEntryDate,omitempty"`
		UserID            int64  `json:"UserID,omitempty"`
		OrderID           int64  `json:"OrderID,omitempty"`
		MDEntryType       string `json:"MDEntryType,omitempty"`
		Broker            string `json:"Broker,omitempty"`
		MDEntryBuyerID    int64  `json:"MDEntryBuyerID,omitempty"`
		Side              string `json:"Side,omitempty"`
		TradeID           int64  `json:"TradeID,omitempty"`
		Symbol            string `json:"Symbol,omitempty"`
		MDEntrySellerID   int64  `json:"MDEntrySellerID,omitempty"`
		SecondaryOrderID  int64  `json:"SecondaryOrderID,omitempty"`
		BRL               int64  `json:"BRL,omitempty"`
		BTC               int64  `json:"BTC,omitempty"`
	} `json:"MDFullGrp,omitempty"`
	MarketDepth int `json:"MarketDepth,omitempty"`
}

// BookItem ...
type BookItem struct {
	Price     float64
	Volume    float64
	MDEntryID int64
}
