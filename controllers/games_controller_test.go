package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"minesweeper-api/errors"
	"minesweeper-api/models"
	"minesweeper-api/services/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_gamesControllerImpl_Create(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	settings := models.Settings{}
	requestByte, _ := json.Marshal(settings)
	requestReader := bytes.NewReader(requestByte)
	requestBody := ioutil.NopCloser(requestReader)

	c.Request = &http.Request{Body: requestBody}

	gamesServiceMock := new(mocks.GamesService)
	gamesServiceMock.On("Create", &settings).Return(&models.Game{}, nil)

	gamesController := NewGamesController(gamesServiceMock)

	gamesController.Create(c)

	expectedJsonString := "{\"id\":\"00000000-0000-0000-0000-000000000000\",\"startedAt\":\"0001-01-01T00:00:00Z\"," +
		"\"endedAt\":null,\"duration\":0,\"settings\":{\"width\":0,\"height\":0,\"minesQuantity\":0}," +
		"\"minefield\":null,\"status\":\"\"}"

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, expectedJsonString, w.Body.String())
}

func Test_gamesControllerImpl_Create_bind_err(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = &http.Request{Body: nil}

	gamesServiceMock := new(mocks.GamesService)

	gamesController := NewGamesController(gamesServiceMock)

	gamesController.Create(c)

	expectedJsonString := "{\"message\":\"invalid request\",\"status\":\"bad request\",\"statusCode\":400}"

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expectedJsonString, expectedJsonString)
}

func Test_gamesControllerImpl_Create_gameService_create_err(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	settings := models.Settings{}
	requestByte, _ := json.Marshal(settings)
	requestReader := bytes.NewReader(requestByte)
	requestBody := ioutil.NopCloser(requestReader)

	c.Request = &http.Request{Body: requestBody}

	gamesServiceMock := new(mocks.GamesService)
	err := errors.NewBadRequestApiError(errors.NewError("bad_request"))
	gamesServiceMock.On("Create", &settings).Return(nil, err)

	gamesController := NewGamesController(gamesServiceMock)

	gamesController.Create(c)

	expectedJsonString := "{\"message\":\"bad_request\",\"status\":\"bad request\",\"status_code\":400}"

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expectedJsonString, w.Body.String())
}

func Test_gamesControllerImpl_FindById(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	uuidParam := uuid.New()
	c.Params = append(c.Params, gin.Param{Key: "game-uuid", Value: uuidParam.String()})

	gamesServiceMock := new(mocks.GamesService)
	gamesServiceMock.On("FindById", &uuidParam, true).Return(&models.Game{}, nil)

	gamesController := NewGamesController(gamesServiceMock)

	gamesController.FindById(c)

	expectedJsonString := "{\"id\":\"00000000-0000-0000-0000-000000000000\",\"startedAt\":\"0001-01-01T00:00:00Z\"," +
		"\"endedAt\":null,\"duration\":0,\"settings\":{\"width\":0,\"height\":0,\"minesQuantity\":0}," +
		"\"minefield\":null,\"status\":\"\"}"

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedJsonString, w.Body.String())
}

func Test_gamesControllerImpl_FindById_err(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	uuidParam := uuid.New()
	c.Params = append(c.Params, gin.Param{Key: "game-uuid", Value: uuidParam.String()})

	gamesServiceMock := new(mocks.GamesService)
	err := errors.NewNotFoundError(errors.NewError("not_found"))
	gamesServiceMock.On("FindById", &uuidParam, true).Return(nil, err)

	gamesController := NewGamesController(gamesServiceMock)

	gamesController.FindById(c)

	expectedJsonString := "{\"message\":\"not_found\",\"status\":\"not found\",\"status_code\":404}"

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, expectedJsonString, w.Body.String())
}

func Test_gamesControllerImpl_ExecuteFieldAction(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	gameUuid := uuid.New()
	c.Params = append(c.Params, gin.Param{Key: "game-uuid", Value: gameUuid.String()})
	fieldUuid := uuid.New()
	c.Params = append(c.Params, gin.Param{Key: "field-uuid", Value: fieldUuid.String()})
	action := "show"
	c.Params = append(c.Params, gin.Param{Key: "action", Value: action})

	gamesServiceMock := new(mocks.GamesService)
	gameStatus := models.GameStatusLost
	gamesServiceMock.On("ExecuteFieldAction", &gameUuid, &fieldUuid, models.FieldStatusShown).
		Return(&gameStatus, nil)
	gamesController := NewGamesController(gamesServiceMock)

	gamesController.ExecuteFieldAction(c)

	expectedJsonString := "{\"game_status\":\"LOST\",\"message\":\"field shown\"}"

	assert.Equal(t, http.StatusAccepted, w.Code)
	assert.Equal(t, expectedJsonString, w.Body.String())
}

func Test_gamesControllerImpl_ExecuteFieldAction_action_error(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	action := "showw"
	c.Params = append(c.Params, gin.Param{Key: "action", Value: action})

	gamesServiceMock := new(mocks.GamesService)

	gamesController := NewGamesController(gamesServiceMock)

	gamesController.ExecuteFieldAction(c)

	expectedJsonString := "{\"message\":\"field action /showw route not found\",\"status\":\"bad request\"," +
		"\"status_code\":400}"

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expectedJsonString, w.Body.String())
}

func Test_gamesControllerImpl_ExecuteFieldAction_error(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	gameUuid := uuid.New()
	c.Params = append(c.Params, gin.Param{Key: "game-uuid", Value: gameUuid.String()})
	fieldUuid := uuid.New()
	c.Params = append(c.Params, gin.Param{Key: "field-uuid", Value: fieldUuid.String()})
	action := "show"
	c.Params = append(c.Params, gin.Param{Key: "action", Value: action})

	gamesServiceMock := new(mocks.GamesService)
	err := errors.NewNotFoundError(errors.NewError("not_found"))
	gamesServiceMock.On("ExecuteFieldAction", &gameUuid, &fieldUuid, models.FieldStatusShown).
		Return(nil, err)
	gamesController := NewGamesController(gamesServiceMock)

	gamesController.ExecuteFieldAction(c)

	expectedJsonString := "{\"message\":\"not_found\",\"status\":\"not found\",\"status_code\":404}"

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, expectedJsonString, w.Body.String())
}
