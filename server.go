package main

import (
	"fmt"
	"net"
)

type Server struct {
	Host, Port string
	ConnectedUsers UserSet
	CommonMsgChannel *chan *Message
}


func (server *Server) Start() {
	// Start Listening on ip:port
	fmt.Printf("Listening for connections at port %s\n", server.Port)
	server.ConnectedUsers = NewUserSet()
	// Create a channel to pass messages to all connections
	channel := make(chan *Message)
	server.CommonMsgChannel = &channel
	// Wait for connections
	go server.WaitForConnections()
	// Listen to messages on Common Channel and Broadcast to other users
	go server.BroadcastMessages()

	for{}
}

func (server *Server) acceptClientConnection(conn net.Conn) {

	user := CreateUser(conn, server.CommonMsgChannel)

	// add user to the server group
	if !server.ConnectedUsers.Add(user) {
		fmt.Println("User with user name already exists. Please try connecting again with another user name")
		// throw error
	}

	user.Connect()
}

func (server *Server) WaitForConnections() {
	listener, err := net.Listen("tcp", server.Host + ":" + server.Port)

	if err != nil {
		fmt.Println("Error starting the server", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection from client", err.Error())
		}

		server.acceptClientConnection(conn)

	}
}

func (server *Server) BroadcastMessages() {
	fmt.Println("waiting message1")
	for {
		// Waiting for a message on the channel
		msg := <- *server.CommonMsgChannel
		fmt.Println("Received message1")
		for user, _ := range server.ConnectedUsers.set {
			fmt.Println("Received message")
			if user.UserName != msg.sender {
				user.ReadChan <- msg
			}
		}
	}
}