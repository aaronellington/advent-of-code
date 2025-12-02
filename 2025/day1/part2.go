package day1

import (
	"log"
	"math"
)

func Part2(filePath string) (int, error) {
	timesAtPositionZero := 0
	position := 50

	LoopOverFile(filePath, func(direction string, amount int) {
		originalPosition := position
		switch direction {
		case "R":
			position += amount
		case "L":
			position -= amount
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
			"The dial is rotated %s%d to point at %d; during this rotation, it points at 0 %d times",
			direction,
			amount,
			position,
			timesAtPositionZero,
		)
	})

	return timesAtPositionZero, nil
}
