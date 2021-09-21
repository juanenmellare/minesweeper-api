package models

import (
	"fmt"
	"github.com/google/uuid"
	"strconv"
)

var MineString = "MINE"

type Field struct {
	ID        uuid.UUID   `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Value     *string     `json:"value"`
	Status    FieldStatus `json:"status"`
	PositionY int         `json:"positionY" gorm:"index:idx_game_position,priority:3"`
	PositionX int         `json:"positionX" gorm:"index:idx_game_position,priority:2"`
	GameId    uuid.UUID   `json:"-" gorm:"index:idx_game_position,priority:1;column:game_id"`
	Game      Game        `json:"-" gorm:"foreignKey:GameId;references:id"`
}

func (f *Field) setValue(value string) {
	f.Value = &value
}

func (f *Field) SetInitialHintValue() {
	f.setValue("1")
}

func (f *Field) IncrementHintValue() {
	intValue, _ := strconv.Atoi(*f.Value)
	value := fmt.Sprintf("%v", intValue+1)
	f.setValue(value)
}

func (f *Field) SetMine() {
	f.setValue(MineString)
}

func (f *Field) IsMine() bool {
	if f.IsNil() {
		return false
	}
	return *f.Value == MineString
}

func (f *Field) IsNil() bool {
	return f.Value == nil
}

func (f *Field) SetPosition(y, x int) {
	f.PositionY = y
	f.PositionX = x
}

func (f *Field) SetInitialStatus() {
	f.Status = FieldStatusHidden
}

func (f *Field) Show() {
	f.Status = FieldStatusShown
}

func (f *Field) SetStatus(candidateStatus FieldStatus) error {
	if err := ValidateFieldStatusTransition(f.Status, candidateStatus); err != nil {
		return err
	}
	f.Status = candidateStatus

	return nil
}
