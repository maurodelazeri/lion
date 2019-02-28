package kraken

// Message ...
type Message struct {
}

// PingPong ...
type PingPong struct {
	Event string `json:"event"`
	Reqid int64  `json:"reqid"`
}

// SubscriptionBook ...
type SubscriptionBook struct {
	Event        string   `json:"event"`
	Pair         []string `json:"pair"`
	Subscription struct {
		Name  string `json:"name"`
		Depth int    `json:"depth"`
	} `json:"subscription"`
}

// SubscriptionTrade ...
type SubscriptionTrade struct {
	Event        string   `json:"event"`
	Pair         []string `json:"pair"`
	Subscription struct {
		Name string `json:"name"`
	} `json:"subscription"`
}

// Subscription ...
type Subscription struct {
	ChannelID    int    `json:"channelID"`
	Event        string `json:"event"`
	Pair         string `json:"pair"`
	Status       string `json:"status"`
	Subscription struct {
		Depth int    `json:"depth"`
		Name  string `json:"name"`
	} `json:"subscription"`
}
