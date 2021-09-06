package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
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
	// TODO: Split the data in multiple files
	// TODO: Configuration for file and input details

	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05.999"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)

	addr := "wss://ws-feed.pro.coinbase.com"

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	log.Infof("connecting to %s", addr)

	connection, _, err := websocket.DefaultDialer.Dial(addr, nil)
	messages.Check(err)
	defer func(connection *websocket.Conn) {
		err := connection.Close()
		if err != nil {
			log.Errorf("Error while closing file %s\n", err)
		}
	}(connection)

	done := make(chan struct{})

	subs := messages.SubscriptionMessage{Type: "subscribe", ProductIds: []string{"BTC-USD"}, Channels: []string{"full", "heartbeat"}}
	subsJSON, err := json.Marshal(&subs)
	messages.Check(err)

	err = connection.WriteMessage(websocket.TextMessage, subsJSON)
	messages.Check(err)

	go func() {
		defer close(done)
		file, err := os.Create("/tmp/coinbase_dump.data")
		messages.Check(err)

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Errorf("Error while closing file %s\n", err)
			}
		}(file)

		for {
			_, message, err := connection.ReadMessage()
			if err != nil {
				log.Warning("read: ", err)
				return
			}
			var messageJSON map[string]interface{}
			messages.Check(json.Unmarshal(message, &messageJSON))
			messageType := messageJSON["type"]
			switch messageType {
			case "received":
				var receivedJSON messages.Received
				messages.Check(json.Unmarshal(message, &receivedJSON))
				log.Infof("Received: %v\n", receivedJSON)
				write(messages.ConvertReceived(&receivedJSON), file)
			case "open":
				var openJSON messages.Open
				messages.Check(json.Unmarshal(message, &openJSON))
				log.Infof("Open: %v\n", openJSON)
				write(messages.ConvertOpen(&openJSON), file)
			case "done":
				var doneJSON messages.Done
				messages.Check(json.Unmarshal(message, &doneJSON))
				log.Infof("Done: %v\n", doneJSON)
				write(messages.ConvertDone(&doneJSON), file)
			case "match":
				var matchJSON messages.Match
				messages.Check(json.Unmarshal(message, &matchJSON))
				log.Infof("Match: %v\n", matchJSON)
				write(messages.ConvertMatch(&matchJSON), file)
			case "change":
				var changeJSON messages.Change
				messages.Check(json.Unmarshal(message, &changeJSON))
			case "activate":
				var activateJSON messages.Activate
				messages.Check(json.Unmarshal(message, &activateJSON))
			case "heartbeat":
				var heartbeatJSON messages.Heartbeat
				messages.Check(json.Unmarshal(message, &heartbeatJSON))
			case "subscriptions":
				log.Infof("Subscribed! %s", messageJSON)
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
			log.Info("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Error("write close:", err)
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
