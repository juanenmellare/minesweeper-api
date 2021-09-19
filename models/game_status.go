package models

type GameStatus string

const (
	StatusInProgress GameStatus = "IN_PROGRESS"
	StatusWon        GameStatus = "WON"
	StatusLost       GameStatus = "LOST"
)
