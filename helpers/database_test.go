package helpers

import (
	"github.com/stretchr/testify/assert"
	"minesweeper-api/errors"
	"net/http"
	"testing"
)

func TestValidateDatabaseTx_not_found(t *testing.T) {
	err := errors.NewError("record not found")

	apiErr := ValidateDatabaseTxError(err, "foo")

	assert.Equal(t, "foo not found", apiErr.Message)
	assert.Equal(t, http.StatusNotFound, apiErr.StatusCode)
}

func TestValidateDatabaseTx_internal_server_error(t *testing.T) {
	err := errors.NewError("panic")

	apiErr := ValidateDatabaseTxError(err, "foo")

	assert.Equal(t, "panic", apiErr.Message)
	assert.Equal(t, http.StatusInternalServerError, apiErr.StatusCode)
}

func TestValidateDatabaseTx(t *testing.T) {
	apiErr := ValidateDatabaseTxError(nil, "foo")

	assert.Nil(t, apiErr)
}
