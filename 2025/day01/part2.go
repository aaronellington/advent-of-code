package day01

import (
	"log"
	"math"
	"strconv"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

type Part2 struct {
	Position int
}

func (s *Part2) SolveLine(line string) int {
	timesAtPositionZero := 0

	direction := string(line[0])

	amount, err := strconv.Atoi(line[1:])
	aoc.PanicOnErr(err)

	originalPosition := s.Position
	switch direction {
	case "R":
		s.Position += amount
	case "L":
		s.Position -= amount
		// Going negative should trigger, but only if moving past 0
		if originalPosition > 0 && s.Position < 1 {
			timesAtPositionZero++
		}
	}

	// Calculate the rotations
	timesAtPositionZero += int(math.Floor(math.Abs(float64(s.Position) / 100)))

	s.Position = s.Position % 100
	if s.Position < 0 {
		s.Position = 100 - s.Position*-1
	}

	log.Printf(
		"The dial is rotated %s%d to point at %d; during this rotation, it points at 0 %d times",
		direction,
		amount,
		s.Position,
		timesAtPositionZero,
	)

	return timesAtPositionZero
}
