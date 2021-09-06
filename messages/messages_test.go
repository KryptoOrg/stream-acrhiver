package messages

import (
	"testing"

	"github.com/krypto-org/krypto-archiver/serialization"
	"github.com/stretchr/testify/assert"
)

func TestParseTimestamp(t *testing.T) {
	tests := []struct {
		name      string
		timestamp string
		expected  int64
	}{
		{"utc nanoseconds", "2014-11-07T08:19:28.464459318Z", 1415348368464459318},
		{"utc microseconds", "2014-11-07T08:19:28.464459Z", 1415348368464459000},
		{"utc microseconds 5 letters", "2014-11-07T08:19:28.46445Z", 1415348368464450000},
		{"utc microseconds 4 letters", "2014-11-07T08:19:28.4648Z", 1415348368464800000},
		{"utc milliseocnds", "2020-11-07T08:19:28.464Z", 1604737168464000000},
		{"utc seocnds", "2020-11-07T08:19:28Z", 1604737168000000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tv, _ := ParseTimestamp(tt.timestamp)
			if tv != tt.expected {
				t.Errorf("ParseTimestamp(%v) = %v != expected (%v)", tt.timestamp, tv, tt.expected)
			}
		})
	}
}

func TestConvertReceived(t *testing.T) {

	received := Received{
		Type:      "received",
		Time:      "2014-11-07T08:19:27.028459Z",
		ProductID: "BTC-USD",
		Sequence:  10,
		OrderID:   "d50ec984-77a8-460a-b958-66f114b0de9b",
		Size:      "1.34",
		Price:     "502.1",
		Side:      "buy",
		OrderType: "limit",
	}

	orderUpdate := serialization.GetRootAsOrderUpdate(ConvertReceived(&received), 0)

	assert.Equal(t, orderUpdate.OrderUpdateType(), serialization.OrderUpdateTypeRECEIVED)
	assert.Equal(t, orderUpdate.Timestamp(), int64(1415348367028459000))
	assert.Equal(t, string(orderUpdate.ProductId()), "BTC-USD")
	assert.Equal(t, orderUpdate.Sequence(), int64(10))
	assert.Equal(t, string(orderUpdate.OrderId()), "d50ec984-77a8-460a-b958-66f114b0de9b")
	assert.Equal(t, orderUpdate.Size(), 1.34)
	assert.Equal(t, orderUpdate.Price(), 502.1)
	assert.Equal(t, orderUpdate.Side(), serialization.SideBUY)
	assert.Equal(t, orderUpdate.OrderType(), serialization.OrderTypeLIMIT)

	received = Received{
		Type:      "received",
		Time:      "2014-11-07T08:19:27.028459Z",
		ProductID: "BTC-USD",
		Sequence:  12,
		OrderID:   "d50ec984-77a8-460a-b958-66f114b0de9b",
		Funds:     "3000.234",
		Side:      "buy",
		OrderType: "market",
	}

	orderUpdate = serialization.GetRootAsOrderUpdate(ConvertReceived(&received), 0)

	assert.Equal(t, orderUpdate.OrderUpdateType(), serialization.OrderUpdateTypeRECEIVED)
	assert.Equal(t, orderUpdate.Timestamp(), int64(1415348367028459000))
	assert.Equal(t, string(orderUpdate.ProductId()), "BTC-USD")
	assert.Equal(t, orderUpdate.Sequence(), int64(12))
	assert.Equal(t, string(orderUpdate.OrderId()), "d50ec984-77a8-460a-b958-66f114b0de9b")
	assert.Equal(t, orderUpdate.Funds(), 3000.234)
	assert.Equal(t, orderUpdate.Side(), serialization.SideBUY)
	assert.Equal(t, orderUpdate.OrderType(), serialization.OrderTypeMARKET)
}

