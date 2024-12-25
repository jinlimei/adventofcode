package day05

import (
	"strconv"
	"strings"
)

type orderingRule struct {
	page   int
	before int
}

func parse(input string) ([]orderingRule, [][]int) {
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

		page   int
		before int

		pageRune   = make([]rune, 0, 2)
		beforeRune = make([]rune, 0, 2)

		configs = make([]orderingRule, 0)

		config orderingRule
	)

	for ; pos < rLen; pos++ {
		chr = runes[pos]

		switch chr {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if afterPipe {
				pageRune = append(pageRune, chr)
			} else {
				beforeRune = append(beforeRune, chr)
			}
		case '|':
			afterPipe = true
			before, _ = strconv.Atoi(string(beforeRune))

			config.before = before

			beforeRune = make([]rune, 0, 2)

		case '\n':
			if !afterPipe {
				break
			}

			page, _ = strconv.Atoi(string(pageRune))

			config.page = page

			if page > 0 && before > 0 {
				configs = append(configs, config)
			}

			pageRune = make([]rune, 0, 2)
			beforeRune = make([]rune, 0, 2)

			page = 0
			before = 0
			afterPipe = false
		}
	}

	return configs
}

func parsePageNumberUpdates(input string) [][]int {
	var (
		runes = []rune(input)
		rLen  = len(runes)

		pos = 0
		chr rune

		lines = make([][]int, 0)
		line  = make([]int, 0)

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
