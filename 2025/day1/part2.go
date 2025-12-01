package day1

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Part2(filePath string) (int, error) {
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
		number, err := strconv.Atoi(lineString[1:])
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
		if position < 0 {
			position = 100 - position*-1
		}

		if position == 0 {
			timesAtPositionZero++
		}

		log.Printf(
			"The dial is rotated %s to point at %d; during this rotation, it points at 0 %d times",
			lineString,
			position,
			timesAtPositionZero,
		)
	}

	return timesAtPositionZero, nil
}
