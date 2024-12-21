package day01

import (
	"os"
	"sort"
	"strconv"
)

func fileAsString(fileName string) (string, error) {
	bits, err := os.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(bits), nil
}

func absInt(a int) int {
	if a < 0 {
		return a * -1
	}

	return a
}

func distance(pairs [][]int) int {
	total := 0

	for _, pair := range pairs {
		total += absInt(pair[0] - pair[1])
	}

	return total
}

func couple(left, right []int) [][]int {
	sort.Ints(left)
	sort.Ints(right)

	pairs := make([][]int, 0)

	for idx, leftInt := range left {
		pairs = append(pairs, []int{leftInt, right[idx]})
	}

	return pairs
}

func parse(intake string) ([]int, []int) {
	var (
		leftInts  = make([]int, 0)
		rightInts = make([]int, 0)
	)

	var (
		runes  = []rune(intake)
		rLen   = len(runes)
		isLeft = true

		left  = make([]rune, 0)
		right = make([]rune, 0)

		chr rune

		leftInt  int
		rightInt int

		//err error
	)

	for pos := 0; pos < rLen; pos++ {
		chr = runes[pos]

		switch chr {
		case ' ':
			isLeft = false
			break
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if isLeft {
				left = append(left, chr)
			} else {
				right = append(right, chr)
			}
			break
		case '\n':
			if len(left) == 0 && len(right) == 0 {
				continue
			}

			leftInt, _ = strconv.Atoi(string(left))
			//fmt.Println(err)
			rightInt, _ = strconv.Atoi(string(right))
			//fmt.Println(err)

			left = make([]rune, 0)
			right = make([]rune, 0)

			leftInts = append(leftInts, leftInt)
			rightInts = append(rightInts, rightInt)

			isLeft = true
			break
		}
	}

	return leftInts, rightInts
}

func similarity(left, right []int) int {
	var (
		leftLen = len(left)
		simVal  = 0
	)

	for k := 0; k < leftLen; k++ {
		simVal += frequency(left[k], right) * left[k]
	}

	return simVal
}

func frequency(num int, list []int) int {
	freq := 0

	for k := 0; k < len(list); k++ {
		if list[k] == num {
			freq++
		}
	}

	return freq
}
