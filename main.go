package main

import (
	"flag"
	"fmt"
	"log"

	"go.roman.zone/craig"
)

func main() {

	var searchURL string
	flag.StringVar(&searchURL, "u", "url", "Craigslist search URL")
	flag.Parse()

	if len(searchURL) == 0 {
		log.Fatal("URL argument is empty")
	}

	result, err := craig.SearchByURL(searchURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d listings:\n", len(result.Listings))
	for _, l := range result.Listings {
		fmt.Printf("$%d | %s | %s\n", l.Price, l.Title, l.URL)
	}
}
