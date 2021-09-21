package router

import (
	"github.com/gin-gonic/gin"
	"minesweeper-api/factories"
)

func New(domainLayers factories.DomainLayersFactory) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", domainLayers.GetHealthChecksController().Ping)

	gamesController := domainLayers.GetGamesController()
	v1 := router.Group("/v1")

	games := v1.Group("/games")
	games.POST("/", gamesController.Create)

	uuid := games.Group("/:game-uuid")
	uuid.GET("/", gamesController.FindById)

	uuid.PUT("/fields/:field-uuid/:action", gamesController.ExecuteFieldAction)
	return router
}
