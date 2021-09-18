package errors

import "net/http"

type ApiError struct {
	Message    string
	Status     string
	StatusCode int
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
