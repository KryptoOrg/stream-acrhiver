package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

// SubscriptionMessage sent to the exchange
type SubscriptionMessage struct {
	Type       string   `json:"type"`
	ProductIds []string `json:"product_ids"`
	Channels   []string `json:"channels"`
}

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

func main() {
	addr := "wss://ws-feed.pro.coinbase.com"
	// filename := "/tmp/archive.data"

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	log.Printf("connecting to %s", addr)

	connection, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer connection.Close()

	done := make(chan struct{})

	subs := SubscriptionMessage{Type: "subscribe", ProductIds: []string{"BTC-USD"}, Channels: []string{"full", "heartbeat"}}
	subsJSON, err := json.Marshal(&subs)

	if err != nil {
		log.Fatal("Json marshalling failed")
	}

	err = connection.WriteMessage(websocket.TextMessage, []byte(string(subsJSON)))

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer close(done)
		for {
			_, message, err := connection.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			var messageJSON map[string]interface{}
			json.Unmarshal([]byte(message), &messageJSON)
			log.Println(messageJSON)
			// messageType := messageJSON["type"]
			// if messageType == "snapshot" {
			// 	var snapshot Snapshot
			// 	json.Unmarshal([]byte(message), &snapshot)
			// 	// log.Println(snapshot)
			// } else if messageType == "l2update" {
			// 	var incremental L2Update
			// 	json.Unmarshal([]byte(message), &incremental)
			// 	log.Println(incremental)
			// }
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
