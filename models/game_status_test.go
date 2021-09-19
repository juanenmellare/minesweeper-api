package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatus(t *testing.T) {
	assert.Equal(t, "IN_PROGRESS", string(StatusInProgress))
	assert.Equal(t, "WON", string(StatusWon))
	assert.Equal(t, "LOST", string(StatusLost))
}
