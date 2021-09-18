package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHealthChecksController(t *testing.T) {
	assert.Implements(t, (*HealthChecksController)(nil), NewHealthChecksController())
}

func Test_healthChecksControllerImpl_Ping(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	healthChecksController := NewHealthChecksController()
	healthChecksController.Ping(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}
