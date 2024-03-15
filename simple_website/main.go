package main

import (
	"log"
	"net/http"
)

// Maps routes and starts the http server
func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))

	log.Println("Starting server at localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
