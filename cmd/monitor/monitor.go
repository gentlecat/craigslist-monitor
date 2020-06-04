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

	go func() {
		// Running the first check immediately
		checkListings(dataClient)

		for range time.NewTicker(5 * time.Minute).C {
			checkListings(dataClient)
		}
	}()

	// Running the first refresh immediately
	getListings(dataClient)

	for range time.NewTicker(5 * time.Minute).C {
		getListings(dataClient)
	}
}

func getListings(dataClient data.DataClient) {

	log.Println("[refresher] Refreshing listings...")

	result, err := craig.SearchByURL(searchURL)
	if err != nil {
		log.Println(err)
		return
	}

	for _, l := range result.Listings {
		dataClient.RecordListing(l)
	}

	log.Printf("[refresher] Done refreshing. Found %d listings.", len(result.Listings))
}

func checkListings(dataClient data.DataClient) {
	log.Println("[checker] Checking known listings...")

	records := dataClient.GetAllListingRecords()

	removedCount := 0

	for _, r := range records {
		_, err := craig.GetListing(r.URL)
		if err != nil {
			log.Printf("[checker] Unable to retrieve listing %s. Assuming it's deleted. Raw err: %s", r.ID, err)
			dataClient.MarkDeleted(r.ID)
			removedCount++
		}

		// Let's not bomb them with too many requests
		time.Sleep(1 * time.Second)
	}

	log.Printf("[checker] Done checking. Marked %d listings as deleted.", removedCount)
}
