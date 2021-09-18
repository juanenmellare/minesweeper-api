package models

type Status string

const (
	StatusInProgress Status = "IN_PROGRESS"
	StatusWon        Status = "WON"
	StatusLost       Status = "LOST"
)
