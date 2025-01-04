package day06

import (
	"fmt"
)

type xyCoord struct {
	x, y int
}

func (xy *xyCoord) clone() *xyCoord {
	return &xyCoord{
		x: xy.x,
		y: xy.y,
	}
}

func (xy *xyCoord) String() string {
	return fmt.Sprintf("(%d,%d)", xy.x, xy.y)
}

func (xy *xyCoord) Equals(xy2 xyCoord) bool {
	return xy.x == xy2.x && xy.y == xy2.y
}

func (xy *xyCoord) Is(x, y int) bool {
	return xy.x == x && xy.y == y
}
