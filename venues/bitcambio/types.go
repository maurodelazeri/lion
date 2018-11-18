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
	SecurityStatusReqID     int      `json:"SecurityStatusReqID,omitempty"`
}

// Message ...
type Message struct {
	M int    `json:"m"`
	I int    `json:"i"`
	N string `json:"n"`
	O string `json:"o"`
}
