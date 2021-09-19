package models

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGame(t *testing.T) {
	game := &Game{}

	assert.Equal(t, uuid.UUID{}, game.ID)
	assert.Equal(t, time.Time{}, game.StartedAt)
	assert.Equal(t, GameStatus(""), game.Status)
	assert.Equal(t, (*[]Field)(nil), game.Minefield)
	assert.Equal(t, Settings{}, game.Settings)
}
