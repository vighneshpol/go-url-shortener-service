package main

import (
	"log"
	"net/http"
)

func main() {

	// Start the HTTP server on port 8080
	log.Println("Starting URL shortener service on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
