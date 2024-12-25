package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jinlimei/adventofcode/golang/library/aoc"
	"github.com/jinlimei/adventofcode/golang/yr2024/day01"
	"github.com/jinlimei/adventofcode/golang/yr2024/day02"
	"github.com/jinlimei/adventofcode/golang/yr2024/day03"
	"github.com/jinlimei/adventofcode/golang/yr2024/day04"
	"github.com/jinlimei/adventofcode/golang/yr2024/day05"
	"github.com/jinlimei/adventofcode/golang/yr2024/day06"
)

var days = map[string]aoc.CodeDay{
	"01": &day01.Day{},
	"02": &day02.Day{},
	"03": &day03.Day{},
	"04": &day04.Day{},
	"05": &day05.Day{},
	"06": &day06.Day{},
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run ./cmd/runner [1-x] [part1|part2] [prompt|actual]")
		return
	}

	var (
		dayArg  = os.Args[1]
		part    = os.Args[2]
		version = os.Args[3]

		group = strings.ToLower(part + version)
	)

	dayInt, err := strconv.Atoi(dayArg)
	if err != nil {
		fmt.Println("day argument must be an integer")
		return
	}

	day, ok := days[fmt.Sprintf("%02d", dayInt)]
	if !ok {
		fmt.Println("day not found")
		return
	}

	fmt.Println("Executing", group)

	switch group {
	case "part1prompt":
		day.Part1Prompt()
	case "part1actual":
		day.Part1Actual()
	case "part2prompt":
		day.Part2Prompt()
	case "part2actual":
		day.Part2Actual()
	default:
		fmt.Println("Unknown group", group)
	}

	fmt.Println("Completed")
}
