package main

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// ScrapeQuotes scrapes quotes from the given URL and returns them as a slice of strings.
func ScrapeQuotes(url string) ([]string, error) {
	log.Printf("Fetching URL: %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch the page")
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var quotes []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "span" {
			for _, attr := range n.Attr {
				if attr.Key == "class" && attr.Val == "text" {
					if n.FirstChild != nil {
						quote := strings.TrimSpace(n.FirstChild.Data)
						quotes = append(quotes, quote)
					}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if len(quotes) == 0 {
		return nil, errors.New("no quotes found on the page")
	}

	log.Printf("Scraped %d quotes\n", len(quotes))
	return quotes, nil
}

