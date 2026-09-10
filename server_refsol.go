// go test server_refsol.go client.go server_test.go -args -log
package main

import (
	"flag"
	"log"
)

var loggingEnabled = flag.Bool("log", false, "enable server debug logging")

func debugf(format string, args ...interface{}) {
	if *loggingEnabled {
		log.Printf(format, args...)
	}
}

type Server struct {
	requests  <-chan string
	responses chan<- string
}

func NewServer(requests <-chan string, responses chan<- string) *Server {
	return &Server{requests: requests, responses: responses}
}

func (s *Server) Run() {
	for message := range s.requests {
		debugf("server received message=%q", message)
		if message == CloseMessage {
			debugf("server received close; stopping without sending a response")
			return // FIX: CLOSE ends the server before it sends another response.
		}

		response := "ACK:" + message
		debugf("server sending response=%q", response)
		s.responses <- response
	}
	debugf("server request channel closed")
}
