package day01

import (
	"log"
	"strconv"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func Part1(lines []string) int {
	position := 50
	timesAtPositionZero := 0

	for _, line := range lines {
		direction := string(line[0])

		amount, err := strconv.Atoi(line[1:])
		aoc.PanicOnErr(err)

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
	}

	return timesAtPositionZero
}
