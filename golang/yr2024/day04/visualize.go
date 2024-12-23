package day04

import (
	"fmt"
)

func visualizeUsed(grid [][]uint8, coordsUsed [][4]xycoord) {
	used := coordsUsedToRawCoords(coordsUsed)

	for x := 0; x < len(grid); x++ {
		if x == 0 {
			fmt.Print("   ")

			for y := 0; y < len(grid[0]); y++ {
				fmt.Printf("%02d ", y)
			}

			fmt.Println()
		}

		fmt.Printf("%02d ", x)

		for y := 0; y < len(grid[x]); y++ {
			if inXYCoordSlice(x, y, used) {
				fmt.Printf(" %s ", string(grid[x][y]))
			} else {
				fmt.Printf("   ")
			}
		}

		fmt.Println()
	}
}

func visualizeMasUsed(grid [][]uint8, peeked []masThatIsXPeekResult) {
	coordsUsed := make([]xycoord, 0, len(peeked)*5)

	for _, peek := range peeked {
		for k := 0; k < 5; k++ {
			p := peek.coords[k]

			if !inXYCoordSlice(p.x, p.y, coordsUsed) {
				coordsUsed = append(coordsUsed, p)
			}
		}
	}

	for x := 0; x < len(grid); x++ {
		if x == 0 {
			fmt.Print("   ")

			for y := 0; y < len(grid[0]); y++ {
				fmt.Printf("%02d ", y)
			}

			fmt.Println()
		}

		fmt.Printf("%02d ", x)

		for y := 0; y < len(grid[x]); y++ {
			if inXYCoordSlice(x, y, coordsUsed) {
				fmt.Printf(" %s ", string(grid[x][y]))
			} else {
				fmt.Printf("   ")
			}
		}

		fmt.Println()
	}
}

func visualize(grid [][]uint8) {
	for x := 0; x < len(grid); x++ {
		if x == 0 {
			fmt.Print("   ")

			for y := 0; y < len(grid[0]); y++ {
				fmt.Printf("%02d ", y)
			}

			fmt.Println()
		}

		fmt.Printf("%02d ", x)

		for y := 0; y < len(grid[x]); y++ {
			fmt.Printf(" %s ", string(grid[x][y]))
		}

		fmt.Println()
	}
}
