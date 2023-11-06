package routes

import (
	"log"
	"net/http"
	handlers "studentsPlayground/Handlers"
)

var port = ":4300"

func Routes() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/ws", handlers.WSEndpoint)
	log.Printf("starting server at port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
