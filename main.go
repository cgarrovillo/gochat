package main

import (
	"fmt"
	"log"
	"net/http"

	"gochat/pkg/websocket"
)

/* Our WebSocket endpoint */
func serveWS(w http.ResponseWriter, r *http.Request)	 {
	fmt.Println(r.Host)

	// upgrade connection to WS connection
	ws, err := websocket.Upgrader.Upgrade(w, r, nil)
	if (err != nil ) {
		log.Println((err))
	}
	
	// Start a goroutine for WS Writer
	go websocket.Writer(ws)
	// listen for messages indefinitely
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple server")
	})

	// map /ws endpoint to ServeWS function
	http.HandleFunc("/ws", serveWS)
} 

func main() {
	fmt.Println("Chat App v0.0.1")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}