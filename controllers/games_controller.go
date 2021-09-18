package controllers

import (
	"github.com/gin-gonic/gin"
	"minesweeper-api/errors"
	"minesweeper-api/models"
	"net/http"
)

type GamesController interface {
	Create(context *gin.Context)
}

type gamesControllerImpl struct{}

func NewGamesController() GamesController {
	return &gamesControllerImpl{}
}

func (g gamesControllerImpl) Create(context *gin.Context) {
	var settings models.Settings
	if err := context.BindJSON(&settings); err != nil {
		apiErr := errors.NewBadRequestApiError(err)
		context.JSON(apiErr.StatusCode, apiErr)
		return
	}

	game := &models.Game{}

	context.JSON(http.StatusCreated, game)
}
