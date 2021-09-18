package models

import "time"

type Game struct {
	StartedAt time.Time
	Settings  Settings
	Minefield [][]Field
	Status    Status
}

func createMinefield(settings Settings) [][]Field {
	minefield := make([][]Field, settings.Width)
	for index := range minefield {
		minefield[index] = make([]Field, settings.Height)
	}

	return minefield
}

func NewGame(settings Settings) (*Game, error) {
	if err := settings.Validate(); err != nil {
		return nil, err
	}

	return &Game{
		StartedAt: time.Now(),
		Settings:  settings,
		Minefield: createMinefield(settings),
		Status:    StatusInProgress,
	}, nil
}
