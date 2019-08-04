package main

import(
	"fmt"
	"io"
	"os"
	"syscall"
	"path/filepath"
)

type MessageLogger struct {
	file string
	MsgChannel chan *Message
}

func (logger *MessageLogger) Start() {
	logger.MsgChannel = make(chan *Message)

	// FIXME get from config

	folderName := "<log_folder_path>"
	fileName := "messages.log"
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
        panic(err)
    }

	path := filepath.Join(folderName, fileName)

	file, err := os.OpenFile(path, syscall.O_WRONLY|syscall.O_CREAT|syscall.O_APPEND, 0666)
	fmt.Println("Logging all messages to " + path)
    if err != nil {
		panic(err)
    }
	defer file.Close()

	for {
		msg := <- logger.MsgChannel
		// Write message to file
		_, err = io.WriteString(file, msg.GetPrintableMessage())

		if err != nil {
			fmt.Println(err)
		}
		file.Sync()
	}
}
