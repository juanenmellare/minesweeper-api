package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_createMinefield(t *testing.T) {
	settingsMock := &Settings{Height: 9, Width: 6, MinesQuantity: 3}

	minefield := *createMinefield(settingsMock)

	assert.Equal(t, settingsMock.Height, len(minefield))
	assert.Equal(t, settingsMock.Width, len(minefield[0]))
	assert.Equal(t, Field{Value: (*string)(nil)}, minefield[0][0])
}

func TestNewGame(t *testing.T) {
	settingsMock := &Settings{Height: 9, Width: 6, MinesQuantity: 3}
	minefield := *createMinefield(settingsMock)

	game, err := NewGame(settingsMock)

	assert.Nil(t, err)
	assert.WithinDuration(t, time.Now(), game.StartedAt, 1*time.Second)
	assert.Equal(t, settingsMock, &game.Settings)
	assert.Equal(t, settingsMock.Height, len(minefield))
	assert.Equal(t, settingsMock.Width, len(minefield[0]))
}

func TestNewGame_err(t *testing.T) {
	settingsMock := &Settings{Height: 9, Width: 6, MinesQuantity: 0}

	game, err := NewGame(settingsMock)

	assert.Nil(t, game)
	assert.NotNil(t, err)
}

func Test_fillMinefieldWithBombs(t *testing.T) {
	settingsMock := &Settings{Height: 3, Width: 3, MinesQuantity: 1}
	minefield := make([][]Field, settingsMock.Height)
	for index := range minefield {
		minefield[index] = make([]Field, settingsMock.Width)
	}

	fillMinefieldWithMines(&minefield, settingsMock)

	minesCounter := 0
	for _, fieldLine := range minefield {
		for _, field := range fieldLine {
			if field.Value != nil && *field.Value == "MINE" {
				minesCounter++
			}
		}
	}

	assert.Equal(t, settingsMock.MinesQuantity, minesCounter)
}
