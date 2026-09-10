// Run against the buggy implementation:
//
//	go test server.go client.go server_test.go
//
// Run against the reference implementation:
//   - go test server_refsol.go client.go server_test.go
//   - go test server_refsol.go client.go server_test.go -args -log
package main

import (
	"testing"
	"time"
)

func TestClientCloseStopsServer(t *testing.T) {
	requests := make(chan string)
	responses := make(chan string)
	server := NewServer(requests, responses)
	client := NewClient(requests, responses)
	finished := make(chan struct{})

	go func() {
		server.Run()
		// blocks until .Run() returns / finishes
		close(finished)
	}()

	// start sending messages to the server
	for _, message := range []string{"hello", "world"} {
		client.Send(message)
		// check to see that we got an Ack from the server
		if got := client.Receive(); got != "ACK:"+message {
			t.Fatalf("response = %q, want %q", got, "ACK:"+message)
		}
	}
	// send the close message to the server
	client.Close()
	select {
	case <-finished:
		// The server correctly stops after receiving CLOSE.
	case <-time.After(100 * time.Millisecond):
		t.Fatal("server did not stop after client sent CLOSE;")
	}
}
