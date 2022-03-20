package main

import (
	"fmt"
	"net/http"

	"gochat/websocket"
)

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple server")
	})

	// map /ws endpoint to ServeWS function
	http.HandleFunc("/ws", websocket.ServeWS)
} 

func main() {
	fmt.Println("Chat App v0.0.1")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}