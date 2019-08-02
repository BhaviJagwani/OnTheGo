package main

import (
	"fmt"
	"net"
)

type Server struct {
	Host, Port string
	ConnectedUsers UserGroup
	CommonMsgChannel chan *Message
	QuitChan chan *Message
}


func (server *Server) Start() {
	server.ConnectedUsers = NewUserGroup()
	// Create a channel to pass messages to all connections
	server.CommonMsgChannel = make(chan *Message)
	server.QuitChan = make(chan *Message)

	// Start Listening on ip:port
	fmt.Printf("Listening for connections at port %s\n", server.Port)
	listener, err := net.Listen("tcp", ":" + server.Port)

	if err != nil {
		fmt.Println("Error starting the server", err.Error())
		return
	}

	// Listen to messages on Common Channel and Broadcast to other users
	go server.HandleMessages()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection from client", err.Error())
		}

		defer conn.Close()
		server.acceptClientConnection(conn)
	}

}

func (server *Server) HandleMessages() {
	for {
		select {
		// Waiting for a message on the channel
		case msg := <- server.CommonMsgChannel:
			for user, _ := range server.ConnectedUsers.set {
				if user.UserName != msg.sender.UserName {
					user.ReadChan <- msg
				}
			}
		case quitMsg := <- server.QuitChan:
			server.ConnectedUsers.Remove(quitMsg.sender)
		}
	}
}

func (server *Server) acceptClientConnection(conn net.Conn) {

	user := CreateUser(conn, &server.CommonMsgChannel, &server.QuitChan)

	// add user to the server group
	if !server.ConnectedUsers.Add(user) {
		fmt.Println("User with user name already exists. Please try connecting again with another user name")
		// throw error
	}

	user.Connect()
}
