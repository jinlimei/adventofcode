package day02

import (
	"strconv"

	"github.com/jinlimei/adventofcode/golang/library/util"
)

type direction int

const (
	dirUnknown direction = iota
	dirIncreasing
	dirDecreasing
)

type validReport struct {
	levels []int
}

func (vr *validReport) isSafe() bool {
	var (
		pos = 0

		cur  int
		nxt  int
		diff int

		startDir = dirUnknown
		currDir  = dirUnknown

		maxLen = len(vr.levels)
	)

	for ; pos < maxLen-1; pos++ {
		cur = vr.levels[pos]
		if pos+1 < maxLen {
			nxt = vr.levels[pos+1]
		}

		if cur-nxt > 0 {
			currDir = dirIncreasing
		} else {
			currDir = dirDecreasing
		}

		if startDir == dirUnknown {
			startDir = currDir
		} else if startDir != currDir {
			return false
		}

		diff = util.AbsInt(cur - nxt)

		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
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
