package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

func WSEndpoint(w http.ResponseWriter, r *http.Request) {
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

// work on the upload first then the sending

func FileTransfer() {

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error creating file", err)
	}
	writer := io.Writer(file)
	n, err := writer.Write([]byte("file trans"))
	fmt.Println(n, err)
	defer file.Close()
}
