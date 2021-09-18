package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"minesweeper-api/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_gamesControllerImpl_Create(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	header := map[string][]string{}

	bodyRequest := models.Settings{}
	requestByte, _ := json.Marshal(bodyRequest)
	requestReader := bytes.NewReader(requestByte)
	request := ioutil.NopCloser(requestReader)

	c.Request = &http.Request{Body: request, Header: header}

	alertsController := NewGamesController()
	alertsController.Create(c)

	expectedJsonString := "{\"StartedAt\":\"0001-01-01T00:00:00Z\",\"Settings\":{\"Width\":0,\"Height\":0," +
		"\"BombsQuantity\":0},\"Minefield\":null,\"Status\":\"\"}"

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, expectedJsonString, w.Body.String())
}

func Test_gamesControllerImpl_Create_err(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	header := map[string][]string{}
	c.Request = &http.Request{Body: nil, Header: header}

	alertsController := NewGamesController()
	alertsController.Create(c)

	expectedJsonString := "{\"Message\":\"invalid request\",\"Status\":\"bad request\",\"StatusCode\":400}"

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expectedJsonString, expectedJsonString)
}
