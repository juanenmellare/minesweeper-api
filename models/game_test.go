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

func Test_fillMinefieldWithMines(t *testing.T) {
	settingsMock := &Settings{Height: 3, Width: 3, MinesQuantity: 3}
	minefield := make([][]Field, settingsMock.Height)
	for index := range minefield {
		minefield[index] = make([]Field, settingsMock.Width)
	}

	minesPositions := fillMinefieldWithMines(&minefield, settingsMock)

	assert.Equal(t, settingsMock.MinesQuantity, len(minesPositions))
}

func Test_fillMinefieldWithHints(t *testing.T) {
	settingsMock := &Settings{Height: 3, Width: 2, MinesQuantity: 2}
	minefield := make([][]Field, settingsMock.Height)
	for index := range minefield {
		minefield[index] = make([]Field, settingsMock.Width)
	}

	positionOne := Position{Y: 0, X: 1}
	positionTwo := Position{Y: 0, X: 0}

	minefield[positionOne.Y][positionOne.X].SetMine()
	minefield[positionTwo.Y][positionTwo.X].SetMine()

	fillMinefieldWithHints(&minefield, settingsMock, []Position{positionOne, positionTwo})

	assert.True(t, minefield[positionOne.Y][positionOne.X].IsMine())
	assert.True(t, minefield[positionTwo.Y][positionTwo.X].IsMine())
	assert.Equal(t, "2", *minefield[1][0].Value)
	assert.Equal(t, "2", *minefield[1][1].Value)
	assert.Nil(t, minefield[2][0].Value)
	assert.Nil(t, minefield[2][1].Value)
}
