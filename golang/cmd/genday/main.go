package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run ./cmd/genday [year] [day]")
		return
	}

	year := os.Args[1]
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		panic(err)
	}

	if yearInt < 2000 || yearInt > 3000 {
		fmt.Println("year must be between 2000 and 3000 (lmao)")
		return
	}

	day := os.Args[2]
	dayInt, err := strconv.Atoi(day)

	if err != nil {
		panic(err)
	}

	if dayInt < 1 || dayInt > 25 {
		fmt.Println("day out of range")
		return
	}

	err = establishDay(yearInt, dayInt)
	if err != nil {
		panic(err)
	}
}
