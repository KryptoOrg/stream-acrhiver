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

func write(bb []byte, f *os.File) {
	_, err := f.Write(bb)
	messages.Check(err)
}

func main() {
	addr := "wss://ws-feed.pro.coinbase.com"

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	log.Printf("connecting to %s", addr)

	connection, _, err := websocket.DefaultDialer.Dial(addr, nil)
	messages.Check(err)
	defer connection.Close()

	done := make(chan struct{})

	subs := messages.SubscriptionMessage{Type: "subscribe", ProductIds: []string{"BTC-USD"}, Channels: []string{"full", "heartbeat"}}
	subsJSON, err := json.Marshal(&subs)
	messages.Check(err)

	err = connection.WriteMessage(websocket.TextMessage, []byte(string(subsJSON)))
	messages.Check(err)

	go func() {
		defer close(done)
		file, err := os.Create("/tmp/coinbase_dump.data")
		messages.Check(err)

		defer file.Close()

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
				log.Printf("Received: %v\n", receivedJSON)
				write(messages.ConvertReceived(&receivedJSON), file)
			case "open":
				var openJSON messages.Open
				json.Unmarshal([]byte(message), &openJSON)
				log.Printf("Open: %v\n", openJSON)
				write(messages.ConvertOpen(&openJSON), file)
			case "done":
				var doneJSON messages.Done
				json.Unmarshal([]byte(message), &doneJSON)
				log.Printf("Done: %v\n", doneJSON)
				write(messages.ConvertDone(&doneJSON), file)
			case "match":
				var matchJSON messages.Match
				json.Unmarshal([]byte(message), &matchJSON)
				log.Printf("Match: %v\n", matchJSON)
				write(messages.ConvertMatch(&matchJSON), file)
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
