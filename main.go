package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"go.roman.zone/craig/retriever"
)

func main() {

	file, err := os.Open("urls.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		url := scanner.Text()

		price, err := retriever.GetPrice(url)
		if err != nil {
			panic(err)
		}

		fmt.Printf("$%d: %s\n", price, url)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done!")
}
