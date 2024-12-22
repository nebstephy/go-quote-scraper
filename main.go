package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"math/rand"
)

var quotes []string

func main() {
	// Dynamically scrape quotes at runtime
	var err error
	quotes, err = ScrapeQuotes("https://quotes.toscrape.com")
	if err != nil {
		log.Fatalf("Failed to scrape quotes: %v", err)
	}

	http.HandleFunc("/quotes", QuotesHandler)
	http.HandleFunc("/random-quote", RandomQuoteHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if no PORT environment variable is found
	}

	fmt.Printf("Server is running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// QuotesHandler returns all quotes in JSON format.
func QuotesHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"quotes":       quotes,
		"total_quotes": len(quotes),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RandomQuoteHandler returns a random quote in JSON format.
func RandomQuoteHandler(w http.ResponseWriter, r *http.Request) {
	if len(quotes) == 0 {
		http.Error(w, "No quotes available", http.StatusNotFound)
		return
	}

	randomIndex := rand.Intn(len(quotes))
	response := map[string]string{
		"quote": quotes[randomIndex],
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

