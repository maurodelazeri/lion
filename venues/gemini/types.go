package gemini

// Message ...
type Message struct {
	Type           string `json:"type,omitempty"`
	EventID        int64  `json:"eventId,omitempty"`
	Timestamp      int    `json:"timestamp,omitempty"`
	Timestampms    int64  `json:"timestampms,omitempty"`
	SocketSequence int    `json:"socket_sequence,omitempty"`
	Events         []struct {
		Type      string `json:"type"`
		Tid       int64  `json:"tid,omitempty"`
		Price     string `json:"price,omitempty"`
		Amount    string `json:"amount,omitempty"`
		MakerSide string `json:"makerSide,omitempty"`
		Side      string `json:"side,omitempty"`
		Remaining string `json:"remaining,omitempty"`
		Delta     string `json:"delta,omitempty"`
		Reason    string `json:"reason,omitempty"`
	} `json:"events,omitempty"`
}

// PingPong ...
type PingPong struct {
	Ping int64 `json:"ping"`
}
