package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	ops := parse("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	//spew.Dump(ops)

	//fmt.Println(sum)
	assert.Equal(t, 161, sumOps(ops))
}
