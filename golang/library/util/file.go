package util

import (
	"fmt"
	"os"
)

func ReadInputFile(year, day uint) (string, error) {
	bits, err := os.ReadFile(
		fmt.Sprintf("yr%04d/day%02d/input.txt", year, day),
	)

	if err != nil {
		return "", err
	}

	return string(bits), nil
}
