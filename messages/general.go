package messages

import (
	"time"
)

// Heartbeat Exchange heartbeat message
type Heartbeat struct {
	Type        string `json:"type"`
	Sequence    int64  `json:"sequence"`
	LastTradeID int64  `json:"last_trade_id"`
	ProductID   string `json:"product_id"`
	Time        string `json:"time"`
}

// ParseTimestamp parse utc formatted timestamp to unix nanos
func ParseTimestamp(ts string) (int64, error) {
	t, e := time.Parse(time.RFC3339, ts)
	return t.UnixNano(), e
}
