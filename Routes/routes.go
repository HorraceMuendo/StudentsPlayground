package routes

import (
	"fmt"
	"log"
	"net/http"
)

var port = ":4300"

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func Routes() {
	http.HandleFunc("/", homePage)
	log.Printf("starting server at port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
