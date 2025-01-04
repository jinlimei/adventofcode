package day06

import (
	"errors"
	"fmt"
)

var (
	ErrInLoop   = errors.New("traversed in a loop")
	ErrInfGuard = errors.New("traversing hit the infinite loop guard")
)

type walk struct {
	walkPaths []xyCoord
	nMap      *nMap
}

func initWalk(nm *nMap) *walk {
	return &walk{
		walkPaths: make([]xyCoord, 0),
		nMap:      nm,
	}
}

func (w *walk) distinctPositions() int {
	return len(w.walkPaths)
}

func (w *walk) traverse() error {
	var (
		guard  = newGuard(w.nMap.guardStart, w.nMap.guardFacing)
		nxtPos xyCoord

		maxWalk  = w.nMap.maxCoord.x * w.nMap.maxCoord.y
		infGuard = 0
	)

	for {
		nxtPos = guard.coordsInFront()

		if !w.nMap.isValidCoords(nxtPos) {
			w.walkPaths = append(w.walkPaths, guard.loc)

			guard.loc = nxtPos
			break
		}

		if w.nMap.guardStart.Equals(nxtPos) && w.nMap.guardFacing == guard.facing {
			return ErrInLoop
		}

		switch w.getEntityOrWalkPathAtCoords(nxtPos) {
		case entityNothing, entityGuard:
			w.walkPaths = append(w.walkPaths, nxtPos)
			guard.loc = nxtPos
			guard.steps++
		case entityWalkPath:
			// We won't re-record walkPath since we've already walked here
			guard.loc = nxtPos
			guard.steps++
		case entityObstacle, entityTempObstacle:
			guard.turnRight()
		default:
			panic("invalid entity provided")
		}

		infGuard++

		if infGuard > maxWalk {
			return ErrInLoop
		}
	}

	return nil
}

func (w *walk) getEntityOrWalkPathAtCoords(xy xyCoord) entity {
	return w.getEntityOrWalkPathAtXY(xy.x, xy.y)
}

func (w *walk) getEntityOrWalkPathAtXY(x, y int) entity {
	var (
		wLen = len(w.walkPaths)
		pos  = 0
	)

	for pos = 0; pos < wLen; pos++ {
		if w.walkPaths[pos].Is(x, y) {
			return entityWalkPath
		}
	}

	return w.nMap.getEntityAtXY(x, y)
}

func (w *walk) visualize() {
	maxX, maxY := w.nMap.maxCoord.x, w.nMap.maxCoord.y

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			ent := w.getEntityOrWalkPathAtXY(x, y)

			fmt.Print(ent)
		}

		fmt.Println()
	}
}
