package day02

import (
	"fmt"
	"strconv"

	"github.com/jinlimei/adventofcode/golang/library/util"
)

type validReport struct {
	levels []int
}

func (vr *validReport) isSafe(useProblemDampener bool) bool {
	var (
		pos = 0

		cur     int
		nxt     int
		diff    int
		diffAbs int

		maxLen = len(vr.levels)

		diffs = make([]int, 0)

		numInvalid int
		numValid   int

		numZero     int
		numPositive int
		numNegative int
	)

	for ; pos < maxLen-1; pos++ {
		cur = vr.levels[pos]
		nxt = vr.levels[pos+1]

		diff = cur - nxt
		diffs = append(diffs, diff)

		diffAbs = util.AbsInt(diff)

		if diffAbs != 0 && (diffAbs < 1 || diffAbs > 3) {
			numInvalid++
		} else {
			numValid++
		}

		if diff > 0 {
			numPositive++
		} else if diff < 0 {
			numNegative++
		} else {
			numZero++
		}
	}

	fmt.Printf("\n  Nums invalid=%d, valid=%d, positive=%d, negative=%d, zero=%d\n", numInvalid, numValid, numPositive, numNegative, numZero)

	var (
		NegPosOkay   bool
		InvValidOkay = numInvalid == 0
	)

	if !useProblemDampener {
		if numZero == 0 && numNegative == 0 && numPositive > 0 {
			NegPosOkay = true
		} else if numZero == 0 && numNegative > 0 && numPositive == 0 {
			NegPosOkay = true
		}

	} else {

		if numZero == 0 && numNegative <= 1 && numPositive > 0 {
			NegPosOkay = true
		} else if numZero == 0 && numNegative > 0 && numPositive <= 1 {
			NegPosOkay = true
		} else if numZero <= 1 && numNegative == 0 && numPositive > 0 {
			NegPosOkay = true
		} else if numZero <= 1 && numNegative > 0 && numPositive == 0 {
			NegPosOkay = true
		}
	}

	return NegPosOkay && InvValidOkay
}

func a() {
	fmt.Print("a")
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
