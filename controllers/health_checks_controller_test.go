package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHealthCheckController(t *testing.T) {
	assert.Implements(t, (*HealthCheckController)(nil), NewHealthCheckController())
}

func Test_healthCheckControllerImpl_Ping(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	healthCheckController := NewHealthCheckController()
	healthCheckController.Ping(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}
