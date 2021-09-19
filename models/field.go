package models

import (
	"fmt"
	"strconv"
)

var mineString = "MINE"

type Field struct {
	Value *string `json:"value"`
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
