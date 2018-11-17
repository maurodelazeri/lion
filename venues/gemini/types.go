package gemini

// Message ...
type Message struct {
	HiveTable string `json:"hivetable,omitempty"`
	Type      string `json:"type,omitempty"`
}

// MessageChannel ...
type MessageChannel struct {
	Name       string   `json:"name"`
	ProductIDs []string `json:"product_ids"`
}
