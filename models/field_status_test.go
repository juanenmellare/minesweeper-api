package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFieldStatus(t *testing.T) {
	assert.Equal(t, "HIDDEN", string(FieldStatusHidden))
	assert.Equal(t, "SHOWN", string(FieldStatusShown))
	assert.Equal(t, "FLAGGED", string(FieldStatusFlagged))
	assert.Equal(t, "QUESTIONED", string(FieldStatusQuestioned))
}

func TestValidateFieldStatusTransition(t *testing.T) {
	err := ValidateFieldStatusTransition(FieldStatusShown, FieldStatusHidden)
	assert.NotNil(t, err)
	assert.Equal(t, "once the field is shown there is no going back", err.Error())
	assert.NotNil(t, ValidateFieldStatusTransition(FieldStatusShown, FieldStatusFlagged))
	err = ValidateFieldStatusTransition(FieldStatusShown, FieldStatusShown)
	assert.NotNil(t, err)
	assert.Equal(t, "already SHOWN", err.Error())
	assert.NotNil(t, ValidateFieldStatusTransition(FieldStatusShown, FieldStatusQuestioned))

	assert.Nil(t, ValidateFieldStatusTransition(FieldStatusHidden, FieldStatusShown))
	assert.Nil(t, ValidateFieldStatusTransition(FieldStatusHidden, FieldStatusQuestioned))
	assert.Nil(t, ValidateFieldStatusTransition(FieldStatusHidden, FieldStatusFlagged))
	assert.NotNil(t, ValidateFieldStatusTransition(FieldStatusHidden, FieldStatusHidden))

	err = ValidateFieldStatusTransition(FieldStatusQuestioned, FieldStatusShown)
	assert.NotNil(t, err)
	assert.Equal(t, "the only available transition for this status is from QUESTIONED is to SHOWN", err.Error())
	assert.NotNil(t, ValidateFieldStatusTransition(FieldStatusQuestioned, FieldStatusQuestioned))
	assert.NotNil(t, ValidateFieldStatusTransition(FieldStatusQuestioned, FieldStatusFlagged))
	assert.Nil(t, ValidateFieldStatusTransition(FieldStatusQuestioned, FieldStatusHidden))

	assert.NotNil(t, ValidateFieldStatusTransition(FieldStatusFlagged, FieldStatusShown))
	assert.NotNil(t, ValidateFieldStatusTransition(FieldStatusFlagged, FieldStatusFlagged))
	assert.NotNil(t, ValidateFieldStatusTransition(FieldStatusFlagged, FieldStatusQuestioned))
	assert.Nil(t, ValidateFieldStatusTransition(FieldStatusFlagged, FieldStatusHidden))
}
