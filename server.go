package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

// RandomQuoteHandler serves a random quote as a JSON response.
func RandomQuoteHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))
	response := map[string]string{
		"quote": quotes[randomIndex],
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// QuotesHandler serves all quotes as a JSON response.
func QuotesHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"total_quotes": len(quotes),
		"quotes":       quotes,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

