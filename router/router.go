package router

import (
	"github.com/gin-gonic/gin"
	"minesweeper-api/factories"
)

func New(domainLayers factories.DomainLayersFactory) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", domainLayers.GetHealthChecksController().Ping)

	gamesController := domainLayers.GetGamesController()
	v1Group := router.Group("/v1")
	{
		v1Group.POST("/games", gamesController.Create)
	}

	return router
}
