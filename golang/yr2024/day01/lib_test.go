package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	left, right := parse(`
3   4
4   3
2   5
1   3
3   9
3   3
`)

	//fmt.Println(left, right)

	assert.Equal(t, 3, left[0])
	assert.Equal(t, 4, right[0])
	assert.Equal(t, 3, left[5])
	assert.Equal(t, 3, right[5])
}

func TestCouple(t *testing.T) {
	left, right := parse(`
3   4
4   3
2   5
1   3
3   9
3   3
`)

	pairs := couple(left, right)

	assert.Equal(t, []int{1, 3}, pairs[0])
	assert.Equal(t, []int{4, 9}, pairs[5])
}

func TestDistance(t *testing.T) {
	left, right := parse(`
3   4
4   3
2   5
1   3
3   9
3   3
`)

	pairs := couple(left, right)

	dist := distance(pairs)

	assert.Equal(t, 11, dist)
}
