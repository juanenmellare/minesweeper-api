package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthChecksController interface {
	Ping(context *gin.Context)
}

type healthChecksControllerImpl struct{}

func NewHealthChecksController() HealthChecksController {
	return &healthChecksControllerImpl{}
}

func (h healthChecksControllerImpl) Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
