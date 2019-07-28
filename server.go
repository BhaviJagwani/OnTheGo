package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"strings"
)

func main() {
	HOST := "localhost"
	PORT := "8554"

	// Start Listening on ip:port
	fmt.Printf("Listening for connections at port %s\n", PORT)
	listener, err := net.Listen("tcp", HOST + ":" + PORT)

	if err != nil {
		fmt.Println("Error starting the server", err.Error())
		os.Exit(1)
	}

	defer listener.Close()

	// Create a channel to pass messages to all connections
	commonMsgChannel := make(chan *Message)

	// Wait for connections
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection from client", err.Error())
		}

		go acceptMessages(conn, commonMsgChannel)

	}
}

func acceptMessages(conn net.Conn, commonMsgChannel chan *Message) {
	// FIXME use ReaderWriter
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	userName := getUserName(writer, reader)

	// channel for this user's messages
	wMsgChannel := make(chan *Message)

	go waitForMessage(wMsgChannel, userName, reader)

	for {
		select {
		case msg := <- commonMsgChannel:
			writer.Write([]byte(fmt.Sprintf("[%s]: %s \n", msg.userName, msg.text)))
			writer.Flush()
		case wMsg := <- wMsgChannel:
			commonMsgChannel <- wMsg
		}
	}
}

func getUserName(writer *bufio.Writer, reader *bufio.Reader) string {
	writer.WriteString("Welcome to OnTheGo Chat Room ! Please enter your user name: \n")
	writer.Flush()

	userName, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Unable to read user name", err.Error())
		// FIXME throw error here
		return ""
	}
	userName = strings.TrimRight(userName, "\r\n")

	writer.WriteString("Hello " + userName + "! Type a message and press Enter to send it to the chat room.\n")
	writer.Flush()

	return userName
}

func waitForMessage(wMsgChannel chan *Message, userName string, reader *bufio.Reader) {
	for {
		// Restricted by the size of the underlying buffer
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected. ", err.Error())
			break
		}

		fmt.Printf("[%v]: %v", userName, message)

		msg := Message{text: message, userName: userName}
		wMsgChannel <- &msg
	}
}

type Message struct{
	text string
	userName string
}