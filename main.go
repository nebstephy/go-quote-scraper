package main

import (
	"log"
	"net/http"
)

var quotes []string // Declare the global quotes variable

func main() {
	// Step 1: Scrape Quotes from the website
	const url = "https://quotes.toscrape.com"
	log.Println("Starting to scrape quotes...")
	var err error
	quotes, err = ScrapeQuotes(url)
	if err != nil {
		log.Fatalf("Failed to scrape quotes: %v", err)
	}
	log.Printf("Successfully scraped %d quotes\n", len(quotes))

	// Step 2: Save Quotes to a File
	log.Println("Saving quotes to file...")
	err = SaveQuotesToFile("quotes.json", quotes)
	if err != nil {
		log.Fatalf("Failed to save quotes to file: %v", err)
	}
	log.Println("Quotes saved to quotes.json")

	// Step 3: Start the HTTP Server
	http.HandleFunc("/quotes", QuotesHandler)
	http.HandleFunc("/random-quote", RandomQuoteHandler) // Add random quote endpoint
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

