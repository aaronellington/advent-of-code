package day04

import (
	"github.com/aaronellington/advent-of-code/internal/aoc"
)

type Part2 struct{}

func (s Part2) SolveLine() aoc.LineSolver { return nil }

func (s Part2) SolveFile() aoc.FileSolver {
	return func(lines []string) int {
		retrievedRollCount := 0

		grid := buildGrid(lines)
		for {
			accessibleRollVectors := accessibleRolls(grid)
			if len(accessibleRollVectors) == 0 {
				break
			}

			for _, vector := range accessibleRollVectors {
				grid[vector] = false
				retrievedRollCount++
			}
		}

		return retrievedRollCount
	}
}
