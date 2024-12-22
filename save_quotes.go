package main

import (
	"encoding/json"
	"os"
)

// SaveQuotesToFile saves a slice of quotes to a file in JSON format.
func SaveQuotesToFile(filename string, quotes []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(quotes)
}

