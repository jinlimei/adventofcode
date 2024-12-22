package day04

import "fmt"

func parse(input string) [][]uint8 {
	var (
		rows = make([][]uint8, 0)
		col  = make([]uint8, 0)

		runes = []rune(input)
		rLen  = len(runes)

		pos = 0
		chr rune
	)

	for ; pos < rLen; pos++ {
		chr = runes[pos]

		switch chr {
		case 'X', 'M', 'A', 'S':
			col = append(col, uint8(chr))
		default:
			col = append(col, '?')
		case '\n':
			if len(col) > 0 {
				rows = append(rows, col)
				col = make([]uint8, 0)
			}
		}
	}

	return rows
}

func inXYCoordSlice(x, y int, sli []xycoord) bool {
	for k := 0; k < len(sli); k++ {
		if sli[k].x == x && sli[k].y == y {
			return true
		}
	}

	return false
}

func coordsUsedToRawCoords(coordsUsed [][4]xycoord) []xycoord {
	out := make([]xycoord, 0)

	for _, row := range coordsUsed {
		for k := 0; k < 4; k++ {
			if !inXYCoordSlice(row[k].x, row[k].y, out) {
				out = append(out, row[k])
			}
		}
	}

	return out
}

type xycoord struct {
	x int
	y int
}

func (xy xycoord) String() string {
	return fmt.Sprintf("(%d,%d)", xy.x, xy.y)
}

type peekResult struct {
	values [4]uint8
	coords [4]xycoord
}

func scanGridForXmas(grid [][]uint8) (int, [][4]xycoord) {
	var (
		gridLen    = len(grid)
		coordsUsed = make([][4]xycoord, 0)
		xmasCount  = 0
	)

	for x := 0; x < gridLen; x++ {
		var (
			row  = grid[x]
			cLen = len(row)
		)

		for y := 0; y < cLen; y++ {
			var (
				peekedRight = peekToRight(x, y, row)
				peekedLeft  = peekToLeft(x, y, row)
				peekedUp    = peekUp(x, y, grid)
				peekedDown  = peekDown(x, y, grid)
				peekedTopL  = peekToTopLeft(x, y, grid)
				peekedTopR  = peekToTopRight(x, y, grid)
				peekedBotR  = peekToBottomRight(x, y, grid)
				peekedBotL  = peekToBottomLeft(x, y, grid)
			)

			if isXMAS(peekedRight.values) {
				foundXMAS(x, y, "Right", peekedRight.coords, grid)
				coordsUsed = append(coordsUsed, peekedRight.coords)
				xmasCount++
			}

			if isXMAS(peekedLeft.values) {
				foundXMAS(x, y, "Left", peekedLeft.coords, grid)
				coordsUsed = append(coordsUsed, peekedLeft.coords)
				xmasCount++
			}

			if isXMAS(peekedUp.values) {
				foundXMAS(x, y, "Up", peekedUp.coords, grid)
				coordsUsed = append(coordsUsed, peekedUp.coords)
				xmasCount++
			}

			if isXMAS(peekedDown.values) {
				foundXMAS(x, y, "Down", peekedDown.coords, grid)
				coordsUsed = append(coordsUsed, peekedDown.coords)
				xmasCount++
			}

			if isXMAS(peekedTopL.values) {
				foundXMAS(x, y, "TopLeft", peekedTopL.coords, grid)
				coordsUsed = append(coordsUsed, peekedTopL.coords)
				xmasCount++
			}

			if isXMAS(peekedTopR.values) {
				foundXMAS(x, y, "TopRight", peekedTopR.coords, grid)
				coordsUsed = append(coordsUsed, peekedTopR.coords)
				xmasCount++
			}

			if isXMAS(peekedBotR.values) {
				foundXMAS(x, y, "BotRight", peekedBotR.coords, grid)
				coordsUsed = append(coordsUsed, peekedBotR.coords)
				xmasCount++
			}

			if isXMAS(peekedBotL.values) {
				foundXMAS(x, y, "BotLeft", peekedBotL.coords, grid)
				coordsUsed = append(coordsUsed, peekedBotL.coords)
				xmasCount++
			}
		}
	}

	return xmasCount, coordsUsed
}

func foundXMAS(x, y int, dir string, coords [4]xycoord, grid [][]uint8) {
	//fmt.Printf(
	//	"FOUND XMAS: POS(%d,%d) (DIR=%s) %+v\n",
	//	x,
	//	y,
	//	dir,
	//	coords,
	//)
}

