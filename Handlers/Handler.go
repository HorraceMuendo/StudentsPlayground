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

	_, payload, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("CLIENT: ", string(payload))

	// // Echo the received message back to the client
	// if err := conn.WriteMessage(message, payload); err != nil {
	// 	log.Println(err)
	// 	return
	// }
}
