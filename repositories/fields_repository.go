package repositories

import (
	"github.com/google/uuid"
	"minesweeper-api/databases"
	"minesweeper-api/errors"
	"minesweeper-api/helpers"
	"minesweeper-api/models"
)

type FieldsRepository interface {
	FindByIdAndGameId(uuid *uuid.UUID, gameUuid *uuid.UUID) (*models.Field, *errors.ApiError)
	Update(field *models.Field) *errors.ApiError
}

type fieldsRepositoryImpl struct {
	database databases.RelationalDatabase
}

func NewFieldsRepository(database databases.RelationalDatabase) FieldsRepository {
	return &fieldsRepositoryImpl{
		database: database,
	}
}

func (f fieldsRepositoryImpl) FindByIdAndGameId(uuid *uuid.UUID, gameUuid *uuid.UUID) (
	*models.Field, *errors.ApiError) {
	var field models.Field

	tx := f.database.Get().
		Where(map[string]interface{}{"id": *uuid, "game_id": *gameUuid}).
		Find(&field).Last(&field)
	baseMessage := "field with uuid " + uuid.String() + " and game uuid " + gameUuid.String()
	if err := helpers.ValidateDatabaseTxError(tx.Error, baseMessage); err != nil {
		return nil, err
	}

	return &field, nil
}

func (f fieldsRepositoryImpl) Update(field *models.Field) *errors.ApiError {
	tx := f.database.Get().Save(&field)
	baseMessage := "field with uuid " + field.ID.String() + " and " + "game uuid " + field.GameId.String()
	if err := helpers.ValidateDatabaseTxError(tx.Error, baseMessage); err != nil {
		return err
	}

	return nil
}
