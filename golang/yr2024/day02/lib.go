package day02

import (
	"fmt"
	"strconv"

	"github.com/jinlimei/adventofcode/golang/library/util"
)

type direction int

func (d direction) String() string {
	switch d {
	case 0:
		return "unknown"
	case 1:
		return "increasing"
	case 2:
		return "decreasing"
	default:
		return "undefined"
	}
}

const (
	dirUnknown direction = iota
	dirIncreasing
	dirDecreasing
)

type validReport struct {
	levels []int
}

func (vr *validReport) hasProblem(baseDir direction, cur, nxt int) bool {
	var (
		currDir = dirUnknown
		diff    = cur - nxt
		absDiff = util.AbsInt(diff)
	)

	currDir = vr.getDirection(cur, nxt)

	if baseDir != currDir {
		return true
	}

	if absDiff < 1 || absDiff > 3 {
		return true
	}

	return false
}

func (vr *validReport) unsafeCount(useProblemDamper bool) int {
	var (
		pos = 0

		cur int
		nxt int
		fut int

		nxtDir = dirUnknown
		futDir = dirUnknown

		maxLen = len(vr.levels)

		problems = 0
	)

	for ; pos < maxLen-1; pos++ {
		cur = vr.levels[pos]
		nxt = -1
		fut = -1

		if pos+1 < maxLen {
			nxt = vr.levels[pos+1]
		}

		if pos+2 < maxLen {
			fut = vr.levels[pos+2]
		}

		if nxtDir == dirUnknown {
			nxtDir = vr.getDirection(cur, nxt)
			futDir = vr.getDirection(cur, fut)
		}

		// Do we have a problem?
		if vr.hasProblem(nxtDir, cur, nxt) {
			//fmt.Printf("  Problem detected (dir=%+v, cur=%d, nxt=%d)\n", nxtDir, cur, nxt)
			//
			//fmt.Printf("  Problem using futDir? %v\n", vr.hasProblem(futDir, cur, nxt))
			//fmt.Printf("  Problem using fut? %v\n", vr.hasProblem(nxtDir, cur, fut))

			// If we aren't using the problem damper _or_ we can't skip 'nxt'
			if !useProblemDamper || fut == -1 {
				fmt.Println("  Recording problem (!useProblemDamper or fut==-1)")
				problems++
				continue
			}

			// If we still have a problem using fut instead of nxt
			// Special case: pos=0, so 'nxtDir' is invalid if we're skipping nxt
			if pos == 0 && vr.hasProblem(futDir, cur, fut) {
				problems++
				continue
			}

			// If we still have a problem using fut instead of nxt
			// Special case: pos>0, so 'nxtDir' is valid
			if pos > 0 && vr.hasProblem(nxtDir, cur, fut) {
				problems++
				continue
			}
		}
	}

	return problems
}

func (vr *validReport) getDirection(cur int, nxt int) direction {
	var currDir direction

	//fmt.Printf("getDirection(%d, %d) = %d\n", cur, nxt, cur-nxt)
	if cur-nxt < 0 {
		currDir = dirIncreasing
	} else {
		currDir = dirDecreasing
	}

	return currDir
}

func (vr *validReport) isSafe(useProblemDampener bool) bool {
	return vr.unsafeCount(useProblemDampener) == 0
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
