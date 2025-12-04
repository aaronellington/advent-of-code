package aoc

import (
	"strings"
)

type Solution interface {
	SolveLine() LineSolver
	SolveFile() FileSolver
}
type FileSolver func(lines []string) int

type LineSolver func(lineIndex int, lines []string) int

func Solve(solution Solution, filePath string) int {
	answer := 0

	split := func(r rune) bool {
		return r == '\n' || r == ','
	}

	lines := strings.FieldsFunc(readFile(filePath), split)

	if fileSolver := solution.SolveFile(); fileSolver != nil {
		return fileSolver(lines)
	} else if lineSolver := solution.SolveLine(); lineSolver != nil {
		for lineIndex := range strings.FieldsFunc(readFile(filePath), split) {
			answer += lineSolver(lineIndex, lines)
		}
	} else {
		panic("no solver")
	}

	return answer
}
