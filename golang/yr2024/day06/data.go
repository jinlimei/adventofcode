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
	coordTempObstacle
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
	case coordTempObstacle:
		return "O"
	default:
		panic("invalid coordObject")
	}
}

type xyCoord struct {
	x, y int
}

func (xy *xyCoord) clone() xyCoord {
	return xyCoord{
		x: xy.x,
		y: xy.y,
	}
}

func (xy *xyCoord) String() string {
	return fmt.Sprintf("(%d,%d)", xy.x, xy.y)
}

type guardMap struct {
	maxCoord      xyCoord
	guard         *theGuard
	tempObstacle  *xyCoord
	obstacles     []xyCoord
	walkPaths     []xyCoord
	tempObstacles []xyCoord
}

func (gm *guardMap) getObjectAtPos(x, y int) coordObject {
	// Handle temporary obstacles.
	if gm.tempObstacle != nil && gm.tempObstacle.x == x && gm.tempObstacle.y == y {
		return coordObstacle
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

	for _, tempObstacle := range gm.tempObstacles {
		if tempObstacle.x == x && tempObstacle.y == y {
			return coordTempObstacle
		}
	}

	if gm.guard.loc.x == x && gm.guard.loc.y == y {
		return coordGuard
	}

	return coordNothing
}

func (gm *guardMap) distinctPositions() int {
	// the +1 is necessary for the *existing* position of the guard.
	return len(gm.walkPaths) + 1
}

func (gm *guardMap) findLoopObstacles() {
	var (
		pos xyCoord

		obstacles = make([]xyCoord, 0)
	)

	gm.guard.reset()

	for {
		pos = gm.guard.coordsInFront()

		// we failed to find a loop
		if !gm.isValidCoords(pos) {
			break
		}

		switch gm.getObjectAtPos(pos.x, pos.y) {
		case coordNothing, coordWalkPath, coordGuard:
			gm.tempObstacle = &xyCoord{pos.x, pos.y}

			if gm.isInLoop(gm.guard.facing, gm.guard.loc) {
				obstacles = append(obstacles, xyCoord{pos.x, pos.y})
			}

			gm.tempObstacle = nil
			gm.guard.loc = pos
		case coordObstacle:
			gm.guard.facing = gm.guard.facing.turnRight()
		default:
			panic("invalid coord object")
		}
	}

	gm.tempObstacles = obstacles
}

func (gm *guardMap) isInLoop(facing guardDirection, pos xyCoord) bool {
	guard := &theGuard{
		start:       pos,
		loc:         pos,
		steps:       0,
		facing:      facing,
		startFacing: facing,
	}

	returnedToStart := false

	// Do we return to our start position!
walkLoop:
	for {
		pos = guard.coordsInFront()

		if !gm.isValidCoords(pos) {
			break
		}

		if pos.x == guard.start.x && pos.y == guard.start.y {
			returnedToStart = true
			break
		}

		switch gm.getObjectAtPos(pos.x, pos.y) {
		case coordNothing, coordWalkPath:
			guard.loc = pos
			guard.steps++
		case coordObstacle:
			guard.facing = guard.facing.turnRight()
		case coordGuard:
			returnedToStart = true
			break walkLoop
		default:
			panic("invalid coord object")
		}
	}

	return returnedToStart
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
			panic("invalid coord object")
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
