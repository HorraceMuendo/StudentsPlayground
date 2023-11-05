package routes

import (
	"fmt"
	"log"
	"net/http"
	handlers "studentsPlayground/Handlers"
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

	// writing back to the client
	err = ws.WriteMessage(1, []byte("Welcome to students playground client"))
	if err != nil {
		log.Println(err)
	}

	handlers.ReadMessage(ws)
}

func Routes() {
	http.HandleFunc("/", wsEndpoint)
	log.Printf("starting server at port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
