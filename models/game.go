package models

import (
	"math/rand"
	"minesweeper-api/errors"
	"time"
)

var borderingPositions = []Position{
	{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: -1},
	{X: 0, Y: -1}, {X: -1, Y: -1}, {X: -1, Y: 0}, {X: -1, Y: 1},
}

type Game struct {
	StartedAt time.Time  `json:"startedAt"`
	Settings  Settings   `json:"settings"`
	Minefield *[][]Field `json:"minefield"`
	Status    Status     `json:"status"`
}

func fillMinefieldWithMines(minefield *[][]Field, settings *Settings) []Position {
	bombsCounter := settings.MinesQuantity

	minesPositions := make([]Position, bombsCounter)

	for ok := true; ok; ok = bombsCounter != 0 {
		yPosition := rand.Intn(settings.Height)
		xPosition := rand.Intn(settings.Width)

		if field := &(*minefield)[yPosition][xPosition]; field.Value == nil {
			field.SetMine()
			minesPositions[bombsCounter-1] = Position{Y: yPosition, X: xPosition}
			bombsCounter--
		}
	}

	return minesPositions
}

func fillMinefieldWithHints(minefield *[][]Field, settings *Settings, minesPositions []Position) {
	for _, position := range minesPositions {
		for _, borderingPosition := range borderingPositions {
			candidatePositionY := position.Y + borderingPosition.Y
			candidatePositionX := position.X + borderingPosition.X

			isCandidateInsideMinefield :=
				candidatePositionY < settings.Height && candidatePositionY != -1 &&
					candidatePositionX < settings.Width && candidatePositionX != -1

			if isCandidateInsideMinefield {
				field := &(*minefield)[candidatePositionY][candidatePositionX]
				if field.Value == nil {
					field.SetInitialHintValue()
				} else if !field.IsMine() {
					field.IncrementHintValue()
				}
			}
		}
	}
}

func createMinefield(settings *Settings) *[][]Field {
	minefield := make([][]Field, settings.Height)
	for index := range minefield {
		minefield[index] = make([]Field, settings.Width)
	}

	minesPositions := fillMinefieldWithMines(&minefield, settings)
	fillMinefieldWithHints(&minefield, settings, minesPositions)

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
