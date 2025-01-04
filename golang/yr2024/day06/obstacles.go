package day06

import (
	"errors"
)

func generateListOfObstacles(nm *nMap) []xyCoord {
	// We're going to traverse the map our initial walk path
	// so we'll have a list of walk path coordinates to try an object at.
	initW := initWalk(nm)
	initW.traverse()

	loopCoords := make([]xyCoord, 0)

	for _, xy := range initW.walkPaths {
		//fmt.Printf("TRYING xy %s...", xy.String())

		// We'll skip the single coordinate in the initial map
		// that is the guard, so we don't do anything insane.
		if nm.getEntityAtCoords(xy) == entityGuard {
			//fmt.Printf(" SKIP due to being guard on map\n")
			continue
		}

		// We're creating a temp obstacle
		nm.tempObstacle = xy.clone()

		w := initWalk(nm)
		err := w.traverse()

		if errors.Is(err, ErrInLoop) {
			//fmt.Printf(" FOUND LOOP!\n")
			loopCoords = append(loopCoords, *nm.tempObstacle)
			nm.tempObstacle = nil
		} else if err != nil {
			//fmt.Printf(" ERROR %v\n", err)

			w.visualize()

			panic("failed")
		} else {
			//fmt.Printf(" NOFIND\n")
		}
	}

	return loopCoords
}
