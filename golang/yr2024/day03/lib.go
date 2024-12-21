package day03

import (
	"fmt"
	"strconv"
)

type op int

const (
	opUnk op = iota
	opMult
	opDo
	opDont
)

func (op op) String() string {
	switch op {
	case opMult:
		return "mul"
	case opDo:
		return "do"
	case opDont:
		return "don't"
	default:
		return "unk"
	}
}

type instruction struct {
	op op
	a1 int64
	a2 int64
}

func (i instruction) String() string {
	return fmt.Sprintf("%s(%d,%d)", i.op, i.a1, i.a2)
}

func (i instruction) do() int64 {
	switch i.op {
	case opMult:
		return i.a1 * i.a2
	default:
		return -1
	}
}

func sumOps(ops []*instruction) int64 {
	var (
		sum = int64(0)
		do  = true
	)

	for _, currOp := range ops {
		switch currOp.op {
		case opMult:
			if do {
				sum += currOp.do()
			}
		case opDo:
			do = true
		case opDont:
			do = false
		default:
			panic("unhandled default case")
		}
	}

	return sum
}

func parse(input string, validOps []op) []*instruction {
	var (
		runes = []rune(input)
		rLen  = len(runes)
	)

	//#region inline functions

	// inline functions
	peek := func(loc int) rune {
		if loc >= 0 && loc < rLen {
			return runes[loc]
		}

		return '\u0000'
	}

	findAt := func(at int, s string) bool {
		var (
			seq  = []rune(s)
			sLen = len(seq)

			i int
		)

		for i = 0; i < sLen; i++ {
			if runes[at+i] != seq[i] {
				return false
			}
		}

		return true
	}

	//seekUntilAny := func(at int, untilAny ...rune) int {
	//	var (
	//		i = 0
	//		p = 0
	//		m = len(untilAny)
	//	)
	//
	//	for i = at; i < rLen; i++ {
	//		for p = 0; p < m; p++ {
	//			if peek(i) == untilAny[p] {
	//				return i
	//			}
	//		}
	//	}
	//
	//	return i
	//}

	seekUntilInvalid := func(at int) (*instruction, int) {
		var (
			i   = 0
			chr rune

			a1 = make([]rune, 0)
			a2 = make([]rune, 0)

			isA1 = true

			a1i int64
			a2i int64
		)

	seekLoop:
		for i = at; i < rLen; i++ {
			chr = runes[i]
			//fmt.Printf("chr '%s' ", string(chr))

			switch chr {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if isA1 {
					a1 = append(a1, chr)
				} else {
					a2 = append(a2, chr)
				}

			case ',':
				isA1 = false
			case ')':
				break seekLoop
			default:
				//fmt.Println()
				return nil, i
			}
		}

		//fmt.Println()

		// a1 and a2 *must* be a maximum of 3 digits
		if len(a1) == 0 || len(a1) > 3 || len(a2) == 0 || len(a2) > 3 {
			fmt.Printf("NEARLY VALID MUL FOUND AT POS '%d', mul(%s,%s)\n", i, string(a1), string(a2))
			fmt.Printf("str? %s\n", string(runes[at-8:i+6]))
			return nil, i
		}

		a1i, _ = strconv.ParseInt(string(a1), 10, 64)
		a2i, _ = strconv.ParseInt(string(a2), 10, 64)

		return &instruction{opMult, a1i, a2i}, i
	}

	//#endregion

	var (
		pos = 0
		nxt int

		chr  rune
		inst *instruction

		out = make([]*instruction, 0)
	)

	for pos = 0; pos < rLen; pos++ {
		chr = peek(pos)

		//fmt.Printf("MAIN '%s' ", string(chr))

		switch chr {
		case 'm':
			// We'll short out and move on if mul( isn't here
			if !findAt(pos, "mul(") {
				break
			}

			//fmt.Printf("found mul( ")
			inst, nxt = seekUntilInvalid(pos + 4)

			if inst != nil {
				out = append(out, inst)
			} else {
				fmt.Printf("found mul( but failed the rest at pos '%d': '%s'\n", pos, string(runes[pos:pos+8]))
			}

			pos = nxt
		case 'd':
			if isValidOp(opDont, validOps) && findAt(pos, "don't()") {
				out = append(out, &instruction{opDont, 0, 0})
				break
			}

			if isValidOp(opDo, validOps) && findAt(pos, "do()") {
				out = append(out, &instruction{opDo, 0, 0})
				break
			}
		}

		//fmt.Println()
	}

	return out
}

func isValidOp(inp op, validOps []op) bool {
	var (
		pos  = 0
		vLen = len(validOps)
	)

	for ; pos < vLen; pos++ {
		if inp == validOps[pos] {
			//fmt.Printf("op %s is valid\n", inp)
			return true
		}
	}

	return false
}
