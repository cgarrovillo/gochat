package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Client connecting to this server
type Client struct {
	ID string
	Conn *websocket.Conn
	Pool *Pool
}

// Message being transported
type Message struct {
	Type int `json:"type"`
	Body string `json:"body"`
}

// A method that operates on a pointer to Client
func (c *Client) Read() {
	defer func() {
		// <- Channel direction operator
		// Send c to Unregister
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}

		// Send message to Broadcast
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}