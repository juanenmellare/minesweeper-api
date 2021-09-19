package factories

import (
	"minesweeper-api/controllers"
	"minesweeper-api/databases"
	"minesweeper-api/repositories"
	"minesweeper-api/services"
)

type DomainLayersFactory interface {
	GetHealthChecksController() controllers.HealthChecksController
	GetGamesController() controllers.GamesController
}

type domainLayersFactoryImpl struct {
	healthChecksController controllers.HealthChecksController
	gamesController        controllers.GamesController
}

func NewDomainLayersFactory(database databases.RelationalDatabase) DomainLayersFactory {
	return &domainLayersFactoryImpl{
		healthChecksController: createHealthChecksController(),
		gamesController:        createGamesController(database),
	}
}

func createHealthChecksController() controllers.HealthChecksController {
	return controllers.NewHealthChecksController()
}

func createGamesController(database databases.RelationalDatabase) controllers.GamesController {
	gamesRepository := repositories.NewGamesRepository(database)
	gamesService := services.NewGamesService(gamesRepository)
	gamesController := controllers.NewGamesController(gamesService)

	return gamesController
}

func (d domainLayersFactoryImpl) GetHealthChecksController() controllers.HealthChecksController {
	return d.healthChecksController
}

func (d domainLayersFactoryImpl) GetGamesController() controllers.GamesController {
	return d.gamesController
}
