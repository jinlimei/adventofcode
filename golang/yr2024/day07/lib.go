package day07

import (
	"strconv"
	"strings"
)

func parse(input string) []*equation {
	input = strings.TrimSpace(input) + "\n"

	var (
		runes = []rune(input)
		rLen  = len(runes)

		pos = 0
		buf = make([]rune, 0)
		chr rune

		num int

		eqs = make([]*equation, 0)
		eq  = newEquation()
	)

	for ; pos < rLen; pos++ {
		chr = runes[pos]

		switch chr {
		case ':':
			num, _ = strconv.Atoi(string(buf))
			buf = make([]rune, 0)
			eq.result = num

			// we'll skip the space after ':'
			if pos+1 < rLen && runes[pos+1] == ' ' {
				pos++
			}

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			buf = append(buf, chr)
		case ' ', '\n':
			num, _ = strconv.Atoi(string(buf))
			buf = make([]rune, 0)
			eq.nums = append(eq.nums, num)

			if chr == '\n' {
				eqs = append(eqs, eq)
				eq = newEquation()
			}
		}
	}

	return eqs
}
