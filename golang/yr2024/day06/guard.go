package day06

import (
	"fmt"
)

type theGuard struct {
	start       xyCoord
	loc         xyCoord
	uniqueSteps int
	facing      guardDirection
	startFacing guardDirection
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
	tg.uniqueSteps = 0
}

type guardDirection uint8

const (
	facingNorth guardDirection = iota
	facingEast
	facingSouth
	facingWest
)

func (gd guardDirection) String() string {
	switch gd {
	case facingNorth:
		return "^"
	case facingEast:
		return ">"
	case facingSouth:
		return "v"
	case facingWest:
		return "<"
	default:
		panic("invalid guardDirection")
	}
}

func (gd guardDirection) turnRight() guardDirection {
	switch gd {
	case facingNorth:
		return facingEast
	case facingEast:
		return facingSouth
	case facingSouth:
		return facingWest
	case facingWest:
		return facingNorth
	default:
		panic("invalid guardDirection")
	}
}
