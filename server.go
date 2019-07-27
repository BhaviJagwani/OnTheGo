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

	fmt.Printf("Listening for connections at port %s\n", PORT)
	listener, err := net.Listen("tcp", HOST + ":" + PORT)

	if err != nil {
		fmt.Println("Error starting the server", err.Error())
		os.Exit(1)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection from client", err.Error())
		}

		go acceptMessage(conn)

	}
}

func acceptMessage(conn net.Conn) {
	// Restricted by the size of the underlying buffer
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println("Client disconnected. ", err.Error())
			break
		}

		fmt.Printf("%s", message)
	}
}