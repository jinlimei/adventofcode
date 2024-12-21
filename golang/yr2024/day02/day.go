package day02

import (
	"fmt"

	"github.com/jinlimei/adventofcode/golang/library/util"
)

type Day struct {
}

func (d Day) actual(useProblemDamper bool) {
	input, err := util.ReadInputFile(2024, 2)
	if err != nil {
		panic(err)
	}

	safeReports := 0
	reports := parse(input)
	fmt.Printf("Loaded %d reports\n", len(reports))

	for _, report := range reports {
		safe := report.isSafe(useProblemDamper)

		fmt.Printf("Report %+v is safe='%v'\n",
			report.levels,
			safe,
		)

		if safe {
			safeReports++
		}
	}

	fmt.Printf("Total Reports: %d\n", len(reports))
	fmt.Printf("Safe Reports: %d\n", safeReports)
	fmt.Printf("Percentage Safe: %.2f\n", (float64(safeReports)/float64(len(reports)))*100)
}
