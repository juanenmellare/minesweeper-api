package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPosition_String(t *testing.T) {
	position := Position{X: 3, Y: 2}

	assert.Equal(t, "{ y: 2, x: 3 }", position.String())
}
