package bitmex

// MessageChannel takes in subscription information
type MessageChannel struct {
	OP   string   `json:"op"`
	Args []string `json:"args"`
}

// Message ...
type Message struct {
	Table  string   `json:"table,omitempty"`
	Action string   `json:"action,omitempty"`
	Keys   []string `json:"keys,omitempty"`
	Types  struct {
		Symbol string `json:"symbol,omitempty"`
		ID     string `json:"id,omitempty"`
		Side   string `json:"side,omitempty"`
		Size   string `json:"size,omitempty"`
		Price  string `json:"price,omitempty"`
	} `json:"types,omitempty"`
	ForeignKeys struct {
		Symbol string `json:"symbol,omitempty"`
		Side   string `json:"side,omitempty"`
	} `json:"foreignKeys"`
	Attributes struct {
		Symbol string `json:"symbol,omitempty"`
		ID     string `json:"id,omitempty"`
	} `json:"attributes,omitempty"`
	Filter struct {
		Symbol string `json:"symbol,omitempty"`
	} `json:"filter,omitempty"`
	Data []struct {
		Symbol string  `json:"symbol,omitempty"`
		ID     int64   `json:"id,omitempty"`
		Side   string  `json:"side,omitempty"`
		Size   int     `json:"size,omitempty"`
		Price  float64 `json:"price,omitempty"`
	} `json:"data,omitempty"`
}

// BookItem ...
type BookItem struct {
	Price  float64
	Volume float64
}
