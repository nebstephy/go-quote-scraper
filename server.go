package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var quotes []string

func init() {
	// Initialize the random seed
	rand.Seed(time.Now().UnixNano())

	// Read quotes from file
	data, err := os.ReadFile("quotes.json") // Updated to use os.ReadFile
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	err = json.Unmarshal(data, &quotes)
	if err != nil {
		log.Fatalf("Error unmarshaling data: %v", err)
	}
}

func quotesHandler(w http.ResponseWriter, r *http.Request) {
	// Return the full list of quotes as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"total_quotes": len(quotes),
		"quotes":       quotes,
	})
}

func randomQuoteHandler(w http.ResponseWriter, r *http.Request) {
	// Pick a random quote
	randomIndex := rand.Intn(len(quotes))

	// Create a response with the random quote
	response := map[string]interface{}{
		"quote": quotes[randomIndex],
	}

	// Set the content type and encode the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Define routes
	http.HandleFunc("/quotes", quotesHandler)
	http.HandleFunc("/random-quote", randomQuoteHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

