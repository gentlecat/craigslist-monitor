package server

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"

	api_handlers "go.roman.zone/craigslist-monitor/server/handlers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func CreateServer() *http.Server {
	router := makeRouter()
	loggingRouter := handlers.LoggingHandler(os.Stdout, router)

	return &http.Server{
		Handler: loggingRouter,
		Addr:    "localhost:8080",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func makeRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", indexHandler)

	r.HandleFunc("/api/list", api_handlers.GetListingsHandler).Methods("GET")
	r.HandleFunc("/api/hide", api_handlers.HideListingHandler).Methods("POST")
	r.HandleFunc("/api/unhide", api_handlers.UnhideListingHandler).Methods("POST")
	r.HandleFunc("/api/note", api_handlers.SetNoteHandler).Methods("POST")

	// Static files
	// TODO: Get static dir path from outside
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("../frontend/static"))))

	return r
}

// executeTemplates is a custom template executor that uses our template
// structure. Should be used when rendering templates based on "base.html"
// template.
func executeTemplates(wr io.Writer, data interface{}, filenames ...string) error {
	filenames = append(filenames, "../frontend/templates/base.html")
	t, err := template.ParseFiles(filenames...)
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(wr, "base", data)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := executeTemplates(w, struct{}{}, "../frontend/templates/index.html")
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
}
