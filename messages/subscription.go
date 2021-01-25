package messages

// SubscriptionMessage sent to the exchange
type SubscriptionMessage struct {
	Type       string   `json:"type"`
	ProductIds []string `json:"product_ids"`
	Channels   []string `json:"channels"`
}
