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

	final, zeroes, clickZeroes := runRotations(rules)

	assert.Equal(t, 3, zeroes)
	assert.Equal(t, 32, final)
	assert.Equal(t, 6, clickZeroes)
}
