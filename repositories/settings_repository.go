package repositories

import (
	"github.com/google/uuid"
	"minesweeper-api/databases"
	"minesweeper-api/errors"
	"minesweeper-api/helpers"
	"minesweeper-api/models"
)

type SettingsRepository interface {
	FindByGameId(gameUuid *uuid.UUID) (*models.Settings, *errors.ApiError)
}

type settingsRepositoryImpl struct {
	database databases.RelationalDatabase
}

func NewSettingsRepository(database databases.RelationalDatabase) SettingsRepository {
	return &settingsRepositoryImpl{
		database: database,
	}
}

func (s settingsRepositoryImpl) FindByGameId(gameUuid *uuid.UUID) (*models.Settings, *errors.ApiError) {
	var settings models.Settings
	tx := s.database.Get().Where(map[string]interface{}{"game_id": *gameUuid}).Find(&settings).Last(&settings)
	baseMessage := "settings with game uuid " + gameUuid.String()
	if err := helpers.ValidateDatabaseTxError(tx.Error, baseMessage); err != nil {
		return nil, err
	}

	return &settings, nil
}
