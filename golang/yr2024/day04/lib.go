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

func visualize(grid [][]uint8) {
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			fmt.Printf("%s", string(grid[x][y]))
		}

		fmt.Println()
	}
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

func visualizeUsed(grid [][]uint8, coordsUsed [][4]xycoord) {
	used := coordsUsedToRawCoords(coordsUsed)

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if inXYCoordSlice(y, x, used) {
				fmt.Printf("%s", string(grid[x][y]))
			} else {
				fmt.Printf(" ")
			}
		}

		fmt.Println()
	}
}

type xycoord struct {
	x int
	y int
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
				peeked4R   = peek4R(x, y, row)
				peeked4L   = peek4L(x, y, row)
				peekedTB   = peek4T(x, y, grid)
				peekedBT   = peek4B(x, y, grid)
				peekedTRBL = peek4DiagPos(x, y, grid)
				peekedBLTR = peek4DiagNeg(x, y, grid)
			)

			if isXMAS(peeked4R.values) {
				coordsUsed = append(coordsUsed, peeked4R.coords)
				xmasCount++
			}

			if isXMAS(peeked4L.values) {
				coordsUsed = append(coordsUsed, peeked4L.coords)
				xmasCount++
			}

			if isXMAS(peekedTB.values) {
				coordsUsed = append(coordsUsed, peekedTB.coords)
				xmasCount++
			}

			if isXMAS(peekedBT.values) {
				coordsUsed = append(coordsUsed, peekedBT.coords)
				xmasCount++
			}

			if isXMAS(peekedTRBL.values) {
				coordsUsed = append(coordsUsed, peekedTRBL.coords)
				xmasCount++
			}

			if isXMAS(peekedBLTR.values) {
				coordsUsed = append(coordsUsed, peekedBLTR.coords)
				xmasCount++
			}
		}
	}

	return xmasCount, coordsUsed
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

func peek4R(x, y int, row []uint8) peekResult {
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

func peek4L(x, y int, row []uint8) peekResult {
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

func peek4T(x, y int, grid [][]uint8) peekResult {
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

func peek4B(x, y int, grid [][]uint8) peekResult {
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

func peek4DiagPos(x, y int, grid [][]uint8) peekResult {
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

func peek4DiagNeg(x, y int, grid [][]uint8) peekResult {
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
