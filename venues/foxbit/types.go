package foxbit

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
	InstrumentID     int    `json:"InstrumentId,omitempty"`
}

// Stream ...
type Stream struct {
	Streaming [][]interface{} `json:"streaming"`
}
