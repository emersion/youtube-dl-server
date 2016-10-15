package main

import (
	"log"
	"net/http"

	"github.com/emersion/youtube-dl-server/server"
)

func main() {
	http.Handle("/api/", http.StripPrefix("/api", server.New()))
	http.Handle("/", http.FileServer(http.Dir("public")))

	log.Println("Starting server on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
