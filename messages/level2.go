package messages

// Snapshot received from the exchange
type Snapshot struct {
	Type      string     `json:"type"`
	ProductID string     `json:"product_id"`
	Bids      [][]string `json:"bids"`
	Asks      [][]string `json:"asks"`
}

// L2Update received from the exchange
type L2Update struct {
	Type      string     `json:"type"`
	ProductID string     `json:"product_id"`
	Time      string     `json:"time"`
	Changes   [][]string `json:"changes"`
}

type PriceLevel struct {
	Price    float64
	Quantity float64
}

type Level2Snapshot struct {
	Sequence int64
	Bids     []PriceLevel
	Asks     []PriceLevel
}
