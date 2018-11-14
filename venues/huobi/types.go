package huobi

// MessageChannel ...
type MessageChannel struct {
	Sub string `json:"sub"`
	ID  string `json:"id"`
}

// Message ...
type Message struct {
	Ch   string `json:"ch,omitempty"`
	Ts   int64  `json:"ts,omitempty"`
	Tick struct {
		Ts      int64       `json:"ts,omitempty"`
		Bids    [][]float64 `json:"bids,omitempty"`
		Asks    [][]float64 `json:"asks,omitempty"`
		Version int64       `json:"version,omitempty"`
		Data    []struct {
			Amount    float64 `json:"amount,omitempty"`
			Ts        int64   `json:"ts,omitempty"`
			Price     float64 `json:"price,omitempty"`
			Direction string  `json:"direction,omitempty"`
		} `json:"data,omitempty"`
	} `json:"tick"`
}
