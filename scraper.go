package main

import (
    "fmt"
    "log"
    "github.com/PuerkitoBio/goquery"
    "net/http"
    "encoding/json"
    "os"
)

func main() {
    // Scrape the quotes from a website
    res, err := http.Get("https://quotes.toscrape.com/") // Example URL
    if err != nil {
        log.Fatal("Error fetching page: ", err)
    }
    defer res.Body.Close()

    if res.StatusCode != 200 {
        log.Fatal("Error: Status Code ", res.StatusCode)
    }

    // Parse the HTML page using goquery
    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        log.Fatal("Error parsing HTML: ", err)
    }

    // Slice to store quotes
    var quotes []string

    // Find and scrape quotes from the page
    doc.Find(".quote .text").Each(func(index int, item *goquery.Selection) {
        quote := item.Text()
        quotes = append(quotes, quote)
    })

    // Save the quotes to quotes.json
    jsonData, err := json.MarshalIndent(quotes, "", "    ")
    if err != nil {
        log.Fatalf("Error marshaling data: %v", err)
    }

    err = os.WriteFile("quotes.json", jsonData, 0644)
    if err != nil {
        log.Fatalf("Error writing file: %v", err)
    }

    fmt.Println("Quotes have been successfully scraped and saved to quotes.json")
}

