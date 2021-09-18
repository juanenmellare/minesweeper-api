package models

import (
	"errors"
	"strconv"
)

type Settings struct {
	Width         int
	Height        int
	BombsQuantity int
}

func buildSettingsMinError(fieldName string, minValue int) error {
	return errors.New("minefield " + fieldName + " must not be less than " + strconv.Itoa(minValue))
}

func (s Settings) Validate() error {
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
