package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_createMinefield(t *testing.T) {
	settingsMock := Settings{Height: 9, Width: 6, BombsQuantity: 3}

	minefield := createMinefield(settingsMock)

	assert.Equal(t, settingsMock.Width, len(minefield))
	assert.Equal(t, settingsMock.Height, len(minefield[0]))
	assert.Equal(t, Field{value: (*int)(nil)}, minefield[0][0])
}

func TestNewGame(t *testing.T) {
	settingsMock := Settings{Height: 9, Width: 6, BombsQuantity: 3}
	minefield := createMinefield(settingsMock)

	game, err := NewGame(settingsMock)

	assert.Nil(t, err)
	assert.WithinDuration(t, time.Now(), game.StartedAt, 1*time.Second)
	assert.Equal(t, settingsMock, game.Settings)
	assert.Equal(t, minefield, game.Minefield)
	assert.Equal(t, minefield, game.Minefield)
}

func TestNewGame_err(t *testing.T) {
	settingsMock := Settings{Height: 9, Width: 6, BombsQuantity: 0}

	game, err := NewGame(settingsMock)

	assert.Nil(t, game)
	assert.NotNil(t, err)
}


