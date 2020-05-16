package parser

import (
	"io"
	"strconv"

	"golang.org/x/net/html"
)

func GetPrice(pageContent io.Reader) (int, error) {
	doc, err := html.Parse(pageContent)
	if err != nil {
		return 0, err
	}
	return findPrice(doc)
}

func findPrice(n *html.Node) (int, error) {
	if n.Type == html.ElementNode && isPriceElement(n) {
		return parsePrice(n.FirstChild.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		price, err := findPrice(c)
		if err == nil {
			return price, nil
		}
	}

	return 0, &NoPriceError{}
}

func isPriceElement(n *html.Node) bool {
	if n.Data == "span" {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "price" {
				return true
			}
		}
	}
	return false
}

func parsePrice(data string) (int, error) {
	if len(data) < 2 {
		return 0, &NoPriceError{}
	}

	value, err := strconv.Atoi(data[1:])
	if err != nil {
		return 0, &NoPriceError{}
	}

	return value, nil
}
