package errors

import (
	"errors"
	"net/http"
)

type ApiError struct {
	Message    string `json:"message"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
}

func NewError(message string) error {
	return errors.New(message)
}

func newApiError(status string, statusCode int, err error) *ApiError {
	return &ApiError{
		Status:     status,
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}

func NewInternalServerApiError(err error) *ApiError {
	return newApiError("internal server error", http.StatusInternalServerError, err)
}

func NewBadRequestApiError(err error) *ApiError {
	return newApiError("bad request", http.StatusBadRequest, err)
}
