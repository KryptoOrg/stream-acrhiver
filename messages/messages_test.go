package messages

import (
	"testing"

	"github.com/krypto-org/krypto-archiver/serialization"
	assert "github.com/stretchr/testify/assert"
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
