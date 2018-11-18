package bitcambio

// MessageChannel ...
type MessageChannel struct {
	MsgType                 string   `json:"MsgType"`
	MDReqID                 int      `json:"MDReqID"`
	SubscriptionRequestType string   `json:"SubscriptionRequestType"`
	MarketDepth             string   `json:"MarketDepth"`
	MDUpdateType            string   `json:"MDUpdateType"`
	MDEntryTypes            []string `json:"MDEntryTypes"`
	Instruments             []string `json:"Instruments"`
}

// Message ...
type Message struct {
	M int    `json:"m"`
	I int    `json:"i"`
	N string `json:"n"`
	O string `json:"o"`
}
