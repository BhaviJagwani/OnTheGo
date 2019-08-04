package main

import(
	"time"
	"fmt"
)

type Message struct{
	text string
	sender *User
	timestamp time.Time
}

func (msg *Message) GetPrintableMessage() string {
	return fmt.Sprintf("[%s %s]: %s \n", formatTime(msg.timestamp), msg.sender.UserName, msg.text)
}