func foundCoordsToGrid(foundCoords [4]xycoord, grid [][]uint8) string {
	out := make([]rune, 0, 4)

	for k := 0; k < 4; k++ {
		out = append(out, rune(grid[foundCoords[k].x][foundCoords[k].y]))
	}

	return string(out)
}

func peekY(y int, row []uint8) uint8 {
	if y >= 0 && y < len(row) {
		return row[y]
	}

	return 0
}

func peekXY(x, y int, grid [][]uint8) uint8 {
	if x < 0 || x >= len(grid) {
		return 0
	}

	if y < 0 || y >= len(grid[x]) {
		return 0
	}

	//fmt.Printf("PEEK (x=%d, y=%d): %d\n", x, y, grid[x][y])
	return grid[x][y]
}

func peekToRight(x, y int, row []uint8) peekResult {
	return peekResult{
		[4]uint8{
			peekY(y+0, row),
			peekY(y+1, row),
			peekY(y+2, row),
			peekY(y+3, row),
		}, [4]xycoord{
			{x, y + 0},
			{x, y + 1},
			{x, y + 2},
			{x, y + 3},
		},
	}
}

func peekToLeft(x, y int, row []uint8) peekResult {
	return peekResult{
		[4]uint8{
			peekY(y-0, row),
			peekY(y-1, row),
			peekY(y-2, row),
			peekY(y-3, row),
		}, [4]xycoord{
			{x, y - 0},
			{x, y - 1},
			{x, y - 2},
			{x, y - 3},
		},
	}
}

func peekDown(x, y int, grid [][]uint8) peekResult {
	return peekResult{
		[4]uint8{
			peekXY(x+0, y, grid),
			peekXY(x+1, y, grid),
			peekXY(x+2, y, grid),
			peekXY(x+3, y, grid),
		},
		[4]xycoord{
			{x + 0, y},
			{x + 1, y},
			{x + 2, y},
			{x + 3, y},
		},
	}

}

func peekUp(x, y int, grid [][]uint8) peekResult {
	return peekResult{
		[4]uint8{
			peekXY(x-0, y, grid),
			peekXY(x-1, y, grid),
			peekXY(x-2, y, grid),
			peekXY(x-3, y, grid),
		},
		[4]xycoord{
			{x - 0, y},
			{x - 1, y},
			{x - 2, y},
			{x - 3, y},
		},
	}
}

func peekToTopRight(x, y int, grid [][]uint8) peekResult {
	return peekResult{
		[4]uint8{
			peekXY(x-0, y+0, grid),
			peekXY(x-1, y+1, grid),
			peekXY(x-2, y+2, grid),
			peekXY(x-3, y+3, grid),
		}, [4]xycoord{
			{x - 0, y + 0},
			{x - 1, y + 1},
			{x - 2, y + 2},
			{x - 3, y + 3},
		},
	}
}

func peekToBottomRight(x, y int, grid [][]uint8) peekResult {
	return peekResult{
		[4]uint8{
			peekXY(x+0, y+0, grid),
			peekXY(x+1, y+1, grid),
			peekXY(x+2, y+2, grid),
			peekXY(x+3, y+3, grid),
		}, [4]xycoord{
			{x + 0, y + 0},
			{x + 1, y + 1},
			{x + 2, y + 2},
			{x + 3, y + 3},
		},
	}
}

func peekToBottomLeft(x, y int, grid [][]uint8) peekResult {
	return peekResult{
		[4]uint8{
			peekXY(x+0, y-0, grid),
			peekXY(x+1, y-1, grid),
			peekXY(x+2, y-2, grid),
			peekXY(x+3, y-3, grid),
		}, [4]xycoord{
			{x + 0, y - 0},
			{x + 1, y - 1},
			{x + 2, y - 2},
			{x + 3, y - 3},
		},
	}
}

func peekToTopLeft(x, y int, grid [][]uint8) peekResult {
	return peekResult{
		[4]uint8{
			peekXY(x-0, y-0, grid),
			peekXY(x-1, y-1, grid),
			peekXY(x-2, y-2, grid),
			peekXY(x-3, y-3, grid),
		}, [4]xycoord{
			{x - 0, y - 0},
			{x - 1, y - 1},
			{x - 2, y - 2},
			{x - 3, y - 3},
		},
	}
}

func isXMAS(peeked [4]uint8) bool {
	if peeked[0] != 'X' {
		return false
	}

	if peeked[1] != 'M' {
		return false
	}

	if peeked[2] != 'A' {
		return false
	}

	if peeked[3] != 'S' {
		return false
	}

	return true
}
