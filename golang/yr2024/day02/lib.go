package day02

import (
	"strconv"
)

func buildIntSliceWithoutElement(old []int, skipEle int) []int {
	var (
		oLen = len(old)
		out  = make([]int, 0, oLen-1)
		pos  = 0
	)

	for ; pos < oLen; pos++ {
		if pos == skipEle {
			continue
		}

		out = append(out, old[pos])
	}

	return out
}

func parse(input string) []validReport {
	var (
		reports = make([]validReport, 0)
		report  = validReport{
			levels: make([]int, 0),
		}

		runes = []rune(input)
		chr   rune

		pos  = 0
		rLen = len(runes)

		numRunes = make([]rune, 0)
		num      int
	)

	for ; pos < rLen; pos++ {
		chr = runes[pos]

		switch chr {
		case ' ', '\n':
			if len(numRunes) == 0 {
				break
			}

			num, _ = strconv.Atoi(string(numRunes))
			report.levels = append(report.levels, num)

			numRunes = make([]rune, 0)
			num = 0

			if chr == '\n' {
				reports = append(reports, report)
				report = validReport{}
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			numRunes = append(numRunes, chr)
		}
	}

	return reports
}
