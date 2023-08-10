package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sunesimonsen/htmx-hackernews/server"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, server); err != nil {
		log.Fatal(err)
	}
}
