package aoc

import (
	"strings"
)

type Solution func(lines []string) int

func Solve(solution Solution, filePath string) int {
	split := func(r rune) bool {
		return r == '\n'
	}

	lines := strings.FieldsFunc(readFile(filePath), split)
	return solution(lines)

}
