package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func Part2(lines []string) int {
	invalidIdentifierSum := 0

	isValidID := func(id int) bool {
		s := fmt.Sprintf("%d", id)

		middle := len(s) / 2
		for charsToCheck := 1; charsToCheck <= middle; charsToCheck++ {
			part := s[0:charsToCheck]
			targetString := strings.Repeat(part, len(s)/charsToCheck)
			if targetString == s {
				return false
			}
		}

		return true
	}

	for _, line := range lines {

		parts := []int{}
		for _, s := range strings.Split(line, "-") {
			n, err := strconv.Atoi(s)
			aoc.PanicOnErr(err)

			parts = append(parts, n)
		}

		start := parts[0]
		end := parts[1]

		for i := start; i <= end; i++ {
			if isValidID(i) {
				continue
			}

			invalidIdentifierSum += i
		}
	}

	return invalidIdentifierSum
}