func TestConvertOpen(t *testing.T) {
	open := Open{
		Type:          "open",
		Time:          "2014-11-07T08:19:27.028459Z",
		ProductID:     "BTC-USD",
		Sequence:      10,
		OrderID:       "d50ec984-77a8-460a-b958-66f114b0de9b",
		Price:         "200.2",
		RemainingSize: "1.00",
		Side:          "sell",
	}

	orderUpdate := serialization.GetRootAsOrderUpdate(ConvertOpen(&open), 0)

	assert.Equal(t, orderUpdate.OrderUpdateType(), serialization.OrderUpdateTypeOPEN)
	assert.Equal(t, orderUpdate.Timestamp(), int64(1415348367028459000))
	assert.Equal(t, string(orderUpdate.ProductId()), "BTC-USD")
	assert.Equal(t, orderUpdate.Sequence(), int64(10))
	assert.Equal(t, string(orderUpdate.OrderId()), "d50ec984-77a8-460a-b958-66f114b0de9b")
	assert.Equal(t, orderUpdate.Price(), 200.2)
	assert.Equal(t, orderUpdate.RemainingSize(), 1.00)
	assert.Equal(t, orderUpdate.Side(), serialization.SideSELL)
}

func TestConvertDone(t *testing.T) {
	/*
	   {
	       "type": "done",
	       "time": ,
	       "product_id": "BTC-USD",
	       "sequence": 10,
	       "price": "200.2",
	       "order_id": ,
	       "reason": "filled", // or "canceled"
	       "side": "sell",
	       "remaining_size": "0"
	   }
	*/

	done := Done{
		Type:          "done",
		Time:          "2014-11-07T08:19:27.028459Z",
		ProductID:     "BTC-USD",
		Sequence:      10,
		Price:         "200.2",
		OrderID:       "d50ec984-77a8-460a-b958-66f114b0de9b",
		Reason:        "filled",
		Side:          "sell",
		RemainingSize: "0.1",
	}

	orderUpdate := serialization.GetRootAsOrderUpdate(ConvertDone(&done), 0)

	assert.Equal(t, orderUpdate.OrderUpdateType(), serialization.OrderUpdateTypeDONE)
	assert.Equal(t, orderUpdate.Timestamp(), int64(1415348367028459000))
	assert.Equal(t, string(orderUpdate.ProductId()), "BTC-USD")
	assert.Equal(t, orderUpdate.Sequence(), int64(10))
	assert.Equal(t, string(orderUpdate.OrderId()), "d50ec984-77a8-460a-b958-66f114b0de9b")
	assert.Equal(t, orderUpdate.Price(), 200.2)
	assert.Equal(t, string(orderUpdate.Reason()), "filled")
	assert.Equal(t, orderUpdate.RemainingSize(), 0.1)
	assert.Equal(t, orderUpdate.Side(), serialization.SideSELL)

}

func TestConvertMatch(t *testing.T) {

	match := Match{
		Type:         "match",
		Time:         "2014-11-07T08:19:27.028459Z",
		TradeID:      10,
		Sequence:     50,
		MakerOrderID: "ac928c66-ca53-498f-9c13-a110027a60e8",
		TakerOrderID: "132fb6ae-456b-4654-b4e0-d681ac05cea1",
		ProductID:    "BTC-USD",
		Size:         "5.23512",
		Price:        "400.23",
		Side:         "sell",
	}

	orderUpdate := serialization.GetRootAsOrderUpdate(ConvertMatch(&match), 0)

	assert.Equal(t, orderUpdate.OrderUpdateType(), serialization.OrderUpdateTypeMATCH)
	assert.Equal(t, orderUpdate.Timestamp(), int64(1415348367028459000))
	assert.Equal(t, string(orderUpdate.ProductId()), "BTC-USD")
	assert.Equal(t, orderUpdate.Sequence(), int64(50))
	assert.Equal(t, orderUpdate.TradeId(), int64(10))
	assert.Equal(t, string(orderUpdate.MakerOrderId()), "ac928c66-ca53-498f-9c13-a110027a60e8")
	assert.Equal(t, string(orderUpdate.TakerOrderId()), "132fb6ae-456b-4654-b4e0-d681ac05cea1")
	assert.Equal(t, orderUpdate.Size(), 5.23512)
	assert.Equal(t, orderUpdate.Price(), 400.23)
	assert.Equal(t, orderUpdate.Side(), serialization.SideSELL)

}
