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

var FieldStatusStrategyMap = map[string]FieldStatus{
	"show":     FieldStatusShown,
	"flag":     FieldStatusFlagged,
	"question": FieldStatusQuestioned,
	"hide":     FieldStatusHidden,
}

func ValidateFieldStatusTransition(current, candidate FieldStatus) error {
	if current == candidate {
		return errors.NewError("already " + string(current))
	}

	switch current {
	case FieldStatusShown:
		return errors.NewError("once the field is shown there is no going back")
	case FieldStatusFlagged, FieldStatusQuestioned:
		if candidate != FieldStatusHidden {
			message := "the only available transition for this status is from " + string(current) +
				" is to " + string(FieldStatusShown)
			return errors.NewError(message)
		}
	}
	return nil
}
