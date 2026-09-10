// Buggy TA implementation.
package main

type Server struct {
	requests  chan Message
	responses chan string
}

func NewServer(requests chan Message, responses chan string) *Server {
	return &Server{requests: requests, responses: responses}
}

func (s *Server) Run() {
	for message := range s.requests {
		s.responses <- "ACK:" + message.Payload
	}
}
