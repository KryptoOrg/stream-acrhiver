package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/krypto-org/krypto-archiver/messages"
)

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

	subs := messages.SubscriptionMessage{Type: "subscribe", ProductIds: []string{"BTC-USD"}, Channels: []string{"full", "heartbeat"}}
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
			messageType := messageJSON["type"]
			switch messageType {
			case "received":
				var receivedJSON messages.Received
				json.Unmarshal([]byte(message), &receivedJSON)
				serialized := messages.ConvertReceived(&receivedJSON)
				log.Printf("Received: %s\n", serialized)
			case "open":
				var openJSON messages.Open
				json.Unmarshal([]byte(message), &openJSON)
				serialized := messages.ConvertOpen(&openJSON)
				log.Printf("Open: %s\n", serialized)
			case "done":
				var doneJSON messages.Done
				json.Unmarshal([]byte(message), &doneJSON)
			case "match":
				var matchJSON messages.Match
				json.Unmarshal([]byte(message), &matchJSON)
			case "change":
				var changeJSON messages.Change
				json.Unmarshal([]byte(message), &changeJSON)
			case "activate":
				var activateJSON messages.Activate
				json.Unmarshal([]byte(message), &activateJSON)
			case "heartbeat":
				var heartbeatJSON messages.Heartbeat
				json.Unmarshal([]byte(message), &heartbeatJSON)
			case "subscriptions":
				log.Printf("Subscribed! %s", messageJSON)
			default:
				log.Fatalf("Received unknown messageType : %s\n", messageType)
			}
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
