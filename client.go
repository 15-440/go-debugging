package main

type Client struct {
	requests  chan Message
	responses chan string
}

func NewClient(requests chan Message, responses chan string) *Client {
	return &Client{requests: requests, responses: responses}
}

func (c *Client) Send(message string) {
	c.requests <- Message{Type: Data, Payload: message}
}

func (c *Client) Receive() string {
	return <-c.responses
}

func (c *Client) Close() {
	c.requests <- Message{Type: Close}
}
