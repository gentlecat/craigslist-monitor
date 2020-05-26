package main

import (
	"flag"
	"log"
	"time"

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

	db, err := data.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dataClient := data.DataClient{
		Database: db,
	}

	dataClient.Init()

	// Running the first refresh immediately
	refreshListing(dataClient)

	for range time.NewTicker(5 * time.Minute).C {
		refreshListing(dataClient)
	}
}

func refreshListing(dataClient data.DataClient) {

	log.Println("Refreshing listings...")

	result, err := craig.SearchByURL(searchURL)
	if err != nil {
		log.Println(err)
	}

	for _, l := range result.Listings {
		dataClient.RecordListing(l)
	}

	log.Printf("Done refreshing. Found %d listings.", len(result.Listings))
}
