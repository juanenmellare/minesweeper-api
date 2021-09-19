package services

import (
	"minesweeper-api/errors"
	"minesweeper-api/models"
	"minesweeper-api/repositories"
)

type GamesService interface {
	Create(settings *models.Settings) (*models.Game, *errors.ApiError)
}

type gamesServiceImpl struct {
	gamesRepository repositories.GamesRepository
}

func NewGamesService(gamesRepository repositories.GamesRepository) GamesService {
	return &gamesServiceImpl{
		gamesRepository: gamesRepository,
	}
}

func (g gamesServiceImpl) Create(settings *models.Settings) (*models.Game, *errors.ApiError) {
	game, err := models.NewGame(settings)
	if err != nil {
		return nil, err
	}

	if err = g.gamesRepository.Create(game); err != nil {
		return nil, err
	}

	return game, nil
}
