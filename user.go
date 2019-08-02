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
	QuitChannel *chan *Message

	Reader *bufio.Reader
	Writer *bufio.Writer
}

func CreateUser(conn net.Conn, commonChannel *chan *Message, quitChannel *chan *Message) *User{
	user := User{}
	
	// FIXME use ReaderWriter
	user.Writer = bufio.NewWriter(conn)
	user.Reader = bufio.NewReader(conn)

	user.PopulateUserName()

	// channel for this user to read messages from
	user.ReadChan = make(chan *Message)
	// channel for this user to write messages to
	user.CommonMsgChannel = commonChannel
	user.QuitChannel = quitChannel
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
			*user.QuitChannel <- & Message{sender: user, timestamp: time.Now()}
			break
		}

		fmt.Printf("[%v]: %v", user.UserName, message)

		msg := Message{text: message, sender: user, timestamp: time.Now()}

		// Write message to all channels
		go func(channel *chan *Message, msg *Message) {
			*channel <- msg
			} (user.CommonMsgChannel, &msg)
	}
}

func (user *User) ListenForMessages() {
	for {
		msg := <- user.ReadChan
		user.Writer.WriteString(fmt.Sprintf("[%s %s]: %s \n", formatTime(msg.timestamp), msg.sender.UserName, msg.text))
		user.Writer.Flush()

	}

}

func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05-07:00")
}

func (user *User) Connect() {
	go user.EnableSendMessage()
	go user.ListenForMessages()
}