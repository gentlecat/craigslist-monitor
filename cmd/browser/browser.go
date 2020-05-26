package main

import (
	"fmt"
	"log"

	"go.roman.zone/craigslist-monitor/server"
)

func main() {
	fmt.Println("Starting server on http://localhost:8080")
	err := server.CreateServer().ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
