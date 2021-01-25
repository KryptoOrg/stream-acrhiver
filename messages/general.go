package messages

// Heartbeat Exchange heartbeat message
type Heartbeat struct {
	Type        string `json:"type"`
	Sequence    int64  `json:"sequence"`
	LastTradeID int64  `json:"last_trade_id"`
	ProductID   string `json:"product_id"`
	Time        string `json:"time"`
}
