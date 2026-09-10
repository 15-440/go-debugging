package main

import (
	"testing"
	"time"
)

func TestClientCloseStopsServer(t *testing.T) {
	requests := make(chan Message)
	responses := make(chan string)
	server := NewServer(requests, responses)
	client := NewClient(requests, responses)
	finished := make(chan struct{})

	go func() {
		server.Run()
		// we block here until .Run() finishes
		close(finished)
	}()

	messages := []string{"hello", "world"}
	for _, message := range messages {
		client.Send(message)
		// check that we are getting the results back properly
		if got := client.Receive(); got != "ACK:"+message {
			t.Fatalf("response = %q, want %q", got, "ACK:"+message)
		}
	}

	client.Close()
	select {
	case <-finished:
		// our server closed in time!
	case <-time.After(100 * time.Millisecond):
		t.Fatal("server did not stop after client sent Close")
	}
}
