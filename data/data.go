package data

import (
	"log"
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"go.roman.zone/craig"
)

type Price struct {
	ID              uint `gorm:"primary_key"`
	ListingRecordID string
	Price           uint

	RecordedAt time.Time
}

type ListingRecord struct {
	ID     string `gorm:"primary_key"`
	Title  string
	Prices []Price
	Note   string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (dbClient *DataClient) GetAllListingRecords() []ListingRecord {
	var recs []ListingRecord
	dbClient.Database.Preload("Prices").Find(&recs)
	return recs
}

func (dbClient *DataClient) FindListingRecord(listingID string) ListingRecord {
	var rec ListingRecord
	dbClient.Database.Preload("Prices").Where("id = ?", listingID).First(&rec)
	return rec
}

func (dbClient *DataClient) RecordListing(listing craig.Listing) {

	existingRecord := dbClient.FindListingRecord(listing.ID)
	if existingRecord.ID != "" {
		log.Printf("Found existing record: %s",
			existingRecord.ID)

		lastRecordedPrice := existingRecord.Prices[len(existingRecord.Prices)-1].Price

		if lastRecordedPrice != listing.Price {
			log.Printf("Price of listing %s has changed. Updating [was $%d, now $%d]",
				listing.ID, lastRecordedPrice, listing.Price)
			existingRecord.Prices = append(existingRecord.Prices, Price{Price: listing.Price, RecordedAt: time.Now()})
			dbClient.Database.Save(&existingRecord)
		}

		return
	}

	newListing := ListingRecord{
		ID:     listing.ID,
		Title:  listing.Title,
		Prices: []Price{Price{Price: listing.Price + 100, RecordedAt: time.Now()}},
	}
	log.Printf("Recording a new listing: %+v", newListing)
	dbClient.Database.Create(&newListing)
}

func (dbClient *DataClient) SetNote(listingID string, note string) {
	dbClient.Database.Model(&ListingRecord{ID: listingID}).Updates(ListingRecord{Note: note})
}
