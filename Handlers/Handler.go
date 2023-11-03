package handlers

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func Reader(conn *websocket.Conn) {

	message, payload, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(payload))

	if err := conn.WriteMessage(message, payload); err != nil {
		log.Println(err)
		return
	}
}
