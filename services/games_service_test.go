package services

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"minesweeper-api/errors"
	"minesweeper-api/models"
	"minesweeper-api/repositories/mocks"
	"net/http"
	"testing"
	"time"
)

func TestNewGamesService(t *testing.T) {
	gamesRepositoryMock := new(mocks.GamesRepository)
	gamesRepositoryMock.On("Create", &models.Game{}).Return(nil)

	gameService := NewGamesService(gamesRepositoryMock, nil)

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

	gameService := NewGamesService(gamesRepositoryMock, nil)

	game, err := gameService.Create(settings)

	assert.WithinDuration(t, gameExpected.StartedAt, game.StartedAt, 1*time.Second)
	assert.Equal(t, gameExpected.Status, game.Status)
	assert.Equal(t, len(minefield), len(*game.Minefield))
	assert.Equal(t, gameExpected.Settings, game.Settings)
	assert.Nil(t, err)
}

func Test_gamesServiceImpl_Create_new_game_error(t *testing.T) {
	gamesRepositoryMock := new(mocks.GamesRepository)
	gamesRepositoryMock.On("Create", &models.Game{}).Return(nil)

	gameService := NewGamesService(gamesRepositoryMock, nil)

	settings := &models.Settings{Width: 0, Height: 3, MinesQuantity: 1}

	game, err := gameService.Create(settings)

	assert.Nil(t, game)
	assert.NotNil(t, err)
}

