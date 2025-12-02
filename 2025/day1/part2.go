package day1

import (
	"log"
	"math"
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

		originalPosition := position
		switch direction {
		case "R":
			position += number
		case "L":
			position -= number
			// Going negative should trigger, but only if moving past 0
			if originalPosition > 0 && position < 1 {
				timesAtPositionZero++
			}
		default:
			panic("wtf")
		}

		// Calculate the rotations
		timesAtPositionZero += int(math.Floor(math.Abs(float64(position) / 100)))

		position = position % 100
		if position < 0 {
			position = 100 - position*-1
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
