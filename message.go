package main

import(
	"time"
)

type Message struct{
	text string
	sender string
	timestamp time.Time
}