func Test_gamesServiceImpl_Create_repository_create_err(t *testing.T) {
	gamesRepositoryMock := new(mocks.GamesRepository)
	errExpected := errors.NewInternalServerApiError(errors.NewError("panic"))
	gamesRepositoryMock.On("Create", mock.AnythingOfType("*models.Game")).Return(errExpected)

	gameService := NewGamesService(gamesRepositoryMock, nil)

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

func Test_gamesServiceImpl_FindById(t *testing.T) {
	gamesRepositoryMock := new(mocks.GamesRepository)
	gameExpected := &models.Game{
		Minefield: &[]models.Field{{Status: models.FieldStatusHidden, Value: &models.MineString}}}

	uuidParam := uuid.New()
	gamesRepositoryMock.On("FindById", &uuidParam, true).Return(gameExpected, nil)

	gameService := NewGamesService(gamesRepositoryMock, nil)

	game, err := gameService.FindById(&uuidParam, true)

	assert.Equal(t, gameExpected, game)
	assert.Nil(t, err)
}

func Test_gamesServiceImpl_FindById_err(t *testing.T) {
	gamesRepositoryMock := new(mocks.GamesRepository)
	errExpected := errors.NewInternalServerApiError(errors.NewError("panic"))
	uuidParam := uuid.New()
	gamesRepositoryMock.On("FindById", &uuidParam, true).Return(nil, errExpected)

	gameService := NewGamesService(gamesRepositoryMock, nil)

	game, err := gameService.FindById(&uuidParam, true)

	assert.Equal(t, errExpected.Status, err.Status)
	assert.Nil(t, game)
}

func Test_executeAfterShow(t *testing.T) {
	settings := models.Settings{Height: 9, Width: 9, MinesQuantity: 3}
	gameUuid := uuid.New()
	game := &models.Game{ID: gameUuid, Settings: settings}
	field := &models.Field{Status: models.FieldStatusShown}

	fieldsRepository := new(mocks.FieldsRepository)
	fieldsRepository.On("FindByGameId", &gameUuid).Return(&[]models.Field{{}, {}}, nil)

	gameService := gamesServiceImpl{nil, fieldsRepository}

	status, err := executeAfterShow(gameService, field, game)

	assert.Nil(t, status)
	assert.Nil(t, err)
}

func Test_executeAfterShow_FindByGameId_error(t *testing.T) {
	settings := models.Settings{Height: 9, Width: 9, MinesQuantity: 3}
	gameUuid := uuid.New()
	game := &models.Game{ID: gameUuid, Settings: settings}
	field := &models.Field{Status: models.FieldStatusShown}

	fieldsRepository := new(mocks.FieldsRepository)
	errExpected := errors.NewInternalServerApiError(errors.NewError("panic"))
	fieldsRepository.On("FindByGameId", &gameUuid).Return(nil, errExpected)

	gameService := gamesServiceImpl{nil, fieldsRepository}

	status, err := executeAfterShow(gameService, field, game)

	assert.Nil(t, status)
	assert.Equal(t, errExpected, err)
}

func Test_executeAfterShow_status_lost(t *testing.T) {
	field := &models.Field{Status: models.FieldStatusShown, Value: &models.MineString}

	gameService := gamesServiceImpl{}

	status, err := executeAfterShow(gameService, field, nil)

	assert.Equal(t, models.GameStatusLost, *status)
	assert.Nil(t, err)
}

func Test_executeAfterFlag(t *testing.T) {
	settings := models.Settings{MinesQuantity: 3}
	gameUuid := uuid.New()
	game := &models.Game{ID: gameUuid, Settings: settings}
	field := &models.Field{Status: models.FieldStatusShown, GameId: gameUuid}

	fieldsRepository := new(mocks.FieldsRepository)
	fieldsRepository.On("FindMineFieldsFlaggedByGame", &gameUuid).Return(&[]models.Field{{}, {}}, nil)

	gameService := gamesServiceImpl{nil, fieldsRepository}

	status, err := executeAfterFlag(gameService, field, game)

	assert.Nil(t, status)
	assert.Nil(t, err)
}

func Test_executeAfterFlag_status_won(t *testing.T) {
	settings := models.Settings{MinesQuantity: 3}
	gameUuid := uuid.New()
	game := &models.Game{ID: gameUuid, Settings: settings}
	field := &models.Field{Status: models.FieldStatusShown, GameId: gameUuid}

	fieldsRepository := new(mocks.FieldsRepository)
	fieldsRepository.On("FindMineFieldsFlaggedByGame", &gameUuid).Return(&[]models.Field{{}, {}, {}}, nil)

	gamesService := gamesServiceImpl{nil, fieldsRepository}

	status, err := executeAfterFlag(gamesService, field, game)

	assert.Equal(t, models.GameStatusWon, *status)
	assert.Nil(t, err)
}

func Test_executeAfterFlag_status_error(t *testing.T) {
	settings := models.Settings{MinesQuantity: 3}
	gameUuid := uuid.New()
	game := &models.Game{ID: gameUuid, Settings: settings}
	field := &models.Field{Status: models.FieldStatusShown, GameId: gameUuid}

	fieldsRepository := new(mocks.FieldsRepository)
	errExpected := errors.NewInternalServerApiError(errors.NewError("panic"))
	fieldsRepository.On("FindMineFieldsFlaggedByGame", &gameUuid).
		Return(nil, errExpected)

	gamesService := gamesServiceImpl{nil, fieldsRepository}

	status, err := executeAfterFlag(gamesService, field, game)

	assert.Nil(t, status)
	assert.Equal(t, errExpected, err)
}

func Test_gamesServiceImpl_ExecuteFieldAction(t *testing.T) {
	gameUuid := uuid.New()
	game := &models.Game{ID: gameUuid, Status: models.GameStatusInProgress}
	fieldUuid := uuid.New()
	field := &models.Field{ID: fieldUuid, Status: models.FieldStatusHidden, GameId: gameUuid, Value: &models.MineString}

	gamesRepositoryMock := new(mocks.GamesRepository)
	gamesRepositoryMock.On("FindById", &gameUuid, false).Return(game, nil)
	gamesRepositoryMock.On("Update", mock.Anything).Return(nil)

	fieldRepositoryMock := new(mocks.FieldsRepository)
	fieldRepositoryMock.On("FindByIdAndGameId", &fieldUuid, &gameUuid).Return(field, nil)
	fieldRepositoryMock.On("Update", mock.Anything).Return(nil)

	gameService := NewGamesService(gamesRepositoryMock, fieldRepositoryMock)

	fieldStatus := models.FieldStatusShown

	status, err := gameService.ExecuteFieldAction(&gameUuid, &fieldUuid, fieldStatus)

	assert.Equal(t, &game.Status, status)
	assert.NotEqual(t, models.GameStatusInProgress, game.Status)
	assert.NotNil(t, game.EndedAt)
	assert.Nil(t, err)
}

func Test_gamesServiceImpl_ExecuteFieldAction_FindById_error(t *testing.T) {
	gameUuid := uuid.New()
	fieldUuid := uuid.New()

	gamesRepositoryMock := new(mocks.GamesRepository)
	errExpected := errors.NewInternalServerApiError(errors.NewError("panic"))
	gamesRepositoryMock.On("FindById", &gameUuid, false).Return(nil, errExpected)

	gameService := NewGamesService(gamesRepositoryMock, nil)

	fieldStatus := models.FieldStatusShown

	status, err := gameService.ExecuteFieldAction(&gameUuid, &fieldUuid, fieldStatus)

	assert.Nil(t, status)
	assert.Equal(t, errExpected, err)
}

func Test_gamesServiceImpl_ExecuteFieldAction_FindById_notInProgress_error(t *testing.T) {
	gameUuid := uuid.New()
	game := &models.Game{ID: gameUuid, Status: models.GameStatusLost}
	fieldUuid := uuid.New()

	gamesRepositoryMock := new(mocks.GamesRepository)
	gamesRepositoryMock.On("FindById", &gameUuid, false).Return(game, nil)

	gameService := NewGamesService(gamesRepositoryMock, nil)

	fieldStatus := models.FieldStatusShown

	status, err := gameService.ExecuteFieldAction(&gameUuid, &fieldUuid, fieldStatus)

	assert.Nil(t, status)
	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	assert.Equal(t, "game "+gameUuid.String()+" is finished", err.Message)
}

func Test_gamesServiceImpl_ExecuteFieldAction_FindByIdAndGameId_error(t *testing.T) {
	gameUuid := uuid.New()
	game := &models.Game{ID: gameUuid, Status: models.GameStatusInProgress}
	fieldUuid := uuid.New()

	gamesRepositoryMock := new(mocks.GamesRepository)
	gamesRepositoryMock.On("FindById", &gameUuid, false).Return(game, nil)

	fieldRepositoryMock := new(mocks.FieldsRepository)
	errExpected := errors.NewInternalServerApiError(errors.NewError("panic"))
	fieldRepositoryMock.On("FindByIdAndGameId", &fieldUuid, &gameUuid).Return(nil, errExpected)

	gameService := NewGamesService(gamesRepositoryMock, fieldRepositoryMock)

	fieldStatus := models.FieldStatusShown

	status, err := gameService.ExecuteFieldAction(&gameUuid, &fieldUuid, fieldStatus)

	assert.Nil(t, status)
	assert.Equal(t, errExpected, err)
}

func Test_gamesServiceImpl_ExecuteFieldAction_setStatus_error(t *testing.T) {
	gameUuid := uuid.New()
	game := &models.Game{ID: gameUuid, Status: models.GameStatusInProgress}
	fieldUuid := uuid.New()
	field := &models.Field{ID: fieldUuid, Status: models.FieldStatusShown, GameId: gameUuid, Value: &models.MineString}

	gamesRepositoryMock := new(mocks.GamesRepository)
	gamesRepositoryMock.On("FindById", &gameUuid, false).Return(game, nil)

	fieldRepositoryMock := new(mocks.FieldsRepository)
	fieldRepositoryMock.On("FindByIdAndGameId", &fieldUuid, &gameUuid).Return(field, nil)

	gameService := NewGamesService(gamesRepositoryMock, fieldRepositoryMock)

	fieldStatus := models.FieldStatusShown

	status, err := gameService.ExecuteFieldAction(&gameUuid, &fieldUuid, fieldStatus)

	assert.Nil(t, status)
	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
}

func Test_gamesServiceImpl_ExecuteFieldAction_fieldUpdate_error(t *testing.T) {
	gameUuid := uuid.New()
	game := &models.Game{ID: gameUuid, Status: models.GameStatusInProgress}
	fieldUuid := uuid.New()
	field := &models.Field{ID: fieldUuid, Status: models.FieldStatusHidden, GameId: gameUuid, Value: &models.MineString}

	gamesRepositoryMock := new(mocks.GamesRepository)
	gamesRepositoryMock.On("FindById", &gameUuid, false).Return(game, nil)

	fieldRepositoryMock := new(mocks.FieldsRepository)
	fieldRepositoryMock.On("FindByIdAndGameId", &fieldUuid, &gameUuid).Return(field, nil)
	errExpected := errors.NewInternalServerApiError(errors.NewError("panic"))
	fieldRepositoryMock.On("Update", mock.Anything).Return(errExpected)

	gameService := NewGamesService(gamesRepositoryMock, fieldRepositoryMock)

	fieldStatus := models.FieldStatusShown

	status, err := gameService.ExecuteFieldAction(&gameUuid, &fieldUuid, fieldStatus)

	assert.Nil(t, status)
	assert.Equal(t, errExpected, err)
}

func Test_gamesServiceImpl_ExecuteFieldAction_flagged_error(t *testing.T) {
	gameUuid := uuid.New()
	game := &models.Game{ID: gameUuid, Status: models.GameStatusInProgress}
	fieldUuid := uuid.New()
	field := &models.Field{ID: fieldUuid, Status: models.FieldStatusHidden, GameId: gameUuid, Value: &models.MineString}

	gamesRepositoryMock := new(mocks.GamesRepository)
	gamesRepositoryMock.On("FindById", &gameUuid, false).Return(game, nil)

	fieldRepositoryMock := new(mocks.FieldsRepository)
	fieldRepositoryMock.On("FindByIdAndGameId", &fieldUuid, &gameUuid).Return(field, nil)
	fieldRepositoryMock.On("Update", mock.Anything).Return(nil)
	errExpected := errors.NewInternalServerApiError(errors.NewError("panic"))
	fieldRepositoryMock.On("FindMineFieldsFlaggedByGame", mock.Anything).
		Return(nil, errExpected)

	gameService := NewGamesService(gamesRepositoryMock, fieldRepositoryMock)

	fieldStatus := models.FieldStatusFlagged

	status, err := gameService.ExecuteFieldAction(&gameUuid, &fieldUuid, fieldStatus)

	assert.Nil(t, status)
	assert.Equal(t, errExpected, err)
}

func Test_gamesServiceImpl_ExecuteFieldAction_game_update_error(t *testing.T) {
	gameUuid := uuid.New()
	game := &models.Game{ID: gameUuid, Status: models.GameStatusInProgress}
	fieldUuid := uuid.New()
	field := &models.Field{ID: fieldUuid, Status: models.FieldStatusHidden, GameId: gameUuid, Value: &models.MineString}

	gamesRepositoryMock := new(mocks.GamesRepository)
	gamesRepositoryMock.On("FindById", &gameUuid, false).Return(game, nil)
	errExpected := errors.NewInternalServerApiError(errors.NewError("panic"))
	gamesRepositoryMock.On("Update", mock.Anything).Return(errExpected)

	fieldRepositoryMock := new(mocks.FieldsRepository)
	fieldRepositoryMock.On("FindByIdAndGameId", &fieldUuid, &gameUuid).Return(field, nil)
	fieldRepositoryMock.On("Update", mock.Anything).Return(nil)

	gameService := NewGamesService(gamesRepositoryMock, fieldRepositoryMock)

	fieldStatus := models.FieldStatusShown

	status, err := gameService.ExecuteFieldAction(&gameUuid, &fieldUuid, fieldStatus)

	assert.Nil(t, status)
	assert.Equal(t, errExpected, err)
}

func Test_showAdjacentFields(t *testing.T) {
	settingsMock := &models.Settings{Height: 3, Width: 3, MinesQuantity: 1}

	minefield := make([][]models.Field, settingsMock.Height)
	for index := range minefield {
		minefield[index] = make([]models.Field, settingsMock.Width)
	}

	for y := 0; y < settingsMock.Height; y++ {
		for x := 0; x < settingsMock.Width; x++ {
			field := &(minefield)[y][x]
			field.SetPosition(y, x)
			field.SetInitialStatus()
		}
	}

	minefield[1][2].SetInitialHintValue()
	minefield[2][1].SetInitialHintValue()
	minefield[2][2].SetMine()

	initialField := &models.Field{PositionY: 0, PositionX: 0, Status: models.FieldStatusHidden}

	fieldsRepositoryMock := new(mocks.FieldsRepository)
	fieldsRepositoryMock.On("Update", mock.Anything).Return(nil)

	gamesService := gamesServiceImpl{nil, fieldsRepositoryMock}

	showAdjacentFields(gamesService, initialField, &minefield, settingsMock)

	assert.Equal(t, models.FieldStatusShown, minefield[1][0].Status)
	assert.Equal(t, models.FieldStatusShown, minefield[2][0].Status)
	assert.Equal(t, models.FieldStatusShown, minefield[0][1].Status)
	assert.Equal(t, models.FieldStatusShown, minefield[1][0].Status)
	assert.Equal(t, models.FieldStatusShown, minefield[2][0].Status)

	assert.Equal(t, models.FieldStatusShown, minefield[1][1].Status)
	assert.Equal(t, models.FieldStatusShown, minefield[1][2].Status)
	assert.Equal(t, models.FieldStatusShown, minefield[1][2].Status)
	assert.Equal(t, models.FieldStatusHidden, minefield[2][2].Status)
}
