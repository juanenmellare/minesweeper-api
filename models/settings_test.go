package models

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_buildSettingsMinError(t *testing.T) {
	apiErr := buildSettingsMinError("field", 99)

	assert.NotNil(t, apiErr)
	assert.Equal(t, "minefield field must not be less than 99", apiErr.Message)
	assert.Equal(t, http.StatusBadRequest, apiErr.StatusCode)
}

func TestSettings_Validate_width(t *testing.T) {
	apiErr := Settings{Width: 0}.Validate()

	assert.NotNil(t, apiErr)
	assert.Equal(t, "minefield width must not be less than 3", apiErr.Message)

}

func TestSettings_Validate_height(t *testing.T) {
	apiErr := Settings{Width: 99, Height: 0}.Validate()

	assert.NotNil(t, apiErr)
	assert.Equal(t, "minefield height must not be less than 3", apiErr.Message)
}

func TestSettings_Validate_bombs_quantity(t *testing.T) {
	apiErr := Settings{Width: 99, Height: 99, BombsQuantity: 0}.Validate()

	assert.NotNil(t, apiErr)
	assert.Equal(t, "minefield bombs quantity must not be less than 1", apiErr.Message)
}
