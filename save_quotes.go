package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    quotes := []string{
        "The world as we have created it is a process of our thinking. It cannot be changed without changing our thinking.",
        "It is our choices, Harry, that show what we truly are, far more than our abilities.",
        "There are only two ways to live your life. One is as though nothing is a miracle. The other is as though everything is a miracle.",
        "The person, be it gentleman or lady, who has not pleasure in a good novel, must be intolerably stupid.",
        "Imperfection is beauty, madness is genius and it's better to be absolutely ridiculous than absolutely boring.",
        // Add more quotes here...
    }

    // Save the quotes as a JSON file
    jsonData, err := json.MarshalIndent(quotes, "", "    ")
    if err != nil {
        log.Fatalf("Error marshaling data: %v", err)
    }

    err = ioutil.WriteFile("quotes.json", jsonData, 0644)
    if err != nil {
        log.Fatalf("Error writing file: %v", err)
    }

    fmt.Println("Quotes have been saved to quotes.json")
}

