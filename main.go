package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var (
	quotes []string
	mu     sync.Mutex
)

// Scrape quotes from the website
func scrapeQuotes() error {
	const baseURL = "https://quotes.toscrape.com/page/%d/"
	mu.Lock()
	defer mu.Unlock()

	quotes = nil // Clear existing quotes
	for page := 1; page <= 10; page++ {
		resp, err := http.Get(fmt.Sprintf(baseURL, page))
		if err != nil {
			return fmt.Errorf("failed to fetch page %d: %v", page, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("error fetching page %d: %s", page, resp.Status)
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to parse page %d: %v", page, err)
		}

		doc.Find(".quote .text").Each(func(i int, s *goquery.Selection) {
			quotes = append(quotes, s.Text())
			// Stop collecting once we reach 100 quotes
			if len(quotes) >= 100 {
				return
			}
		})

		if len(quotes) >= 100 {
			break
		}
	}

	return nil
}

func getQuotes(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(quotes)
}

func main() {
	// Initial scrape
	if err := scrapeQuotes(); err != nil {
		fmt.Println("Failed to scrape quotes:", err)
		return
	}

	http.HandleFunc("/quotes", getQuotes)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

