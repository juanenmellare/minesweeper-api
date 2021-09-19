package models

import (
	"minesweeper-api/errors"
	"time"
)

type Game struct {
	StartedAt time.Time `json:"startedAt"`
	Settings  Settings `json:"settings"`
	Minefield [][]Field `json:"minefield"`
	Status    Status `json:"status"`
}

func createMinefield(settings *Settings) [][]Field {
	minefield := make([][]Field, settings.Width)
	for index := range minefield {
		minefield[index] = make([]Field, settings.Height)
	}

	return minefield
}

func NewGame(settings *Settings) (*Game, *errors.ApiError) {
	if err := settings.Validate(); err != nil {
		return nil, err
	}

	return &Game{
		StartedAt: time.Now(),
		Settings:  *settings,
		Minefield: createMinefield(settings),
		Status:    StatusInProgress,
	}, nil
}
