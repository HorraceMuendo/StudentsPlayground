package handlers

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

var Upgraderr = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ReadMessage(conn *websocket.Conn) {

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
