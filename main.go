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
		fileHandler := messages.FileHandler{
			Frequency: int64(60 * time.Second),
			Filename:  "coinbase_dump.data",
			Directory: "/tmp/coinbase",
		}

		defer messages.Close(&fileHandler)

		for {
			err := messages.Update(&fileHandler)
			if err != nil {
				log.Error("read: ", err)
				return
			}
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
				write(messages.ConvertReceived(&receivedJSON), fileHandler.File)
			case "open":
				var openJSON messages.Open
				messages.Check(json.Unmarshal(message, &openJSON))
				write(messages.ConvertOpen(&openJSON), fileHandler.File)
			case "done":
				var doneJSON messages.Done
				messages.Check(json.Unmarshal(message, &doneJSON))
				write(messages.ConvertDone(&doneJSON), fileHandler.File)
			case "match":
				var matchJSON messages.Match
				messages.Check(json.Unmarshal(message, &matchJSON))
				write(messages.ConvertMatch(&matchJSON), fileHandler.File)
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
