package main

type MessageType int

const (
	Data MessageType = iota
	Close
)

type Message struct {
	Type    MessageType
	Payload string
}
