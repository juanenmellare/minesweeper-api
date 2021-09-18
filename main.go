package main

import (
	"log"
	"minesweeper-api/router"
)

func main() {
	logger := log.Default()

	port := ":8080"
	if err := router.New().Run(port); err != nil {
		logger.Fatalf("Error while trying to create the router: " + err.Error())
	}
}
