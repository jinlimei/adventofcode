package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run ./cmd/genday [day]")
		return
	}

	day := os.Args[1]
	dayInt, err := strconv.Atoi(day)

	if err != nil {
		panic(err)
	}

	if dayInt < 1 || dayInt > 25 {
		fmt.Println("day out of range")
		return
	}

	err = establishDay(dayInt)
	if err != nil {
		panic(err)
	}
}
