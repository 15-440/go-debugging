// Buggy TA implementation.
package main

// type Message struct {
//     message.Type
// 	   message.Payload
// }

import (
	"flag"
	"log"
)

var loggingEnabled = flag.Bool("log", false, "enable server logging")

type Server struct {
	requests  chan Message
	responses chan string
}

func NewServer(requests chan Message, responses chan string) *Server {
	return &Server{requests: requests, responses: responses}
}

func (s *Server) Run() {
	for message := range s.requests {
		if *loggingEnabled {
			log.Printf("server received type=%d payload=%q", message.Type, message.Payload)
		}
		response := "ACK:" + message.Payload
		if *loggingEnabled {
			log.Printf("server sending response=%q", response)
		}
		s.responses <- response

	}
}
