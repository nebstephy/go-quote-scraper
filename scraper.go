package main

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// ScrapeQuotes scrapes quotes from a given URL.
func ScrapeQuotes(url string) ([]string, error) {
	log.Println("Fetching URL:", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching the page: %v", err)
		return nil, errors.New("failed to fetch the page")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code: %d", resp.StatusCode)
		return nil, errors.New("failed to fetch the page")
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Printf("Error parsing the page: %v", err)
		return nil, errors.New("failed to parse the page")
	}

	var quotes []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "span" {
			for _, attr := range n.Attr {
				if attr.Key == "class" && attr.Val == "text" {
					if n.FirstChild != nil {
						quote := strings.TrimSpace(n.FirstChild.Data)
						if quote != "" {
							quotes = append(quotes, quote)
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if len(quotes) == 0 {
		log.Println("No quotes found on the page")
		return nil, errors.New("no quotes found on the page")
	}

	log.Printf("Scraped %d quotes", len(quotes))
	return quotes, nil
}

