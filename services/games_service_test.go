package services

import (
	"github.com/stretchr/testify/assert"
	"minesweeper-api/models"
	"testing"
	"time"
)

func TestNewGamesService(t *testing.T) {
	gameService := NewGamesService()

	assert.Implements(t, (*GamesService)(nil), gameService)
}

func Test_gamesServiceImpl_Create(t *testing.T) {
	gameService := NewGamesService()

	settings := &models.Settings{Width: 3, Height: 3, MinesQuantity: 1}

	game, err := gameService.Create(settings)

	gameExpected, _ := models.NewGame(settings)

	assert.WithinDuration(t, gameExpected.StartedAt, game.StartedAt, 1*time.Second)
	assert.Equal(t, gameExpected.Status, game.Status)
	assert.Equal(t, settings.Height, len(*game.Minefield))
	assert.Equal(t, settings.Height, len((*game.Minefield)[0]))
	assert.Equal(t, gameExpected.Settings, game.Settings)
	assert.Nil(t, err)
}

func Test_gamesServiceImpl_Create_new_game_err(t *testing.T) {
	gameService := NewGamesService()

	settings := &models.Settings{Width: 0, Height: 3, MinesQuantity: 1}

	game, err := gameService.Create(settings)

	assert.Nil(t, game)
	assert.NotNil(t, err)
}
