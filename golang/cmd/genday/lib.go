package main

import (
	"fmt"
	"os"
	"strings"
)

func dirExists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func makeDir(path string) error {
	return os.Mkdir(path, 0755)
}

func makeFile(path, contents string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write([]byte(contents))
	if err != nil {
		return err
	}

	return nil
}

func establishDay(day int) error {
	dayDir := fmt.Sprintf("yr2024/day%02d", day)

	if dirExists(dayDir) {
		return fmt.Errorf("%s already exists", dayDir)
	}

	err := makeDir(dayDir)
	if err != nil {
		return err
	}

	vars := map[string]string{
		"package": fmt.Sprintf("day%02d", day),
	}

	err = makeFile(dayDir+"/day.go", translate(dayGo, vars))
	if err != nil {
		return err
	}

	err = makeFile(dayDir+"/lib.go", translate(libGo, vars))
	if err != nil {
		return err
	}

	vars["day"] = "1"
	err = makeFile(dayDir+"/part1.go", translate(partGo, vars))
	if err != nil {
		return err
	}

	vars["day"] = "2"
	err = makeFile(dayDir+"/part2.go", translate(partGo, vars))
	if err != nil {
		return err
	}

	return nil
}

func translate(s string, vars map[string]string) string {
	for key, value := range vars {
		s = strings.ReplaceAll(s, "%%"+key+"%%", value)
	}

	return s
}
