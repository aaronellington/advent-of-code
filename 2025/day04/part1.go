package day04

import (
	"github.com/aaronellington/advent-of-code/internal/aoc"
)

type Part1 struct{}

func (s Part1) SolveLine() aoc.LineSolver { return nil }

func (s Part1) SolveFile() aoc.FileSolver {
	return func(lines []string) int {
		grid := buildGrid(lines)

		return len(accessibleRolls(grid))
	}
}
