package messages

import (
	"strconv"

	flatbuffers "github.com/google/flatbuffers/go"
	serialization "github.com/krypto-org/krypto-archiver/serialization"
)

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

func parseSide(side string) serialization.Side {
	if side == "buy" {
		return serialization.SideBUY
	} else if side == "sell" {
		return serialization.SideSELL
	}
	return serialization.SideUNKNOWN
}

func parseOrderType(orderType string) serialization.OrderType {
	if orderType == "limit" {
		return serialization.OrderTypeLIMIT
	} else if orderType == "market" {
		return serialization.OrderTypeMARKET
	}
	return serialization.OrderTypeUNKNOWN
}

func parseCheckedFloat(value string) float64 {
	var parsed float64 = 0

	if value != "" {
		fundsParsed, e := strconv.ParseFloat(value, 64)
		Check(e)
		parsed = fundsParsed
	}

	return parsed
}

// ConvertReceived received to flatbuffer type
func ConvertReceived(message *Received) []byte {

	builder := flatbuffers.NewBuilder(1024)
	productID := builder.CreateString(message.ProductID)
	orderID := builder.CreateString(message.OrderID)

	size := parseCheckedFloat(message.Size)

	price := parseCheckedFloat(message.Price)

	side := parseSide(message.Side)
	orderType := parseOrderType(message.OrderType)

	timestamp, e := ParseTimestamp(message.Time)
	Check(e)

	funds := parseCheckedFloat(message.Funds)

	serialization.OrderUpdateStart(builder)
	serialization.OrderUpdateAddTimestamp(builder, timestamp)
	serialization.OrderUpdateAddOrderUpdateType(builder, serialization.OrderUpdateTypeRECEIVED)
	serialization.OrderUpdateAddSequence(builder, message.Sequence)
	serialization.OrderUpdateAddSize(builder, size)
	serialization.OrderUpdateAddPrice(builder, price)
	serialization.OrderUpdateAddSide(builder, side)
	serialization.OrderUpdateAddOrderType(builder, orderType)
	serialization.OrderUpdateAddProductId(builder, productID)
	serialization.OrderUpdateAddOrderId(builder, orderID)
	serialization.OrderUpdateAddFunds(builder, funds)
	orderUpdate := serialization.OrderUpdateEnd(builder)

	builder.Finish(orderUpdate)
	return builder.FinishedBytes()
}

// ConvertOpen open to flatbuffer type
func ConvertOpen(message *Open) []byte {

	builder := flatbuffers.NewBuilder(1024)
	productID := builder.CreateString(message.ProductID)
	orderID := builder.CreateString(message.OrderID)

	remainingSize, e := strconv.ParseFloat(message.RemainingSize, 64)
	Check(e)

	price, e := strconv.ParseFloat(message.Price, 64)
	Check(e)

	side := parseSide(message.Side)

	timestamp, e := ParseTimestamp(message.Time)
	Check(e)

	serialization.OrderUpdateStart(builder)
	serialization.OrderUpdateAddTimestamp(builder, timestamp)
	serialization.OrderUpdateAddOrderUpdateType(builder, serialization.OrderUpdateTypeOPEN)
	serialization.OrderUpdateAddSequence(builder, message.Sequence)
	serialization.OrderUpdateAddPrice(builder, price)
	serialization.OrderUpdateAddSide(builder, side)
	serialization.OrderUpdateAddRemainingSize(builder, remainingSize)
	serialization.OrderUpdateAddProductId(builder, productID)
	serialization.OrderUpdateAddOrderId(builder, orderID)
	orderUpdate := serialization.OrderUpdateEnd(builder)

	builder.Finish(orderUpdate)
	return builder.FinishedBytes()
}

// ConvertDone done to flatbuffer type
func ConvertDone(message *Done) []byte {

	builder := flatbuffers.NewBuilder(1024)
	productID := builder.CreateString(message.ProductID)
	orderID := builder.CreateString(message.OrderID)
	reason := builder.CreateString(message.Reason)

	remainingSize := parseCheckedFloat(message.RemainingSize)
	price := parseCheckedFloat(message.Price)
	side := parseSide(message.Side)

	timestamp, e := ParseTimestamp(message.Time)
	Check(e)

	serialization.OrderUpdateStart(builder)
	serialization.OrderUpdateAddTimestamp(builder, timestamp)
	serialization.OrderUpdateAddOrderUpdateType(builder, serialization.OrderUpdateTypeDONE)
	serialization.OrderUpdateAddSequence(builder, message.Sequence)
	serialization.OrderUpdateAddPrice(builder, price)
	serialization.OrderUpdateAddSide(builder, side)
	serialization.OrderUpdateAddRemainingSize(builder, remainingSize)
	serialization.OrderUpdateAddProductId(builder, productID)
	serialization.OrderUpdateAddOrderId(builder, orderID)
	serialization.OrderUpdateAddReason(builder, reason)
	orderUpdate := serialization.OrderUpdateEnd(builder)

	builder.Finish(orderUpdate)
	return builder.FinishedBytes()
}

// ConvertMatch match to flatbuffer type
func ConvertMatch(message *Match) []byte {

	builder := flatbuffers.NewBuilder(1024)
	productID := builder.CreateString(message.ProductID)
	makerOrderID := builder.CreateString(message.MakerOrderID)
	takerOrderID := builder.CreateString(message.TakerOrderID)

	size, e := strconv.ParseFloat(message.Size, 64)
	Check(e)

	price, e := strconv.ParseFloat(message.Price, 64)
	Check(e)

	side := parseSide(message.Side)

	timestamp, e := ParseTimestamp(message.Time)
	Check(e)

	serialization.OrderUpdateStart(builder)
	serialization.OrderUpdateAddTimestamp(builder, timestamp)
	serialization.OrderUpdateAddOrderUpdateType(builder, serialization.OrderUpdateTypeMATCH)
	serialization.OrderUpdateAddTradeId(builder, message.TradeID)
	serialization.OrderUpdateAddSequence(builder, message.Sequence)
	serialization.OrderUpdateAddPrice(builder, price)
	serialization.OrderUpdateAddSide(builder, side)
	serialization.OrderUpdateAddSize(builder, size)
	serialization.OrderUpdateAddProductId(builder, productID)
	serialization.OrderUpdateAddMakerOrderId(builder, makerOrderID)
	serialization.OrderUpdateAddTakerOrderId(builder, takerOrderID)
	orderUpdate := serialization.OrderUpdateEnd(builder)

	builder.Finish(orderUpdate)
	return builder.FinishedBytes()
}
