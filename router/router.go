package router

import (
	"github.com/gin-gonic/gin"
	"minesweeper-api/controllers"
)

func New() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controllers.NewHealthChecksController().Ping)

	return router
}
