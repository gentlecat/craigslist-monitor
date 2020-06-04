package data

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type ListingRecord struct {
	ID          string `gorm:"primary_key"`
	URL         string
	Title       string
	Description string
	Prices      []Price
	Images      []Image
	Note        string

	Hidden  bool
	Deleted bool

	PostedAt  time.Time
	UpdatedAt time.Time
}

type Price struct {
	ID              uint `gorm:"primary_key"`
	ListingRecordID string
	Price           uint

	RecordedAt time.Time
}

type Image struct {
	ID              uint `gorm:"primary_key"`
	ListingRecordID string
	URL             string
}
