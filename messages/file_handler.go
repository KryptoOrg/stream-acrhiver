package messages

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

type FileHandler struct {
	Timestamp int64
	Frequency int64
	Filename  string
	File      *os.File
	Directory string
}

func GenerateNewFileName(handler *FileHandler) string {
	date := time.Unix(0, handler.Timestamp).Format("20060102")
	datedDirectory := filepath.Join(handler.Directory, date)
	if _, err := os.Stat(datedDirectory); os.IsNotExist(err) {
		err := os.Mkdir(datedDirectory, os.ModePerm)
		Check(err)
	}
	timestampedFilename := fmt.Sprintf("%v_%s", handler.Timestamp, handler.Filename)
	fullPath := filepath.Join(datedDirectory, timestampedFilename)
	return fullPath
}

func Update(handler *FileHandler) error {
	currentTime := (time.Now().UnixNano() / handler.Frequency) * handler.Frequency
	if currentTime >= handler.Timestamp || handler.File == nil {
		handler.Timestamp = currentTime + handler.Frequency
		if handler.File != nil {
			Close(handler)
		}
		filename := GenerateNewFileName(handler)
		log.Infof("Updated Filename to %s\n", filename)
		file, err := os.Create(filename)
		handler.File = file
		return err
	}
	return nil
}

func Close(handler *FileHandler) {
	log.Infof("Closing %s\n", handler.File.Name())
	err := handler.File.Close()
	if err != nil {
		log.Fatalf("Error while closing File %s\n", err)
	}
}
