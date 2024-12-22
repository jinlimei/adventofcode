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
