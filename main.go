package main

import (
	"fmt"

	"go.roman.zone/craig/parser"
	"go.roman.zone/craig/retriever"
)

func main() {
	url := "PUT YOUR URL HERE"

	respose, err := retriever.GetPage(url)
	if err != nil {
		panic(err)
	}
	defer respose.Body.Close()

	price, err := parser.GetPrice(respose.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(price)
}
