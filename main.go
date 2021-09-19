package main

import (
	"log"
	"minesweeper-api/factories"
	"minesweeper-api/router"
)

func main() {
	logger := log.Default()

	domainLayersFactory := factories.NewDomainLayersFactory()

	port := ":8080"
	if err := router.New(domainLayersFactory).Run(port); err != nil {
		logger.Fatalf("Error while trying to create the router: " + err.Error())
	}
}
