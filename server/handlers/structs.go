package handlers

import (
	"time"

	"go.roman.zone/craigslist-monitor/data"
)

type Price struct {
	Price      uint      `json:"price"`
	RecordedAt time.Time `json:"recorded_at"`
}

type Listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	Prices      []Price   `json:"prices"`
	Images      []string  `json:"images"`
	Note        string    `json:"note"`
	Hidden      bool      `json:"isHidden"`
	PostedAt    time.Time `json:"postedAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func convertListingRecords(records []data.ListingRecord) []Listing {
	converted := make([]Listing, len(records))
	for i, record := range records {
		converted[i] = convertListingRecord(record)
	}
	return converted
}

func convertListingRecord(l data.ListingRecord) Listing {
	return Listing{
		ID:          l.ID,
		Title:       l.Title,
		URL:         l.URL,
		Description: l.Description,
		Prices:      convertPrices(l.Prices),
		Images:      convertImages(l.Images),
		Note:        l.Note,
		Hidden:      l.Hidden,
		PostedAt:    l.PostedAt,
		UpdatedAt:   l.UpdatedAt,
	}
}

func convertPrices(prices []data.Price) []Price {
	converted := make([]Price, len(prices))
	for i, price := range prices {
		converted[i] = convertPrice(price)
	}
	return converted
}

func convertPrice(price data.Price) Price {
	return Price{
		Price:      price.Price,
		RecordedAt: price.RecordedAt,
	}
}

func convertImages(images []data.Image) []string {
	converted := make([]string, len(images))
	for i, img := range images {
		converted[i] = img.URL
	}
	return converted
}
