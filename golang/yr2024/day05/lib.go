package day05

import (
	"slices"
)

// validate operations for part1 (at time of writing, anyway, part2 may be different)
//
// 1. Identify which updates are already in the *right* order.
func validate(rules []orderingRule, updates []update) ([]update, []update) {
	var (
		goodOut = make([]update, 0)
		badOut  = make([]update, 0)
	)

	for _, u := range updates {
		var (
			pos        int
			uLen       = len(u)
			anyInvalid = false
		)

	ruleLoop:
		for _, rule := range rules {
			for pos = 0; pos < uLen; pos++ {
				//fmt.Printf("Number: %d\n", u[pos])
				// We found one! Let's see if this num1 is *num0* .num0
				if rule.num0 == u[pos] {
					//fmt.Printf("rule Check '%s', %d must be *before* %d\n", rule, rule.num0, rule.num1)
					//fmt.Printf("using num0 %d\n", rule.num0)
					// We need to check all numbers preceding num0 here
					// So we start at 0 and move towards current pos.
					for k := 0; k < pos; k++ {
						//fmt.Printf("checking %d\n", u[k])
						if rule.num1 == u[k] {
							anyInvalid = true
							break ruleLoop
						}
					}
				}
			}
		}

		if anyInvalid {
			badOut = append(badOut, u)
		} else {
			goodOut = append(goodOut, u)
		}

	}

	return goodOut, badOut
}

func fixInvalid(rules []orderingRule, invalidUpdates []update) []update {
	for _, invalid := range invalidUpdates {
		slices.SortFunc(invalid, func(num0, num1 int) int {
			for _, rule := range rules {
				if rule.num0 == num0 && rule.num1 == num1 {
					return -1
				}
			}

			return 0
		})
	}

	return invalidUpdates
}

func getMid(u update) int {
	return u[len(u)/2]
}

func sumMids(updates []update) int {
	out := 0
	for _, u := range updates {
		out += getMid(u)
	}

	return out
}
