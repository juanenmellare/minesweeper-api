package controllers

import (
	"github.com/gin-gonic/gin"
	"minesweeper-api/errors"
	"minesweeper-api/models"
	"minesweeper-api/services"
	"net/http"
)

type GamesController interface {
	Create(context *gin.Context)
}

type gamesControllerImpl struct{
	gamesService services.GamesService
}

func NewGamesController(gamesService services.GamesService) GamesController {
	return &gamesControllerImpl{
		gamesService: gamesService,
	}
}

func (g gamesControllerImpl) Create(context *gin.Context) {
	var settings models.Settings
	if err := context.BindJSON(&settings); err != nil {
		apiErr := errors.NewBadRequestApiError(err)
		context.JSON(apiErr.StatusCode, apiErr)
		return
	}

	game, err := g.gamesService.Create(&settings)
	if err != nil {
		context.JSON(err.StatusCode, err)
		return
	}

	context.JSON(http.StatusCreated, game)
}
