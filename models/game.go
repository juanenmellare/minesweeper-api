package models

import (
	"math/rand"
	"minesweeper-api/errors"
	"time"
)

type Game struct {
	StartedAt time.Time  `json:"startedAt"`
	Settings  Settings   `json:"settings"`
	Minefield *[][]Field `json:"minefield"`
	Status    Status     `json:"status"`
}

func fillMinefieldWithMines(minefield *[][]Field, settings *Settings) {
	bombsCounter := settings.MinesQuantity

	for ok := true; ok; ok = bombsCounter != 0 {
		yPosition := rand.Intn(settings.Height)
		xPosition := rand.Intn(settings.Width)

		if field := &(*minefield)[yPosition][xPosition]; field.Value == nil {
			mineString := "MINE"
			field.Value = &mineString
			bombsCounter--
		}
	}
}

func createMinefield(settings *Settings) *[][]Field {
	minefield := make([][]Field, settings.Height)
	for index := range minefield {
		minefield[index] = make([]Field, settings.Width)
	}

	fillMinefieldWithMines(&minefield, settings)

	return &minefield
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
