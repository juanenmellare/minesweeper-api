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
	minefield := make([]models.Field, settings.Height*settings.Width)
	gameExpected := &models.Game{
		StartedAt: time.Now(), Settings: *settings, Minefield: &minefield, Status: models.GameStatusInProgress,
	}

	gameService := NewGamesService(gamesRepositoryMock)

	game, err := gameService.Create(settings)

	assert.WithinDuration(t, gameExpected.StartedAt, game.StartedAt, 1*time.Second)
	assert.Equal(t, gameExpected.Status, game.Status)
	assert.Equal(t, len(minefield), len(*game.Minefield))
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

func Test_fillMinefieldWithMines(t *testing.T) {
	settingsMock := &models.Settings{Height: 3, Width: 3, MinesQuantity: 3}
	minefield := make([][]models.Field, settingsMock.Height)
	for index := range minefield {
		minefield[index] = make([]models.Field, settingsMock.Width)
	}

	minesPositions := fillMinefieldWithMines(&minefield, settingsMock)

	assert.Equal(t, settingsMock.MinesQuantity, len(minesPositions))
}

func Test_fillMinefieldWithHints(t *testing.T) {
	settingsMock := &models.Settings{Height: 3, Width: 2, MinesQuantity: 2}

	minefield := make([][]models.Field, settingsMock.Height)
	for index := range minefield {
		minefield[index] = make([]models.Field, settingsMock.Width)
	}

	positionOne := models.Position{Y: 0, X: 1}
	positionTwo := models.Position{Y: 0, X: 0}

	minefield[positionOne.Y][positionOne.X].SetMine()
	minefield[positionTwo.Y][positionTwo.X].SetMine()

	fillMinefieldWithHints(&minefield, settingsMock, []models.Position{positionOne, positionTwo})

	assert.True(t, minefield[positionOne.Y][positionOne.X].IsMine())
	assert.True(t, minefield[positionTwo.Y][positionTwo.X].IsMine())
	assert.Equal(t, "2", *minefield[1][0].Value)
	assert.Equal(t, "2", *minefield[1][1].Value)
	assert.Nil(t, minefield[2][0].Value)
	assert.Nil(t, minefield[2][1].Value)
}

func Test_fillFieldsPositionsAndFlat(t *testing.T) {
	settingsMock := &models.Settings{Height: 3, Width: 2, MinesQuantity: 2}

	minefield := make([][]models.Field, settingsMock.Height)
	for index := range minefield {
		minefield[index] = make([]models.Field, settingsMock.Width)
	}

	minefieldFlatted := *fillFieldsPositionsAndFlat(settingsMock, &minefield)

	assert.Equal(t, 0, minefieldFlatted[0].PositionY)
	assert.Equal(t, 0, minefieldFlatted[0].PositionX)
	assert.Equal(t, models.FieldStatusHidden, minefieldFlatted[0].Status)
	assert.Equal(t, 0, minefieldFlatted[1].PositionY)
	assert.Equal(t, 1, minefieldFlatted[1].PositionX)
	assert.Equal(t, models.FieldStatusHidden, minefieldFlatted[1].Status)
	assert.Equal(t, 1, minefieldFlatted[2].PositionY)
	assert.Equal(t, 0, minefieldFlatted[2].PositionX)
	assert.Equal(t, models.FieldStatusHidden, minefieldFlatted[2].Status)
	assert.Equal(t, 1, minefieldFlatted[3].PositionY)
	assert.Equal(t, 1, minefieldFlatted[3].PositionX)
	assert.Equal(t, models.FieldStatusHidden, minefieldFlatted[3].Status)
	assert.Equal(t, 2, minefieldFlatted[4].PositionY)
	assert.Equal(t, 0, minefieldFlatted[4].PositionX)
	assert.Equal(t, models.FieldStatusHidden, minefieldFlatted[4].Status)
	assert.Equal(t, 2, minefieldFlatted[5].PositionY)
	assert.Equal(t, 1, minefieldFlatted[5].PositionX)
	assert.Equal(t, models.FieldStatusHidden, minefieldFlatted[5].Status)

}
