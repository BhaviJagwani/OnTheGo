package main

import(
	"time"
)

type Message struct{
	text string
	sender *User
	timestamp time.Time
}