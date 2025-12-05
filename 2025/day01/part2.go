package day01

import (
	"log"
	"math"
	"strconv"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func Part2(lines []string) int {
	position := 50
	timesAtPositionZero := 0

	for _, line := range lines {
		direction := string(line[0])

		amount, err := strconv.Atoi(line[1:])
		aoc.PanicOnErr(err)

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
	}

	return timesAtPositionZero
}
