package day06

import (
	"fmt"
)

type nMap struct {
	maxCoord     xyCoord
	guardStart   xyCoord
	guardFacing  direction
	obstacles    []xyCoord
	foundTemps   []xyCoord
	tempObstacle *xyCoord
}

func initMap() *nMap {
	return &nMap{
		obstacles:  make([]xyCoord, 0),
		foundTemps: make([]xyCoord, 0),
	}
}

func (nm *nMap) isValidXY(x, y int) bool {
	if x < 0 || x > nm.maxCoord.x {
		return false
	}

	if y < 0 || y > nm.maxCoord.y {
		return false
	}

	return true
}

func (nm *nMap) isValidCoords(xy xyCoord) bool {
	return nm.isValidXY(xy.x, xy.y)
}

func (nm *nMap) getEntityAtXY(x, y int) entity {

	if nm.guardStart.Is(x, y) {
		return entityGuard
	}

	if nm.tempObstacle != nil && nm.tempObstacle.Is(x, y) {
		return entityObstacle
	}

	var (
		oLen = len(nm.obstacles)
		tLen = len(nm.foundTemps)
		pos  = 0
	)

	for pos = 0; pos < oLen; pos++ {
		if nm.obstacles[pos].Is(x, y) {
			return entityObstacle
		}
	}

	for pos = 0; pos < tLen; pos++ {
		if nm.foundTemps[pos].Is(x, y) {
			return entityTempObstacle
		}
	}

	return entityNothing
}

func (nm *nMap) getEntityAtCoords(xy xyCoord) entity {
	return nm.getEntityAtXY(xy.x, xy.y)
}

func (nm *nMap) visualize() {
	for y := 0; y < nm.maxCoord.y; y++ {
		for x := 0; x < nm.maxCoord.x; x++ {
			fmt.Print(nm.getEntityAtXY(x, y))
		}

		fmt.Println()
	}
}
