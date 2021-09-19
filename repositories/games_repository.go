package repositories

import (
	"minesweeper-api/databases"
	"minesweeper-api/errors"
	"minesweeper-api/models"
)

//go:generate mockery --name GamesRepository --output mocks
type GamesRepository interface {
	Create(game *models.Game) *errors.ApiError
}

type gamesRepositoryImpl struct {
	database databases.RelationalDatabase
}

func NewGamesRepository(database databases.RelationalDatabase) GamesRepository {
	return &gamesRepositoryImpl{
		database: database,
	}
}

func (g gamesRepositoryImpl) Create(game *models.Game) *errors.ApiError {
	tx := g.database.Get().Create(&game)
	if err := tx.Error; err != nil {
		return errors.NewInternalServerApiError(err)
	}

	return nil
}
