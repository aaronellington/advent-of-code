package aoc

import (
	"strings"
)

type Solution interface {
	SolveLine(line string, lines []string) int
}

func Solve(solution Solution, filePath string) int {
	answer := 0

	split := func(r rune) bool {
		return r == '\n' || r == ','
	}

	lines := strings.FieldsFunc(readFile(filePath), split)
	for _, line := range strings.FieldsFunc(readFile(filePath), split) {
		answer += solution.SolveLine(line, lines)
	}

	return answer
}
