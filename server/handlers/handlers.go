package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"go.roman.zone/craigslist-monitor/data"
)

func GetListingsHandler(w http.ResponseWriter, r *http.Request) {

	db, err := data.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dataClient := data.DataClient{
		Database: db,
	}

	records := dataClient.GetAllListingRecords()

	b, err := json.Marshal(convertListingRecords(records))
	if err != nil {
		http.Error(w, "Internal error.", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

type HideInput struct {
	ListingID string `json:"listingId"`
}

func HideListingHandler(w http.ResponseWriter, r *http.Request) {

	var p HideInput

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := data.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dataClient := data.DataClient{
		Database: db,
	}

	if len(p.ListingID) < 1 {
		http.Error(w, "You have to specify `listingId`.", http.StatusBadRequest)
		return
	}

	dataClient.Hide(p.ListingID)

	w.WriteHeader(http.StatusOK)
}

func UnhideListingHandler(w http.ResponseWriter, r *http.Request) {

	var p HideInput

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := data.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dataClient := data.DataClient{
		Database: db,
	}

	if len(p.ListingID) < 1 {
		http.Error(w, "You have to specify `listingId`.", http.StatusBadRequest)
		return
	}

	dataClient.Unhide(p.ListingID)

	w.WriteHeader(http.StatusOK)
}

func SetNoteHandler(w http.ResponseWriter, r *http.Request) {

	type SetNoteInput struct {
		ListingID string `json:"listingId"`
		Note      string `json:"note"`
	}

	var p SetNoteInput

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := data.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dataClient := data.DataClient{
		Database: db,
	}

	if len(p.ListingID) < 1 {
		http.Error(w, "You have to specify `listingId`.", http.StatusBadRequest)
		return
	}

	dataClient.SetNote(p.ListingID, p.Note)

	w.WriteHeader(http.StatusOK)
}
