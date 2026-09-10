// go test server.go client.go server_test.go
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
		s.responses <- "ACK:" + message
	}
}
