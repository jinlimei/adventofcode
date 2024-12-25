package day05

import (
	"fmt"
	"github.com/jinlimei/adventofcode/golang/library/util"
)

/*
--- Part Two ---

While the Elves get to work printing the correctly-ordered updates, you have a little time to fix the rest of them.

For each of the incorrectly-ordered updates, use the page ordering rules to put the page numbers in the right order.
For the above example, here are the three incorrectly-ordered updates and their correct orderings:

    75,97,47,61,53 becomes 97,75,47,61,53.
    61,13,29 becomes 61,29,13.
    97,13,75,29,47 becomes 97,75,47,29,13.

After taking only the incorrectly-ordered updates and ordering them correctly, their middle page numbers are 47, 29,
and 47. Adding these together produces 123.

Find the updates which are not in the correct order. What do you get if you add up the middle page numbers after
correctly ordering just those updates?
*/

func (d Day) Part2Prompt() {
	const input = `
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`

	rules, updates := parse(input)
	//spew.Dump(rules)

	_, invalid := validate(rules, updates)
	fmt.Println("\nINVALID:")
	for _, v := range invalid {
		fmt.Println(v)
	}

	fixed := fixInvalid(rules, invalid)

	fmt.Println("\nFIXED:")
	for _, v := range fixed {
		fmt.Println(v)
	}

	fmt.Println("\nFIXED SUMS", sumMids(fixed))
}

func (d Day) Part2Actual() {
	input, err := util.ReadInputFile(2024, 5)

	if err != nil {
		panic(err)
	}

	rules, updates := parse(input)

	_, invalid := validate(rules, updates)

	fixed := fixInvalid(rules, invalid)

	fmt.Println(sumMids(fixed))
}
