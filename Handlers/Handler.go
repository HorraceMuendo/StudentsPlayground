package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn []websocket.Conn
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/muendo/Desktop/studentsPlayground/index.html")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	//checks if the incoming request can connect

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// Attempting to upragrade the connection to a websocket

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("The client is connected \n")

	conn = append(conn, *ws)

	for {

		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		// log to the terminal
		log.Println("%s CLIENT : %s\n", ws.RemoteAddr(), string(msg))

		for _, client := range conn {
			if err = client.WriteMessage(msgType, msg); err != nil {
				fmt.Println(err)
			}
		}

	}

}

// _, payload, err := conn.ReadMessage()
// if err != nil {
// 	log.Println(err)
// 	return
// }

// fmt.Println("CLIENT: ", string(payload))

// // Echo the received message back to the client
// if err := conn.WriteMessage(message, payload); err != nil {
// 	log.Println(err)
// 	return
// }

// dummy code for writing to client from server

// for {
// 	var message string
// 	fmt.Scanln(&message)
// 	m := []byte(message)

// 	// writing back to the client
// 	err = ws.WriteMessage(websocket.TextMessage, m)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	go handlers.ReadMessage(ws)

// }
