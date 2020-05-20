package retriever

import (
	"net/http"

	"go.roman.zone/craig/parser"
)

func GetPrice(url string) (int, error) {

	respose, err := getPage(url)
	if err != nil {
		return 0, err
	}
	defer respose.Body.Close()

	return parser.ExtractPrice(respose.Body)
}

func getPage(url string) (*http.Response, error) {
	return http.Get(url)
}
