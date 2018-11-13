package okex

// MessageTrade ...
type MessageTrade struct {
	Channel string     `json:"channel,omitempty"`
	Data    [][]string `json:"data,omitempty"`
	Binary  int        `json:"binary,omitempty"`
}

// MessageBook ...
type MessageBook struct {
	Binary  int    `json:"binary,omitempty"`
	Channel string `json:"channel,omitempty"`
	Data    struct {
		Asks      [][]string `json:"asks,omitempty"`
		Bids      [][]string `json:"bids,omitempty"`
		Timestamp int64      `json:"timestamp,omitempty"`
	} `json:"data"`
}

// MessageChannel ...
type MessageChannel struct {
	Binary  int    `json:"binary,omitempty"`
	Channel string `json:"channel,omitempty"`
	Event   string `json:"event,omitempty"`
	Data    struct {
		Result  bool   `json:"result,omitempty"`
		Channel string `json:"channel,omitempty"`
	} `json:"data,omitempty"`
}
