package errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewApiError(t *testing.T) {
	status := "internal server error"
	statusCode := http.StatusInternalServerError
	err := errors.New("panic")

	apiError := newApiError(status, statusCode, err)

	assert.Equal(t, status, apiError.Status)
	assert.Equal(t, statusCode, apiError.StatusCode)
	assert.Equal(t, err.Error(), apiError.Message)
}

func TestNewInternalServerApiError(t *testing.T) {
	err := errors.New("panic")

	apiError := NewInternalServerApiError(err)

	assert.Equal(t, "internal server error", apiError.Status)
	assert.Equal(t, http.StatusInternalServerError, apiError.StatusCode)
	assert.Equal(t, err.Error(), apiError.Message)
}

func TestNewBadRequestApiError(t *testing.T) {
	err := errors.New("panic")

	apiError := NewBadRequestApiError(err)

	assert.Equal(t, "bad request", apiError.Status)
	assert.Equal(t, http.StatusBadRequest, apiError.StatusCode)
	assert.Equal(t, err.Error(), apiError.Message)
}

func TestNewError(t *testing.T) {
	message := "panic"
	err := NewError(message)

	assert.NotNil(t, err)
	assert.Equal(t, message, err.Error())
}

func TestNewNotFoundError(t *testing.T) {
	err := errors.New("panic")

	apiError := NewNotFoundError(err)

	assert.Equal(t, "not found", apiError.Status)
	assert.Equal(t, http.StatusNotFound, apiError.StatusCode)
	assert.Equal(t, err.Error(), apiError.Message)
}
