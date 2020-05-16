package retriever

import (
	"log"
	"net/http"
)

func GetPage(url string) (*http.Response, error) {
	log.Printf("Retrieving %s\n", url)
	return http.Get(url)
}
