package models

import (
	"fmt"
	"github.com/google/uuid"
	"strconv"
)

var mineString = "MINE"

type Field struct {
	ID        uuid.UUID `json:"-" gorm:"type:uuid;default:uuid_generate_v4()"`
	Value     *string   `json:"value"`
	PositionY int       `json:"positionY"`
	PositionX int       `json:"positionX"`
	GameId    uuid.UUID `json:"-"`
	Game      Game      `json:"-" gorm:"foreignKey:GameId;references:id"`
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
	f.setValue(mineString)
}

func (f *Field) IsMine() bool {
	return *f.Value == mineString
}

func (f *Field) IsNil() bool {
	return f.Value == nil
}
