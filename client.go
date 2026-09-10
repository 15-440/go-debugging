package main

const CloseMessage = "CLOSE"

type Client struct {
	requests  chan<- string
	responses <-chan string
}

func NewClient(requests chan<- string, responses <-chan string) *Client {
	return &Client{requests: requests, responses: responses}
}

func (c *Client) Send(message string) {
	c.requests <- message
}

func (c *Client) Receive() string {
	return <-c.responses
}

func (c *Client) Close() {
	c.Send(CloseMessage)
}
