package data

import "github.com/jinzhu/gorm"

const (
	dbFileName = "data.sqlite3"
)

type DataClient struct {
	Database gorm.DB
}

// OpenDB opens a new connection to the database.
//
// Don't forget to close it after your are done using it!
func OpenDB() (db *gorm.DB, err error) {
	return gorm.Open("sqlite3", dbFileName)
}

func (dbClient *DataClient) Init() {
	dbClient.Database.AutoMigrate(&ListingRecord{}, &Price{})
}
