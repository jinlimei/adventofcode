package day01

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	dialStartPosition = 50
	maxDial           = 100
)

type Direction uint8

const (
	DirectionL Direction = 'L'
	DirectionR Direction = 'R'
)

func (d Direction) String() string {
	switch d {
	case DirectionL:
		return "L"
	case DirectionR:
		return "R"
	default:
		return "IDK"
	}
}

type DialRule struct {
	Direction Direction
	Rotate    int
}

func parseRotationFile(str string) []DialRule {
	str = strings.TrimSpace(str)
	lines := strings.Split(str, "\n")

	rules := make([]DialRule, 0, len(lines))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		rotate, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		rules = append(rules, DialRule{
			Direction: Direction(line[0]),
			Rotate:    rotate,
		})
	}

	return rules
}

func runRotations(rules []DialRule) (int, int, int) {
	dial := dialStartPosition

	zeroes := 0
	clickZeroes := 0

	for _, rule := range rules {
		fmt.Printf("dial at %02d - ", dial)

		switch rule.Direction {
		case DirectionL:
			dial = dial - rule.Rotate

			// apparently golang modulus does not handle negative numbers in
			// the way we want it to, thus we do this to make sure we return
			// to the positive number space.
			if dial <= 0 {
				dial = maxDial + dial
			}
		case DirectionR:
			dial = dial + rule.Rotate

			if dial >= maxDial {
				dial = dial % maxDial
			}
		}

		fmt.Printf("rotated '%s' by %02d to be %02d\n", rule.Direction.String(), rule.Rotate, dial)

		if dial == 0 {
			zeroes++
		}
	}

	return dial, zeroes, clickZeroes
}

func runClickRotations(rules []DialRule) (int, int) {
	var (
		dial   = dialStartPosition
		zeroes = 0
	)

	fmt.Printf("dial at %d\n", dial)

	for _, rule := range rules {
		var (
			start  = dial
			clicks = 0
			addl   = 0

			flags []string
		)

		if rule.Rotate >= maxDial {
			flags = append(flags, fmt.Sprintf("C0/%02.2f/%d", float64(rule.Rotate)/float64(maxDial), int(rule.Rotate/maxDial)-1))
			addl = int(rule.Rotate/maxDial) - 1
		}

		switch rule.Direction {
		case DirectionL:
			dial = dial - rule.Rotate
			if dial <= 0 {
				magnitude := int(rule.Rotate/maxDial) + 1
				dial = ((maxDial * magnitude) + dial) % maxDial

				if start != 0 {
					flags = append(flags, fmt.Sprintf("C1/%d", dial))
					clicks++
				}
			}
		case DirectionR:
			dial = dial + rule.Rotate
			if dial > maxDial && start != 0 {
				flags = append(flags, fmt.Sprintf("C3/%d", dial))
				clicks++
			}

			if dial >= maxDial {
				dial = dial % maxDial
			}
		}

		if dial == 0 && clicks == 0 {
			flags = append(flags, fmt.Sprintf("C4/%d", dial))
			clicks++
		}

		flagOut := ""
		if len(flags) > 0 {
			flagOut = fmt.Sprintf(" (flags %s)", strings.Join(flags, " "))
		}
		fmt.Printf("rotated %s%d to be %d%s\n",
			rule.Direction.String(),
			rule.Rotate,
			dial,
			flagOut,
		)

		zeroes += clicks + addl
	}

	return dial, zeroes
}
