package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleRotations(t *testing.T) {
	assert.True(t, true)

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
	assert.Equal(t, 10, len(rules))

	final, zeroes := runRotations(rules)

	assert.Equal(t, 3, zeroes)
	assert.Equal(t, 32, final)
}

func TestRunClickRotations1(t *testing.T) {
	var testInput = `
L48
L3
R6
`

	rules := parseRotationFile(testInput)
	assert.Equal(t, 3, len(rules))

	final, clickZeroes := runClickRotations(rules)
	assert.Equal(t, 5, final)
	assert.Equal(t, 2, clickZeroes)
}

func TestRunClickRotations2(t *testing.T) {
	var testInput = `
L50
L5
`
	rules := parseRotationFile(testInput)
	assert.Equal(t, 2, len(rules))

	final, clickZeroes := runClickRotations(rules)
	assert.Equal(t, 95, final)
	assert.Equal(t, 1, clickZeroes)
}

func TestRunClickRotations3(t *testing.T) {
	var testInput = `
R1000
`

	rules := parseRotationFile(testInput)
	assert.Equal(t, 1, len(rules))

	final, clickZeroes := runClickRotations(rules)
	assert.Equal(t, 50, final)
	assert.Equal(t, 10, clickZeroes)
}

func TestRunClickRotations4(t *testing.T) {
	var testInput = `
L68
L30
R48
`
	rules := parseRotationFile(testInput)
	assert.Equal(t, 3, len(rules))

	final, clickZeroes := runClickRotations(rules)
	assert.Equal(t, 0, final)
	assert.Equal(t, 2, clickZeroes)
}

func TestRunClickRotations5(t *testing.T) {
	var testInput = `
L50
L5
R60
L55
`

	rules := parseRotationFile(testInput)
	assert.Equal(t, 4, len(rules))

	final, clickZeroes := runClickRotations(rules)
	assert.Equal(t, 0, final)
	assert.Equal(t, 3, clickZeroes)
}

func TestRunClickRotations6(t *testing.T) {
	var testInput = `
L50
L1
L99
R14
L82
`

	rules := parseRotationFile(testInput)
	assert.Equal(t, 5, len(rules))

	final, clickZeroes := runClickRotations(rules)
	assert.Equal(t, 32, final)
	assert.Equal(t, 3, clickZeroes)
}
