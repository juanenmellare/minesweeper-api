package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameStatus(t *testing.T) {
	assert.Equal(t, "IN_PROGRESS", string(GameStatusInProgress))
	assert.Equal(t, "WON", string(GameStatusWon))
	assert.Equal(t, "LOST", string(GameStatusLost))
}
