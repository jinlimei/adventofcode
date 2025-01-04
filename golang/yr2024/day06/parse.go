package day06

import (
	"strings"
)

func parseGuardDirection(g rune) direction {
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

func parse(input string) *nMap {
	// Let's simplify the parsing here and avoid having to deal
	// with empty newlines at start since I love doing that for Part1Prompt/Part2Prompt
	input = strings.TrimSpace(input) + "\n"

	var (
		runes = []rune(input)
		rLen  = len(runes)

		pos = 0
		chr rune

		xPos = 0

		nm = initMap()

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
			nm.obstacles = append(nm.obstacles, xyCoord{
				x: xPos,
				y: maxY,
			})

			//fmt.Printf("found obstacle at (%d,%d)\n", xPos, maxY)
			xPos++
		case '^', '>', 'v', '<':
			nm.guardStart = xyCoord{
				x: xPos,
				y: maxY,
			}

			nm.guardFacing = parseGuardDirection(chr)

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

	nm.maxCoord = xyCoord{maxX, maxY}

	return nm
}
