package router

import (
	"github.com/gin-gonic/gin"
	"minesweeper-api/controllers"
)

func New() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controllers.NewHealthChecksController().Ping)

	v1Group := router.Group("/v1")
	{
		v1Group.POST("/games", controllers.NewGamesController().Create)
	}

	return router
}
