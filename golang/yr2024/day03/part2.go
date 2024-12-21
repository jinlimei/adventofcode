package day03

import (
	"fmt"

	"github.com/jinlimei/adventofcode/golang/library/util"
)

/*
--- Part Two ---

As you scan through the corrupted memory, you notice that some of the conditional statements are also still intact.
If you handle some of the uncorrupted conditional statements in the program, you might be able to get an even more
accurate result.

There are two new instructions you'll need to handle:

    The do() instruction enables future mul instructions.
    The don't() instruction disables future mul instructions.

Only the most recent do() or don't() instruction applies. At the beginning of the program, mul instructions are enabled.

For example:

xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
 mul(2,4)           don't() mul(5,5)            mul(11,8)  do() mul(8,5)
 mul(2,4)                                                       mul(8,5)

This corrupted memory is similar to the example from before, but this time the mul(5,5) and mul(11,8) instructions
are disabled because there is a don't() instruction before them. The other mul instructions function normally,
including the one at the end that gets re-enabled by a do() instruction.

This time, the sum of the results is 48 (2*4 + 8*5).

Handle the new instructions; what do you get if you add up all of the results of just the enabled multiplications?
*/

func (d Day) Part2Prompt() {
	const mem = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

	ops := parse(mem, []op{opMult, opDo, opDont})

	fmt.Println("PART 2 PROMPT:", sumOps(ops))
}

func (d Day) Part2Actual() {
	input, err := util.ReadInputFile(2024, 3)

	if err != nil {
		panic(err)
	}

	ops := parse(input, []op{opMult, opDo, opDont})
	fmt.Printf("Ops: %d\n", len(ops))
	fmt.Printf("Sum of Ops: %d\n", sumOps(ops))
}
