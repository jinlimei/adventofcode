package day07

import "log"

type equationOp uint8

const validOpsCount = 2

const (
	eqOpUnknown equationOp = iota
	eqOpAdd
	eqOpMultiply
)

func (op equationOp) String() string {
	switch op {
	case eqOpAdd:
		return "+"
	case eqOpMultiply:
		return "*"
	default:
		panic("invalid equation op")
	}
}

func (op equationOp) Exec(a, b int) int {
	switch op {
	case eqOpAdd:
		return a + b
	case eqOpMultiply:
		return a * b
	default:
		panic("invalid equation op")
	}
}

func validOps() [2]equationOp {
	return [2]equationOp{
		eqOpAdd,
		eqOpMultiply,
	}
}

type equation struct {
	result int
	nums   []int
	ops    []equationOp
	valid  bool
}

func newEquation() *equation {
	return &equation{
		nums: make([]int, 0),
		ops:  make([]equationOp, 0),
	}
}

func (eq *equation) solve() {
	var (
		nLen = len(eq.nums)
		pos  = 0

		n1 int
		n2 int

		vOps = validOps()
	)

	for ; pos < nLen-1; pos++ {
		n1 = eq.nums[pos]
		n2 = eq.nums[pos+1]

		for k := 0; k < validOpsCount; k++ {
			start := vOps[k].Exec(n1, n2)
			log.Printf("test: %v", start)

			for i := pos; i < nLen-1; i++ {
				n1 = eq.nums[i]
				n2 = eq.nums[i+1]
			}
		}
	}
}
