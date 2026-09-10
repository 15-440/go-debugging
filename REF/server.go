//go:build ignore

// Copy this file over the root server.go during the lesson.
// It is intentionally excluded from normal Go builds so the root TA exercise
// remains the only package that VS Code and go test analyze.
package main

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

		switch message.Type {
		case Close:
			if *loggingEnabled {
				log.Printf("server received Close; stopping without sending a response")
			}
			return
		case Data:
			response := "ACK:" + message.Payload
			if *loggingEnabled {
				log.Printf("server sending response=%q", response)
			}
			s.responses <- response
		}
	}
}
