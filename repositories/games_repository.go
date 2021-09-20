package repositories

import (
	"github.com/google/uuid"
	"minesweeper-api/databases"
	"minesweeper-api/errors"
	"minesweeper-api/helpers"
	"minesweeper-api/models"
)

type GamesRepository interface {
	Create(game *models.Game) *errors.ApiError
	FindById(uuid *uuid.UUID, hasToPreload bool) (*models.Game, *errors.ApiError)
	Update(game *models.Game) *errors.ApiError
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

func (g gamesRepositoryImpl) FindById(id *uuid.UUID, hasToPreload bool) (*models.Game, *errors.ApiError) {
	var game models.Game
	game.ID = *id

	database := g.database.Get()
	if hasToPreload {
		database = database.Preload("Minefield").Preload("Settings")
	}
	tx := database.Find(&game).Last(&game)
	baseMessage := "game with uuid " + id.String()
	if err := helpers.ValidateDatabaseTxError(tx.Error, baseMessage); err != nil {
		return nil, err
	}

	return &game, nil
}

func (g gamesRepositoryImpl) Update(game *models.Game) *errors.ApiError {
	tx := g.database.Get().Save(&game)
	baseMessage := "game with uuid " + game.ID.String()
	if err := helpers.ValidateDatabaseTxError(tx.Error, baseMessage); err != nil {
		return err
	}

	return nil
}
