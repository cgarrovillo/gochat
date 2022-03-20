package websocket

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	/* Check the origin of connection */
	CheckOrigin: func(r *http.Request) bool {return true},
}

/* Defines a reader which will listen for new messages being sent
to our WS endpoint */
func Reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))

		if err:= conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}


func Writer(conn *websocket.Conn) {
	for {
		fmt.Println(("Sending"))
		messageType, r, err := conn.NextReader()
		if err != nil {
			fmt.Println(err)
			return
		}

		w, err := conn.NextWriter(messageType)
		if _, err := io.Copy(w, r); err != nil {
			fmt.Println(err)
			return
		}

		if err := w.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}
}