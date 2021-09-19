package models

import (
	"github.com/google/uuid"
	"math/rand"
	"minesweeper-api/errors"
	"time"
)

var borderingPositions = []Position{
	{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: -1},
	{X: 0, Y: -1}, {X: -1, Y: -1}, {X: -1, Y: 0}, {X: -1, Y: 1},
}

type Game struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	StartedAt  time.Time `json:"startedAt"`
	Settings   Settings  `json:"settings" gorm:"foreignKey:id"`
	SettingsID uuid.UUID `json:"-" gorm:"foreignKey:ID;references:SettingsID;constraint:OnUpdate:CASCADE"`
	Minefield  []Field   `json:"minefield" gorm:"GameId"`
	Status     Status    `json:"status"`
}

func fillMinefieldWithMines(minefield *[][]Field, settings *Settings) []Position {
	bombsCounter := settings.MinesQuantity

	minesPositions := make([]Position, bombsCounter)

	for ok := true; ok; ok = bombsCounter != 0 {
		yPosition := rand.Intn(settings.Height)
		xPosition := rand.Intn(settings.Width)

		if field := &(*minefield)[yPosition][xPosition]; field.IsNil() {
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
				if field.IsNil() {
					field.SetInitialHintValue()
				} else if !field.IsMine() {
					field.IncrementHintValue()
				}
			}
		}
	}
}

func fillFieldsPositionsAndFlat(settings *Settings, minefield *[][]Field) []Field {
	minefieldSlice := make([]Field, settings.Height*settings.Width)

	for y := 0; y < settings.Height; y++ {
		for x := 0; x < settings.Width; x++ {
			field := &(*minefield)[y][x]
			field.PositionY = y
			field.PositionX = x
			minefieldSlice[y+x] = *field
		}
	}

	return minefieldSlice
}

func createMinefield(settings *Settings) *[]Field {
	minefield := make([][]Field, settings.Height)
	for index := range minefield {
		minefield[index] = make([]Field, settings.Width)
	}

	minesPositions := fillMinefieldWithMines(&minefield, settings)
	fillMinefieldWithHints(&minefield, settings, minesPositions)
	minefieldFlatted := fillFieldsPositionsAndFlat(settings, &minefield)

	return &minefieldFlatted
}

func NewGame(settings *Settings) (*Game, *errors.ApiError) {
	if err := settings.Validate(); err != nil {
		return nil, err
	}

	return &Game{
		StartedAt: time.Now(),
		Settings:  *settings,
		Minefield: *createMinefield(settings),
		Status:    StatusInProgress,
	}, nil
}
