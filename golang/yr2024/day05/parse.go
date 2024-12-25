package day05

import (
	"fmt"
	"strconv"
	"strings"
)

type update []int

func (u update) String() string {
	var (
		buf  strings.Builder
		uLen = len(u)
	)

	for idx, val := range u {
		buf.WriteString(strconv.Itoa(val))
		if idx+1 < uLen {
			buf.WriteRune(',')
		}
	}

	return buf.String()
}

type orderingRule struct {
	num0 int
	num1 int
}

func (or orderingRule) String() string {
	return fmt.Sprintf("%d|%d", or.num0, or.num1)
}

func parse(input string) ([]orderingRule, []update) {
	parts := strings.Split(input, "\n\n")

	return parseOrderingRules(parts[0]),
		parsePageNumberUpdates(parts[1])
}

func parseOrderingRules(input string) []orderingRule {
	var (
		runes = []rune(input)
		rLen  = len(runes)

		pos = 0
		chr rune

		afterPipe = false

		num1 int
		num0 int

		num0Rune = make([]rune, 0, 2)
		num1Rune = make([]rune, 0, 2)

		configs = make([]orderingRule, 0)

		config orderingRule
	)

	for ; pos < rLen; pos++ {
		chr = runes[pos]

		switch chr {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if afterPipe {
				num0Rune = append(num0Rune, chr)
			} else {
				num1Rune = append(num1Rune, chr)
			}
		case '|':
			afterPipe = true
			num0, _ = strconv.Atoi(string(num1Rune))

			config.num0 = num0
		case '\n':
			if !afterPipe {
				break
			}

			num1, _ = strconv.Atoi(string(num0Rune))

			config.num1 = num1

			if num1 > 0 && num0 > 0 {
				configs = append(configs, config)
			}

			num0Rune = make([]rune, 0, 2)
			num1Rune = make([]rune, 0, 2)

			num1 = 0
			num0 = 0
			afterPipe = false
		}
	}

	return configs
}

func parsePageNumberUpdates(input string) []update {
	var (
		runes = []rune(input)
		rLen  = len(runes)

		pos = 0
		chr rune

		lines = make([]update, 0)
		line  = make(update, 0)

		numRunes = make([]rune, 0, 2)
		num      int
	)

	for ; pos < rLen; pos++ {
		chr = runes[pos]

		switch chr {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			numRunes = append(numRunes, chr)
		case ',':
			num, _ = strconv.Atoi(string(numRunes))
			line = append(line, num)
			numRunes = make([]rune, 0, 2)
		case '\n':
			if len(numRunes) == 0 || len(line) == 0 {
				break
			}

			num, _ = strconv.Atoi(string(numRunes))
			line = append(line, num)
			numRunes = make([]rune, 0, 2)

			lines = append(lines, line)
			line = make([]int, 0)
		}
	}

	return lines
}
