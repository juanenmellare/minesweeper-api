package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildSettingsMinError(t *testing.T) {
	err := buildSettingsMinError("field", 99)

	assert.NotNil(t, err)
	assert.Equal(t, "minefield field must not be less than 99", err.Error())
}

func TestSettings_Validate_width(t *testing.T) {
	err := Settings{Width: 0}.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "minefield width must not be less than 3", err.Error())

}

func TestSettings_Validate_height(t *testing.T) {
	err := Settings{Width: 99, Height: 0}.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "minefield height must not be less than 3", err.Error())
}

func TestSettings_Validate_bombs_quantity(t *testing.T) {
	err := Settings{Width: 99, Height: 99, BombsQuantity: 0}.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, "minefield bombs quantity must not be less than 1", err.Error())
}
