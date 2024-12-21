package day02

import "github.com/jinlimei/adventofcode/golang/library/util"

type direction int

const (
	dirUnknown direction = iota
	dirIncreasing
	dirDecreasing
)

type validReport struct {
	levels []int
}

func (vr *validReport) isSafeWithDamper() bool {
	ok, _ := isValidReport(vr.levels)
	// Simple case: it's already safe without any drastic wildness
	if ok {
		return true
	}

	for pos := 0; pos < len(vr.levels); pos++ {
		check := buildIntSliceWithoutElement(vr.levels, pos)

		ok, _ = isValidReport(check)

		if ok {
			return true
		}
	}
	
	return false
}

func (vr *validReport) isSafeSimple() bool {
	ok, _ := isValidReport(vr.levels)

	return ok
}

func isValidReport(levels []int) (bool, int) {
	var (
		pos      = 0
		cur      int
		nxt      int
		diff     int
		startDir = dirUnknown
		currDir  = dirUnknown
		maxLen   = len(levels)
	)

	for ; pos < maxLen-1; pos++ {
		cur = levels[pos]

		if pos+1 < maxLen {
			nxt = levels[pos+1]
		}

		if cur-nxt > 0 {
			currDir = dirIncreasing
		} else {
			currDir = dirDecreasing
		}

		if startDir == dirUnknown {
			startDir = currDir
		} else if startDir != currDir {
			return false, pos
		}

		diff = util.AbsInt(cur - nxt)

		if diff < 1 || diff > 3 {
			return false, pos
		}
	}

	return true, -1
}
