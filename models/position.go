package models

import "strconv"

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p Position) String() string {
	yString := strconv.Itoa(p.Y)
	xString := strconv.Itoa(p.X)

	return "{ y: " + yString + ", x: " + xString + " }"

}
