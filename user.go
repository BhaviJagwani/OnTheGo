package main

import (
	"bufio"
	"net"
	"fmt"
	"strings"
	"time"
)

type User struct {
	UserName string
	ReadChan chan *Message
	CommonMsgChannel *chan *Message

	Reader *bufio.Reader
	Writer *bufio.Writer
}

func CreateUser(conn net.Conn, commonChannel *chan *Message) *User{
	user := User{}
	
	// FIXME use ReaderWriter
	user.Writer = bufio.NewWriter(conn)
	user.Reader = bufio.NewReader(conn)

	user.PopulateUserName()

	// channel for this user to read messages from
	user.ReadChan = make(chan *Message)
	// channel for this user to write messages to
	user.CommonMsgChannel = commonChannel
	return &user
}

func (user *User) PopulateUserName() {
	user.Writer.WriteString("Welcome to OnTheGo Chat Room ! Please enter your user name: \n")
	user.Writer.Flush()

	userName, err := user.Reader.ReadString('\n')

	if err != nil {
		fmt.Println("Unable to read user name", err.Error())
		// FIXME throw error here
		return
	}
	userName = strings.TrimRight(userName, "\r\n")

	user.Writer.WriteString("Hello " + userName + "! Type a message and press Enter to send it to the chat room.\n")
	user.Writer.Flush()

	user.UserName = userName
}

func (user *User) EnableSendMessage() {
	for {
		// Restricted by the size of the underlying buffer
		message, err := user.Reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected. ", err.Error())
			break
		}

		fmt.Printf("[%v]: %v", user.UserName, message)

		msg := Message{text: message, sender: user.UserName, timestamp: time.Now()}

		// Write message to all channels
		go func(channel *chan *Message, msg *Message) {
			*channel <- msg
			} (user.CommonMsgChannel, &msg)
	}
}

func (user *User) ListenForMessages() {
	for {
		msg := <- user.ReadChan
		fmt.Println("User message")
		user.Writer.WriteString(fmt.Sprintf("[%s]: %s \n", msg.sender, msg.text))
		user.Writer.Flush()

	}

}

func (user *User) Connect() {
	go user.EnableSendMessage()
	go user.ListenForMessages()
}