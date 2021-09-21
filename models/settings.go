package models

import (
	"github.com/google/uuid"
	"minesweeper-api/errors"
	"strconv"
)

type Settings struct {
	ID            uuid.UUID `json:"-" gorm:"type:uuid;default:uuid_generate_v4()"`
	Width         int       `json:"width"`
	Height        int       `json:"height"`
	MinesQuantity int       `json:"minesQuantity"`
}

func buildSettingsMinMaxError(fieldName string, minValue, maxValue int) *errors.ApiError {
	err := errors.NewError("minefield " + fieldName + " must be bigger than " + strconv.Itoa(minValue) +
		" and less than " + strconv.Itoa(maxValue))
	return errors.NewBadRequestApiError(err)
}

func (s Settings) Validate() *errors.ApiError {
	const minWidth int = 3
	const maxWidth int = 30
	if minWidth > s.Width || maxWidth < s.Width {
		return buildSettingsMinMaxError("width", minWidth, maxWidth)
	}

	const minHeight int = 3
	const maxHeight int = 16
	if minHeight > s.Height || maxHeight < s.Height {
		return buildSettingsMinMaxError("height", minHeight, maxHeight)
	}

	const minMinesQuantity int = 1
	maxMinesQuantity := (s.Height * s.Width) - 1
	if minMinesQuantity > s.MinesQuantity || maxMinesQuantity < s.MinesQuantity {
		return buildSettingsMinMaxError("mines quantity", minMinesQuantity, maxMinesQuantity)
	}

	return nil
}
