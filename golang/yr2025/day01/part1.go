package day01

import (
	"log"

	"github.com/jinlimei/adventofcode/golang/library/util"
)

func (d Day) Part1Prompt() {

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
	final, zeroes, _ := runRotations(rules)

	log.Printf("Zeroes: %02d", zeroes)
	log.Printf("Final: %02d", final)
}

func (d Day) Part1Actual() {
	input, err := util.ReadInputFile(2025, 1)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	rules := parseRotationFile(input)
	log.Printf("retrieved %d rules", len(rules))

	final, zeroes, _ := runRotations(rules)
	log.Printf("final result is %d, zeroes %d", final, zeroes)
}
