package day01

import (
	"log"
	"strconv"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

type Part1 struct {
	Position int
}

func (s *Part1) SolveFile() aoc.FileSolver { return nil }

func (s *Part1) SolveLine() aoc.LineSolver {
	return func(lineIndex int, lines []string) int {
		line := lines[lineIndex]
		timesAtPositionZero := 0

		direction := string(line[0])

		amount, err := strconv.Atoi(line[1:])
		aoc.PanicOnErr(err)

		switch direction {
		case "R":
			s.Position += amount
		case "L":
			s.Position -= amount
		}

		s.Position = s.Position % 100
		if s.Position < 0 {
			s.Position = 100 - s.Position*-1
		}

		if s.Position == 0 {
			timesAtPositionZero++
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
}
