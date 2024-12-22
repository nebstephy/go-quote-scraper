package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

var quotes []string

// QuotesHandler handles the /quotes endpoint
func QuotesHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"quotes":       quotes,
		"total_quotes": len(quotes),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RandomQuoteHandler handles the /random-quote endpoint
func RandomQuoteHandler(w http.ResponseWriter, r *http.Request) {
	if len(quotes) == 0 {
		http.Error(w, "No quotes available", http.StatusInternalServerError)
		return
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))
	response := map[string]string{
		"quote": quotes[randomIndex],
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

