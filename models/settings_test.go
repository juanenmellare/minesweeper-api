package models

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_buildSettingsMinError(t *testing.T) {
	apiErr := buildSettingsMinMaxError("field", 1, 99)

	assert.NotNil(t, apiErr)
	assert.Equal(t, "minefield field must be bigger than 1 and less than 99", apiErr.Message)
	assert.Equal(t, http.StatusBadRequest, apiErr.StatusCode)
}

func TestSettings_Validate_width(t *testing.T) {
	apiErr := Settings{Width: 0}.Validate()

	assert.NotNil(t, apiErr)
	assert.Equal(t, "minefield width must be bigger than 3 and less than 30", apiErr.Message)

}

func TestSettings_Validate_height(t *testing.T) {
	apiErr := Settings{Width: 30, Height: 0}.Validate()

	assert.NotNil(t, apiErr)
	assert.Equal(t, "minefield height must be bigger than 3 and less than 16", apiErr.Message)
}

func TestSettings_Validate_mines_quantity(t *testing.T) {
	apiErr := Settings{Width: 30, Height: 16, MinesQuantity: 0}.Validate()

	assert.NotNil(t, apiErr)
	assert.Equal(t, "minefield mines quantity must be bigger than 1 and less than 479", apiErr.Message)
}

func TestSettings_Validate(t *testing.T) {
	apiErr := Settings{Width: 3, Height: 3, MinesQuantity: 1}.Validate()

	assert.Nil(t, apiErr)
}
