package factories

import (
	"minesweeper-api/controllers"
	"minesweeper-api/services"
)

type DomainLayersFactory interface {
	createHealthChecksController()
	GetHealthChecksController() controllers.HealthChecksController
	createGamesController()
	GetGamesController() controllers.GamesController
}

type domainLayersFactoryImpl struct {
	healthChecksController controllers.HealthChecksController
	gamesController controllers.GamesController
}

func NewDomainLayersFactory() DomainLayersFactory {
	domainLayersFactory := &domainLayersFactoryImpl{}

	domainLayersFactory.createHealthChecksController()
	domainLayersFactory.createGamesController()

	return domainLayersFactory
}

func (d *domainLayersFactoryImpl) createHealthChecksController() {
	healthChecksController := controllers.NewHealthChecksController()

	d.healthChecksController = healthChecksController
}

func (d domainLayersFactoryImpl) GetHealthChecksController() controllers.HealthChecksController {
	return d.healthChecksController
}

func (d *domainLayersFactoryImpl) createGamesController() {
	gamesService := services.NewGamesService()
	gamesController := controllers.NewGamesController(gamesService)

	d.gamesController = gamesController
}

func (d domainLayersFactoryImpl) GetGamesController() controllers.GamesController {
	return d.gamesController
}
