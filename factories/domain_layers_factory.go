package factories

import (
	"minesweeper-api/controllers"
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

func NewDomainLayersFactory() DomainLayersFactory {
	return &domainLayersFactoryImpl{
		healthChecksController: createHealthChecksController(),
		gamesController:        createGamesController(),
	}
}

func createHealthChecksController() controllers.HealthChecksController {
	return controllers.NewHealthChecksController()
}

func createGamesController() controllers.GamesController {
	gamesService := services.NewGamesService()
	gamesController := controllers.NewGamesController(gamesService)

	return gamesController
}

func (d domainLayersFactoryImpl) GetHealthChecksController() controllers.HealthChecksController {
	return d.healthChecksController
}

func (d domainLayersFactoryImpl) GetGamesController() controllers.GamesController {
	return d.gamesController
}
