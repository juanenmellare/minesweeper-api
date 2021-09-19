package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestField_IncrementHintValue(t *testing.T) {
	value := "1"
	field := Field{Value: &value}

	field.IncrementHintValue()

	assert.Equal(t, "2", *field.Value)
}

func TestField_IsNil(t *testing.T) {
	field := Field{}

	assert.True(t, field.IsNil())
}

func TestField_IsNil_false(t *testing.T) {
	value := "foo"
	field := Field{Value: &value}

	assert.False(t, field.IsNil())
}

func TestField_IsMine(t *testing.T) {
	field := Field{Value: &mineString}

	assert.True(t, field.IsMine())
}

func TestField_IsMine_false(t *testing.T) {
	value := "foo"
	field := Field{Value: &value}

	assert.False(t, field.IsMine())
}

func TestField_SetInitialHintValue(t *testing.T) {
	field := Field{}

	field.SetInitialHintValue()

	assert.Equal(t, "1", *field.Value)
}

func TestField_SetMine(t *testing.T) {
	field := Field{}

	field.SetMine()

	assert.Equal(t, mineString, *field.Value)
}

func TestField_setValue(t *testing.T) {
	field := Field{}
	value := "foo"

	field.setValue(value)

	assert.Equal(t, value, *field.Value)
}

func TestField_SetPosition(t *testing.T) {
	field := Field{}
	y := 0
	x := 1

	field.SetPosition(y, x)

	assert.Equal(t, y, field.PositionY)
	assert.Equal(t, x, field.PositionX)
}
