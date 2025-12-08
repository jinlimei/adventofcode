package day01

import (
	"log"

	"github.com/jinlimei/adventofcode/golang/library/util"
)

func (d Day) Part2Prompt() {

	var testInput = `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

	rules := parseRotationFile(testInput)
	final, clickZeroes := runClickRotations(rules)

	log.Printf("Zeroes: %d", clickZeroes)
	log.Printf("Final: %d", final)
}

func (d Day) Part2Actual() {
	input, err := util.ReadInputFile(2025, 1)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	rules := parseRotationFile(input)
	log.Printf("retrieved %d rules", len(rules))

	final, clickZeroes := runClickRotations(rules)
	log.Printf("final result is %d, click zeroes %d", final, clickZeroes)
}
