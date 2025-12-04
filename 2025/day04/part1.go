package day04

import (
	"log"
	"strings"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

type Part1 struct{}

func (s Part1) SolveFile() aoc.FileSolver {
	return func(lines []string) int {
		accessibleRolls := 0
		maxRolls := 4
		grid := buildGrid(lines)

		for rowIndex, row := range lines {
			outputLine := strings.Split(row, "")

			line := lines[rowIndex]
			for columnIndex, v := range line {
				// Skip if not a roll
				if v != '@' {
					continue
				}

				toCheck := []Vector{
					// Up
					{X: columnIndex, Y: rowIndex - 1},
					// Down
					{X: columnIndex, Y: rowIndex + 1},
					// Left
					{X: columnIndex - 1, Y: rowIndex},
					// Right
					{X: columnIndex + 1, Y: rowIndex},
					// Top Right
					{X: columnIndex + 1, Y: rowIndex - 1},
					// Top Left
					{X: columnIndex - 1, Y: rowIndex - 1},
					// bottom Right
					{X: columnIndex + 1, Y: rowIndex + 1},
					// bottom Left
					{X: columnIndex - 1, Y: rowIndex + 1},
				}

				// Check adjacent rolls
				adjacentRolls := 0
				for _, v := range toCheck {
					if grid[v] {
						outputLine[columnIndex] = "X"
						adjacentRolls++
					}
				}

				// Skip if too many adjacent rolls
				if adjacentRolls >= maxRolls {
					continue
				}

				accessibleRolls++
			}

			log.Println(strings.Join(outputLine, ""))
		}

		return accessibleRolls
	}
}

func (s Part1) SolveLine() aoc.LineSolver {
	return nil
}

type Vector struct {
	X int
	Y int
}

func buildGrid(lines []string) map[Vector]bool {
	grid := map[Vector]bool{}
	for rowIndex, line := range lines {
		for columnIndex, x := range line {
			grid[Vector{
				X: columnIndex,
				Y: rowIndex,
			}] = x == '@'
		}
	}

	return grid
}
