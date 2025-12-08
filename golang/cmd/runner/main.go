package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jinlimei/adventofcode/golang/library/aoc"
	yr2024day01 "github.com/jinlimei/adventofcode/golang/yr2024/day01"
	yr2024day02 "github.com/jinlimei/adventofcode/golang/yr2024/day02"
	yr2024day03 "github.com/jinlimei/adventofcode/golang/yr2024/day03"
	yr2024day04 "github.com/jinlimei/adventofcode/golang/yr2024/day04"
	yr2024day05 "github.com/jinlimei/adventofcode/golang/yr2024/day05"
	yr2024day06 "github.com/jinlimei/adventofcode/golang/yr2024/day06"
	yr2024day07 "github.com/jinlimei/adventofcode/golang/yr2024/day07"

	yr2025day01 "github.com/jinlimei/adventofcode/golang/yr2025/day01"
)

var challenges = map[int]map[string]aoc.CodeDay{
	2024: {
		"01": &yr2024day01.Day{},
		"02": &yr2024day02.Day{},
		"03": &yr2024day03.Day{},
		"04": &yr2024day04.Day{},
		"05": &yr2024day05.Day{},
		"06": &yr2024day06.Day{},
		"07": &yr2024day07.Day{},
	},
	2025: {
		"01": &yr2025day01.Day{},
	},
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run ./cmd/runner [year] [1-x] [part1|part2] [prompt|actual]")
		return
	}

	var (
		yearArg = os.Args[1]
		dayArg  = os.Args[2]
		part    = os.Args[3]
		version = os.Args[4]

		group = strings.ToLower(part + version)
	)

	yearInt, err := strconv.Atoi(yearArg)
	if err != nil {
		fmt.Println("year argument must be an integer")
		return
	}

	dayInt, err := strconv.Atoi(dayArg)
	if err != nil {
		fmt.Println("day argument must be an integer")
		return
	}

	challengeYear, ok := challenges[yearInt]
	if !ok {
		fmt.Println("unknown year")
		return
	}

	day, ok := challengeYear[fmt.Sprintf("%02d", dayInt)]
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
