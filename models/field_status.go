package models

import (
	"minesweeper-api/errors"
)

type FieldStatus string

const (
	FieldStatusHidden     FieldStatus = "HIDDEN"
	FieldStatusShown      FieldStatus = "SHOWN"
	FieldStatusFlagged    FieldStatus = "FLAGGED"
	FieldStatusQuestioned FieldStatus = "QUESTIONED"
)

func ValidateFieldStatusTransition(current, candidate FieldStatus) error {
	if current == candidate {
		return errors.NewError("unnecessary status transition")
	}

	switch current {
	case FieldStatusShown:
		return errors.NewError("once the field is shown there is no going back")
	case FieldStatusFlagged, FieldStatusQuestioned:
		if candidate != FieldStatusHidden {
			message := "the only available transition is from " + string(current) + " is to " + string(FieldStatusShown)
			return errors.NewError(message)
		}
	}
	return nil
}
