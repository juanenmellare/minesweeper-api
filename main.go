package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"minesweeper-api/configs"
	"minesweeper-api/databases"
	"minesweeper-api/factories"
	"minesweeper-api/router"

	"log"
)

func main() {
	logger := log.Default()
	config := configs.NewConfig()

	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		config.GetDatabaseUser(), config.GetDatabasePass(), config.GetDatabaseHost(),
		config.GetDatabasePort(), config.GetDatabaseName())

	relationalDatabase := databases.NewConnection(gorm.Open(postgres.Open(connectionString), &gorm.Config{}))
	relationalDatabase.DoMigration()

	domainLayersFactory := factories.NewDomainLayersFactory(relationalDatabase)

	port := ":" + config.GetPort()
	if err := router.New(domainLayersFactory).Run(port); err != nil {
		logger.Fatalf("Error while trying to create the router: " + err.Error())
	}
}
