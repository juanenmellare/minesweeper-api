package models

import (
	"github.com/google/uuid"
	"time"
)

type Game struct {
	ID         uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	StartedAt  time.Time  `json:"startedAt"`
	Settings   Settings   `json:"settings" gorm:"foreignKey:id"`
	SettingsID uuid.UUID  `json:"-" gorm:"foreignKey:ID;references:SettingsID"`
	Minefield  *[]Field   `json:"minefield" gorm:"GameId"`
	Status     GameStatus `json:"status"`
}
