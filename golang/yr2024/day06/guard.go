package day06

import (
	"fmt"
)

type theGuard struct {
	start       xyCoord
	loc         xyCoord
	steps       int
	facing      direction
	startFacing direction
}

func newGuard(start xyCoord, facing direction) *theGuard {
	return &theGuard{
		start:       start,
		loc:         start,
		steps:       0,
		facing:      facing,
		startFacing: facing,
	}
}

func (tg *theGuard) coordsInFront() xyCoord {
	switch tg.facing {
	case facingNorth:
		return xyCoord{tg.loc.x, tg.loc.y - 1}
	case facingEast:
		return xyCoord{tg.loc.x + 1, tg.loc.y}
	case facingSouth:
		return xyCoord{tg.loc.x, tg.loc.y + 1}
	case facingWest:
		return xyCoord{tg.loc.x - 1, tg.loc.y}
	default:
		panic("guard is dead or something idk")
	}
}

func (tg *theGuard) turnRight() {
	tg.facing = tg.facing.turnRight()
}

func (tg *theGuard) String() string {
	return fmt.Sprintf(
		"start=%s facing %s,current=%s facing %s",
		tg.start.String(),
		tg.startFacing,
		tg.loc.String(),
		tg.facing,
	)
}

func (tg *theGuard) reset() {
	tg.loc = tg.start
	tg.facing = tg.startFacing
	tg.steps = 0
}
