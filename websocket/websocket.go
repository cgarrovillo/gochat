package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	/* Check the origin of connection */
	CheckOrigin: func(r *http.Request) bool {return true},
}

/* Defines a reader which will listen for new messages being sent
to our WS endpoint */
func reader(conn *websocket.Conn) {
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

/* Our WebSocket endpoint */
func ServeWS(w http.ResponseWriter, r *http.Request)	 {
	fmt.Println(r.Host)

	// upgrade connection to WS connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if (err != nil ) {
		log.Println((err))
	}

	// listen indefinitely
	reader(ws)
}

