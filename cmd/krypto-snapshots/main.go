package main

import (
	"encoding/json"
	"github.com/krypto-org/krypto-archiver/messages"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("https://api-public.sandbox.pro.coinbase.com/products/BTC-USD/book?level=2")
	messages.Check(err)
	responseData, err := ioutil.ReadAll(response.Body)
	messages.Check(err)

	var messageJSON map[string]interface{}
	messages.Check(json.Unmarshal(responseData, &messageJSON))
	bytes := messages.CreateL2Snapshot("BTC-USD", messageJSON)
	log.Info(bytes)
}
