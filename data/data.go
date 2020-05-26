package data

import (
	"log"
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go.roman.zone/craig"
)

func (dbClient *DataClient) GetAllListingRecords() []ListingRecord {
	var recs []ListingRecord
	dbClient.Database.Preload("Prices").Preload("Images").Find(&recs)
	return recs
}

func (dbClient *DataClient) FindListingRecord(listingID string) ListingRecord {
	var rec ListingRecord
	dbClient.Database.Preload("Prices").Where("id = ?", listingID).First(&rec)
	return rec
}

func (dbClient *DataClient) RecordListing(listing craig.Listing) {

	// TODO: Add duplicate checks by name

	existingRecord := dbClient.FindListingRecord(listing.ID)
	if existingRecord.ID != "" {

		lastRecordedPrice := existingRecord.Prices[len(existingRecord.Prices)-1].Price

		converted := convertListing(listing)
		converted.Hidden = existingRecord.Hidden
		converted.Note = existingRecord.Note

		if lastRecordedPrice != listing.Price {
			log.Printf("Price of listing %s has changed. Updating [was $%d, now $%d]",
				listing.ID, lastRecordedPrice, listing.Price)
			converted.Prices = append(existingRecord.Prices, Price{Price: listing.Price, RecordedAt: time.Now()})
		}

		dbClient.Database.Save(&converted)

		return
	}

	newListing := convertListing(listing)
	log.Printf("Recording a new listing: %s", newListing.URL)
	dbClient.Database.Create(&newListing)
}

func convertListing(listing craig.Listing) ListingRecord {
	return ListingRecord{
		ID:          listing.ID,
		URL:         listing.URL,
		Title:       listing.Title,
		Description: listing.Description,
		Prices: []Price{{
			Price:      listing.Price,
			RecordedAt: time.Now(),
		}},
		Images:    convertImages(listing.Images),
		Hidden:    false,
		PostedAt:  *listing.PostedAt,
		UpdatedAt: *listing.UpdatedAt,
	}
}

func convertImages(images []craig.Image) []Image {
	converted := make([]Image, len(images))
	for i, img := range images {
		converted[i] = Image{
			URL: img.Large,
		}
	}
	return converted
}

func (dbClient *DataClient) SetNote(listingID string, note string) {
	dbClient.Database.Model(&ListingRecord{ID: listingID}).Updates(ListingRecord{Note: note})
}

func (dbClient *DataClient) Hide(listingID string) {
	log.Printf("Hiding listing %s", listingID)
	dbClient.Database.Model(&ListingRecord{ID: listingID}).Updates(ListingRecord{Hidden: true})
}

func (dbClient *DataClient) Unhide(listingID string) {
	log.Printf("Un-hiding listing %s", listingID)
	dbClient.Database.Model(&ListingRecord{ID: listingID}).Updates(ListingRecord{Hidden: false})
}
