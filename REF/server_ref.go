// Buggy TA implementation.
package main

// type Message struct {
//     message.Type
// 	   message.Payload
// }

import (
	"flag"
	"log"
	"os"
)

var loggingEnabled = flag.Bool("log", false, "enable server logging")

type Server struct {
	requests  chan Message
	responses chan string
	logger    *log.Logger
}

func NewServer(requests chan Message, responses chan string) *Server {
	f, err := os.Create("server.log")
	if err != nil {
		panic(err)
	}

	logger := log.New(f, "", log.LstdFlags)

	return &Server{requests: requests, responses: responses, logger: logger}
}

func (s *Server) Run() {
	for message := range s.requests {
		if *loggingEnabled {
			// log.Printf("server received type=%d payload=%q", message.Type, message.Payload)
			s.logger.Printf("server received type=%d payload=%q", message.Type, message.Payload)
		}

		switch message.Type {
		case Close:
			return
		case Data:
			response := "ACK:" + message.Payload
			if *loggingEnabled {
				// log.Printf("server sending response=%q", response)
				s.logger.Printf("server sending response=%q", response)
			}
			s.responses <- response
		}

	}
}
