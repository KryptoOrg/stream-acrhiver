package messages

// Received Exchange received message:
type Received struct {
	Type      string `json:"type"`
	Time      string `json:"time"`
	ProductID string `json:"product_id"`
	Sequence  int64  `json:"sequence"`
	OrderID   string `json:"order_id"`
	Size      string `json:"size"`
	Price     string `json:"price"`
	Side      string `json:"side"`
	OrderType string `json:"order_type"`
	Funds     string `json:"funds"`
}

// Open Exchange open order message
type Open struct {
	Type          string `json:"type"`
	Time          string `json:"time"`
	ProductID     string `json:"product_id"`
	Sequence      int64  `json:"sequence"`
	OrderID       string `json:"order_id"`
	Price         string `json:"price"`
	RemainingSize string `json:"remaining_size"`
	Side          string `json:"side"`
}

// Done Exchange order done message
type Done struct {
	Type          string `json:"type"`
	Time          string `json:"time"`
	ProductID     string `json:"product_id"`
	Sequence      int64  `json:"sequence"`
	OrderID       string `json:"order_id"`
	Price         string `json:"price"`
	Side          string `json:"side"`
	Reason        string `json:"reason"`
	RemainingSize string `json:"remaining_size"`
}

// Match Exchange order match message
type Match struct {
	Type         string `json:"type"`
	TradeID      int64  `json:"trade_id"`
	Time         string `json:"time"`
	ProductID    string `json:"product_id"`
	Sequence     int64  `json:"sequence"`
	MakerOrderID string `json:"maker_order_id"`
	TakerOrderID string `json:"taker_order_id"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	Side         string `json:"side"`
}

// Change Exchange order change message
type Change struct {
	Type      string `json:"type"`
	Time      string `json:"time"`
	ProductID string `json:"product_id"`
	Sequence  int64  `json:"sequence"`
	OrderID   string `json:"order_id"`
	NewSize   string `json:"new_size"`
	OldSize   string `json:"old_size"`
	Price     string `json:"price"`
	Side      string `json:"side"`
}

// Activate Exchange activate ordre message
type Activate struct {
	Type      string `json:"type"`
	ProductID string `json:"product_id"`
	Timestamp string `json:"timestamp"`
	UserID    string `json:"user_id"`
	ProfileID string `json:"profile_id"`
	OrderID   string `json:"order_id"`
	StopType  string `json:"stop_type"`
	Side      string `json:"side"`
	StopPrice string `json:"stop_price"`
	Size      string `json:"size"`
	Funds     string `json:"funds"`
	Private   bool   `json:"private"`
}
