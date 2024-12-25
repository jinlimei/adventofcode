package day05

import "github.com/davecgh/go-spew/spew"

func validate(rules []orderingRule, updates [][]int) []string {
	var (
		issues = make([]string, 0)
	)

	spew.Dump(rules[0])

	return issues
}
