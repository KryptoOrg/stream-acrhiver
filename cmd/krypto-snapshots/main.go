package main

import (
	"encoding/json"
	"github.com/krypto-org/krypto-archiver/messages"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	response, err := http.Get("https://api-public.sandbox.pro.coinbase.com/products/BTC-USD/book?level=2")
	messages.Check(err)
	responseData, err := ioutil.ReadAll(response.Body)
	messages.Check(err)

	var messageJSON map[string]interface{}
	messages.Check(json.Unmarshal(responseData, &messageJSON))
	sequence := int64(messageJSON["sequence"].(float64))

	snapshot := messages.Level2Snapshot{
		Sequence: sequence,
	}

	for _, bid := range messageJSON["bids"].([]interface{}) {
		split := bid.([]interface{})
		price, err := strconv.ParseFloat(split[0].(string), 10)
		messages.Check(err)
		quantity, err := strconv.ParseFloat(split[1].(string), 10)
		messages.Check(err)
		snapshot.Bids = append(snapshot.Bids, messages.PriceLevel{Price: price, Quantity: quantity})
	}

	for _, ask := range messageJSON["asks"].([]interface{}) {
		split := ask.([]interface{})
		price, err := strconv.ParseFloat(split[0].(string), 10)
		messages.Check(err)
		quantity, err := strconv.ParseFloat(split[1].(string), 10)
		messages.Check(err)
		snapshot.Asks = append(snapshot.Asks, messages.PriceLevel{Price: price, Quantity: quantity})
	}

	log.Info(snapshot)
}
