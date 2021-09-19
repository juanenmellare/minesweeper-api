package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
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

	alertsController := NewGamesController(gamesServiceMock)

	alertsController.Create(c)

	expectedJsonString := "{\"startedAt\":\"0001-01-01T00:00:00Z\",\"settings\":{\"width\":0,\"height\":0," +
		"\"bombsQuantity\":0},\"minefield\":null,\"status\":\"\"}"

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, expectedJsonString, w.Body.String())
}

func Test_gamesControllerImpl_Create_bind_err(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = &http.Request{Body: nil}

	gamesServiceMock := new(mocks.GamesService)

	alertsController := NewGamesController(gamesServiceMock)

	alertsController.Create(c)

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

	alertsController := NewGamesController(gamesServiceMock)

	alertsController.Create(c)

	expectedJsonString := "{\"message\":\"bad_request\",\"status\":\"bad request\",\"status_code\":400}"

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expectedJsonString, w.Body.String())
}
