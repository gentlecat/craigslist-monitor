package main

import (
	"flag"
	"log"

	"go.roman.zone/craig"
	"go.roman.zone/craigslist-monitor/data"
)

var (
	searchURL string
)

func main() {

	flag.StringVar(&searchURL, "u", "url", "Craigslist search URL")
	flag.Parse()

	if len(searchURL) == 0 {
		log.Fatal("URL argument is empty")
	}

	result, err := craig.SearchByURL(searchURL)
	if err != nil {
		log.Fatal(err)
	}

	db, err := data.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dataClient := data.DataClient{
		Database: db,
	}

	dataClient.Init()

	log.Printf("Found %d listings:\n", len(result.Listings))
	for _, l := range result.Listings {
		dataClient.RecordListing(l)
	}
}
