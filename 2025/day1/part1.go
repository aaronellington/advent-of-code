package day1

import (
	"os"
	"strconv"
	"strings"
)

func Part1(filePath string) (int, error) {
	timesAtPositionZero := 0
	position := 50

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	for _, lineString := range strings.Split(string(fileBytes), "\n") {
		if lineString == "" {
			continue
		}
		direction := string(lineString[0])
		number, err := strconv.Atoi(lineString[1:]) // Converts "123" to an integer
		if err != nil {
			return 0, err
		}

		switch direction {
		case "R":
			position += number
		case "L":
			position -= number
		default:
			panic("wtf")
		}

		position = position % 100
		if position == 0 {
			timesAtPositionZero++
		}
	}

	return timesAtPositionZero, nil
}
