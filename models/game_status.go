package models

type GameStatus string

const (
	GameStatusInProgress GameStatus = "IN_PROGRESS"
	GameStatusWon        GameStatus = "WON"
	GameStatusLost       GameStatus = "LOST"
)
