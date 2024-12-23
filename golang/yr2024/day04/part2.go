package day04

import (
	"fmt"

	"github.com/jinlimei/adventofcode/golang/library/util"
)

/*
--- Part Two ---

The Elf looks quizzically at you. Did you misunderstand the assignment?

Looking for the instructions, you flip over the word search to find that this isn't actually an XMAS puzzle;
it's an X-MAS puzzle in which you're supposed to find two MAS in the shape of an X. One way to achieve that
is like this:

M.S
.A.
M.S

Irrelevant characters have again been replaced with . in the above diagram. Within the X, each MAS can be written
forwards or backwards.

Here's the same example from before, but this time all of the X-MASes have been kept instead:

.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........

In this example, an X-MAS appears 9 times.

Flip the word search from the instructions back over to the word search side and try again. How many times
does an X-MAS appear?
*/

func (d Day) Part2Prompt() {
	const input = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

	grid := parse(input)
	visualize(grid)

	fmt.Println()
	fmt.Println()

	masCount, coordsUsed := scanGridForMasThatIsX(grid)
	visualizeMasUsed(grid, coordsUsed)

	fmt.Printf("\n\nNumber Of MAS that is X: %d\n", masCount)
}

func (d Day) Part2Actual() {
	input, err := util.ReadInputFile(2024, 4)

	if err != nil {
		panic(err)
	}

	grid := parse(input)

	masCount, _ := scanGridForMasThatIsX(grid)

	fmt.Printf("\n\nNumber Of MAS taht is X: %d\n", masCount)
}
