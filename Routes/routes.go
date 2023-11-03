package routes

import (
	"fmt"
	"log"
	"net/http"
	websockets "studentsPlayground/Websockets"
)

var port = ":4300"

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	//checks if the incoming request can connect
	websockets.Upgraderr.CheckOrigin = func(r *http.Request) bool { return true }
}

func Routes() {
	http.HandleFunc("/", homePage)
	log.Printf("starting server at port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
