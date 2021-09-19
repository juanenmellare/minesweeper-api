package models

import (
	"minesweeper-api/errors"
	"strconv"
)

type Settings struct {
	Width         int `json:"width"`
	Height        int `json:"height"`
	BombsQuantity int `json:"bombsQuantity"`
}

func buildSettingsMinError(fieldName string, minValue int) *errors.ApiError {
	err := errors.NewError("minefield " + fieldName + " must not be less than " + strconv.Itoa(minValue))
	return errors.NewBadRequestApiError(err)
}

func (s Settings) Validate() *errors.ApiError {
	const minWidth int = 3
	const minHeight int = 3
	const minBombsQuantity int = 1

	if minWidth > s.Width {
		return buildSettingsMinError("width", minHeight)
	}
	if minHeight > s.Height {
		return buildSettingsMinError("height", minHeight)
	}
	if minBombsQuantity > s.BombsQuantity {
		return buildSettingsMinError("bombs quantity", minBombsQuantity)
	}

	return nil
}
