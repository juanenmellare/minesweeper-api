package services

import (
	"minesweeper-api/errors"
	"minesweeper-api/models"
)

type GamesService interface {
	Create(settings *models.Settings) (*models.Game, *errors.ApiError)
}

type gamesServiceImpl struct{}

func NewGamesService() GamesService {
	return &gamesServiceImpl{}
}

func (g gamesServiceImpl) Create(settings *models.Settings) (*models.Game, *errors.ApiError) {
	game, err := models.NewGame(settings)
	if err != nil {
		return nil, err
	}

	return game, nil
}
