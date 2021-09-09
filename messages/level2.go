package messages

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/krypto-org/krypto-archiver/serialization"
	"strconv"
)

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

func CreateL2Snapshot(symbol string, json map[string]interface{}) []byte {
	builder := flatbuffers.NewBuilder(1024)
	symbolOffset := builder.CreateString(symbol)
	sequence := int64(json["sequence"].(float64))

	bids := json["bids"].([]interface{})
	asks := json["asks"].([]interface{})

	serialization.MarketByLevelSnapshotStartBidsVector(builder, len(bids))
	for _, bid := range bids {
		split := bid.([]interface{})
		price, err := strconv.ParseFloat(split[0].(string), 10)
		Check(err)
		quantity, err := strconv.ParseFloat(split[1].(string), 10)
		Check(err)
		serialization.CreatePriceLevel(builder, price, quantity)
	}
	bidsOffset := builder.EndVector(len(bids))

	serialization.MarketByLevelSnapshotStartAsksVector(builder, len(asks))
	for _, ask := range asks {
		split := ask.([]interface{})
		price, err := strconv.ParseFloat(split[0].(string), 10)
		Check(err)
		quantity, err := strconv.ParseFloat(split[1].(string), 10)
		Check(err)
		serialization.CreatePriceLevel(builder, price, quantity)
	}
	asksOffset := builder.EndVector(len(asks))

	serialization.MarketByLevelSnapshotStart(builder)
	serialization.MarketByLevelSnapshotAddSequence(builder, sequence)
	serialization.MarketByLevelSnapshotAddSymbol(builder, symbolOffset)
	serialization.MarketByLevelSnapshotAddBids(builder, bidsOffset)
	serialization.MarketByLevelSnapshotAddAsks(builder, asksOffset)
	snapshot := serialization.MarketByLevelSnapshotEnd(builder)

	builder.Finish(snapshot)
	return builder.FinishedBytes()
}
