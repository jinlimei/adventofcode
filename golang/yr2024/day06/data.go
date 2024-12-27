package day06

import (
	"fmt"
)

type coordObject uint8

const (
	coordNothing coordObject = iota
	coordObstacle
	coordWalkPath
	coordGuard
)

func (co coordObject) String() string {
	switch co {
	case coordNothing:
		return "."
	case coordObstacle:
		return "#"
	case coordWalkPath:
		return "X"
	case coordGuard:
		return "g"
	default:
		panic("invalid coordObject")
	}
}

type xyCoord struct {
	x, y int
}

func (xy *xyCoord) String() string {
	return fmt.Sprintf("(%d,%d)", xy.x, xy.y)
}

type guardMap struct {
	maxCoord  xyCoord
	guard     *theGuard
	obstacles []xyCoord
	walkPaths []xyCoord
}

func (gm *guardMap) getObjectAtPos(x, y int) coordObject {
	if gm.guard.loc.x == x && gm.guard.loc.y == y {
		return coordGuard
	}

	for _, obstacle := range gm.obstacles {
		if obstacle.x == x && obstacle.y == y {
			return coordObstacle
		}
	}

	for _, walkPath := range gm.walkPaths {
		if walkPath.x == x && walkPath.y == y {
			return coordWalkPath
		}
	}

	return coordNothing
}

func (gm *guardMap) distinctPositions() int {
	// the +1 is necessary for the *existing* position of the guard.
	return len(gm.walkPaths) + 1
}

func (gm *guardMap) traverse() {
	var (
		pos xyCoord
	)

	// we're going to reset the guard position
	gm.guard.reset()

	for {
		pos = gm.guard.coordsInFront()

		// We've completed our routine, as we've walked off the map
		// and presumably off a cliff and/or into the ocean.
		if !gm.isValidCoords(pos) {
			fmt.Printf("OH NO THE GUARD FELL OFF A CLIFF AND/OR INTO THE OCEAN: %s\n", pos.String())
			fmt.Printf("WHAT IS %s\n", gm.getObjectAtPos(pos.x, 0))

			fmt.Printf("ROW %d\n", pos.x)
			for k := 0; k < gm.maxCoord.y; k++ {
				fmt.Printf("%s", gm.getObjectAtPos(pos.x, k))
			}

			fmt.Println()

			// we'll still place the guard off into nowhere
			gm.guard.loc = pos
			break
		}

		switch gm.getObjectAtPos(pos.x, pos.y) {
		case coordNothing:
			// we're fine to proceed!
			// let's record our walk position
			gm.walkPaths = append(gm.walkPaths, pos)
			gm.guard.loc = pos
			gm.guard.steps++
		case coordObstacle:
			// we are not fine to proceed we need to change direction
			gm.guard.facing = gm.guard.facing.turnRight()
		case coordWalkPath:
			// it's fine that we're walking
			// but we won't re-record the walk path
			gm.guard.loc = pos
			gm.guard.steps++
		case coordGuard:
			panic("what the actual fuck we've talked into ourselves, burn everything down")
		default:
			panic("unhandled default case")
		}
	}
}

func (gm *guardMap) isValidCoords(xy xyCoord) bool {
	if xy.x < 0 || xy.y < 0 {
		return false
	}

	if xy.x >= gm.maxCoord.x || xy.y >= gm.maxCoord.y {
		return false
	}

	return true
}

func (gm *guardMap) visualize() {
	for y := 0; y < gm.maxCoord.y; y++ {
		for x := 0; x < gm.maxCoord.x; x++ {
			pos := gm.getObjectAtPos(x, y)

			if pos == coordGuard {
				fmt.Print(gm.guard.facing)
			} else {
				fmt.Print(pos)
			}
		}

		fmt.Println()
	}
}
