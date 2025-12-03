package day01

import (
	"log"
)

func Part1(filePath string) (int, error) {
	timesAtPositionZero := 0
	position := 50

	LoopOverFile(filePath, func(direction string, amount int) {
		switch direction {
		case "R":
			position += amount
		case "L":
			position -= amount
		}

		position = position % 100
		if position < 0 {
			position = 100 - position*-1
		}

		if position == 0 {
			timesAtPositionZero++
		}

		log.Printf(
			"The dial is rotated %s%d to point at %d; during this rotation, it points at 0 %d times",
			direction,
			amount,
			position,
			timesAtPositionZero,
		)
	})

	return timesAtPositionZero, nil
}
