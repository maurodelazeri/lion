package poloniex

// Message ...
type Message struct {
	HiveTable string `json:"hivetable,omitempty"`
	Type      string `json:"type,omitempty"`
}

// MessageChannel ...
// Subscription stores a subscription
type MessageChannel struct {
	Command string `json:"command"`
	Channel string `json:"channel"`
}
