package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
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
	commonMsgChannel := make(chan string)

	// Wait for connections
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection from client", err.Error())
		}

		go acceptMessages(conn, commonMsgChannel)

	}
}

func acceptMessages(conn net.Conn, commonMsgChannel chan string) {
	// FIXME use ReaderWriter
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	userName := getUserName(writer, reader)

	// channel for this user's messages
	wMsgChannel := make(chan string)

	go waitForMessage(wMsgChannel, userName, reader)

	for {
		select {
		case msg := <- commonMsgChannel:
			fmt.Println("receive message on common channel")
			writer.WriteString(msg + "\n")
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
	writer.WriteString("Hello " + userName + "! Type a message and press Enter to send it to the chat room.\n")
	writer.Flush()

	return userName
}

func waitForMessage(wMsgChannel chan string, userName string, reader *bufio.Reader) {
	for {
		// Restricted by the size of the underlying buffer
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected. ", err.Error())
			break
		}

		fmt.Printf("[%v]: %v", userName, message)

		wMsgChannel <- message
	}
}