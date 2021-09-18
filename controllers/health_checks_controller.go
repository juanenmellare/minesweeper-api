package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthCheckController interface {
	Ping(context *gin.Context)
}

type healthCheckControllerImpl struct{}

func NewHealthCheckController() HealthCheckController {
	return &healthCheckControllerImpl{}
}

func (h healthCheckControllerImpl) Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
