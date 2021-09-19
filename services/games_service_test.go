package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"minesweeper-api/errors"
	"minesweeper-api/models"
	"minesweeper-api/repositories/mocks"
	"testing"
	"time"
)

func TestNewGamesService(t *testing.T) {
	gamesRepositoryMock := new(mocks.GamesRepository)
	gamesRepositoryMock.On("Create", &models.Game{}).Return(nil)

	gameService := NewGamesService(gamesRepositoryMock)

	assert.Implements(t, (*GamesService)(nil), gameService)
}

func Test_gamesServiceImpl_Create(t *testing.T) {
	gamesRepositoryMock := new(mocks.GamesRepository)
	gamesRepositoryMock.On("Create", mock.AnythingOfType("*models.Game")).Return(nil)

	settings := &models.Settings{Width: 3, Height: 3, MinesQuantity: 1}
	gameExpected, _ := models.NewGame(settings)

	gameService := NewGamesService(gamesRepositoryMock)

	game, err := gameService.Create(settings)

	assert.WithinDuration(t, gameExpected.StartedAt, game.StartedAt, 1*time.Second)
	assert.Equal(t, gameExpected.Status, game.Status)
	assert.Equal(t, settings.Height*settings.Width, len(game.Minefield))
	assert.Equal(t, gameExpected.Settings, game.Settings)
	assert.Nil(t, err)
}

func Test_gamesServiceImpl_Create_new_game_err(t *testing.T) {
	gamesRepositoryMock := new(mocks.GamesRepository)
	gamesRepositoryMock.On("Create", &models.Game{}).Return(nil)

	gameService := NewGamesService(gamesRepositoryMock)

	settings := &models.Settings{Width: 0, Height: 3, MinesQuantity: 1}

	game, err := gameService.Create(settings)

	assert.Nil(t, game)
	assert.NotNil(t, err)
}

func Test_gamesServiceImpl_Create_repository_create_err(t *testing.T) {
	gamesRepositoryMock := new(mocks.GamesRepository)
	errExpected := errors.NewInternalServerApiError(errors.NewError("panic"))
	gamesRepositoryMock.On("Create", mock.AnythingOfType("*models.Game")).Return(errExpected)

	gameService := NewGamesService(gamesRepositoryMock)

	settings := &models.Settings{Width: 3, Height: 3, MinesQuantity: 1}

	game, err := gameService.Create(settings)

	assert.Nil(t, game)
	assert.NotNil(t, err)
	assert.Equal(t, errExpected.StatusCode, err.StatusCode)
	assert.Equal(t, errExpected.Message, err.Message)
}
