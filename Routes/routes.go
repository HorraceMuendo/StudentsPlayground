package routes

import (
	"fmt"
	"log"
	"net/http"
	handlers "studentsPlayground/Handlers"

	"github.com/gorilla/websocket"
)

var port = ":4300"

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

// check on the package importation

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	//checks if the incoming request can connect

	handlers.Upgraderr.CheckOrigin = func(r *http.Request) bool { return true }

	// Attempting to upragrade the connection to a websocket

	ws, err := handlers.Upgraderr.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("The client is connected \n")

	var message string
	fmt.Scanln(&message)
	m := []byte(message)

	// writing back to the client
	err = ws.WriteMessage(websocket.TextMessage, m)
	if err != nil {
		log.Println(err)
	}

	go handlers.ReadMessage(ws)
}

func Routes() {
	http.HandleFunc("/ws", wsEndpoint)
	log.Printf("starting server at port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
