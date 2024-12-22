package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Use the PORT environment variable, default to 8080 for local testing
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("Defaulting to port 8080")
	}

	// Set up routes
	http.HandleFunc("/quotes", QuotesHandler)
	http.HandleFunc("/random-quote", RandomQuoteHandler)

	// Start the server
	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

