package day06

import (
	"strings"
)

func parseGuardDirection(g rune) guardDirection {
	switch g {
	case '^':
		return facingNorth
	case '>':
		return facingEast
	case 'v':
		return facingSouth
	case '<':
		return facingWest
	default:
		panic("undefined guard direction")
	}
}

func parse(input string) *guardMap {
	// Let's simplify the parsing here and avoid having to deal
	// with empty newlines at start since I love doing that for Part1Prompt/Part2Prompt
	input = strings.TrimSpace(input) + "\n"

	var (
		runes = []rune(input)
		rLen  = len(runes)

		pos = 0
		chr rune

		xPos = 0

		gMap = &guardMap{
			guard:     &theGuard{},
			obstacles: make([]xyCoord, 0),
			walkPaths: make([]xyCoord, 0),
		}

		maxX = 0
		maxY = 0
	)

	for ; pos < rLen; pos++ {
		chr = runes[pos]

		//fmt.Printf("(%d,%d) %s\n", xPos, maxY, string(chr))

		switch chr {
		case '.':
			xPos++
		case '#':
			gMap.obstacles = append(gMap.obstacles, xyCoord{
				x: xPos,
				y: maxY,
			})

			//fmt.Printf("found obstacle at (%d,%d)\n", xPos, maxY)
			xPos++
		case '^', '>', 'v', '<':
			gMap.guard.start = xyCoord{
				x: xPos,
				y: maxY,
			}

			gMap.guard.loc = gMap.guard.start
			gMap.guard.startFacing = parseGuardDirection(chr)
			gMap.guard.facing = gMap.guard.startFacing
			xPos++
		case '\n':
			maxY++

			if xPos > maxX {
				maxX = xPos
			}

			xPos = 0
		default:
			xPos++
		}

	}

	gMap.maxCoord = xyCoord{maxX, maxY}

	return gMap
}
