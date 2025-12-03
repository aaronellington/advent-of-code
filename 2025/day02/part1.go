package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

type Part1 struct{}

func (s Part1) SolveLine(line string) int {
	parts := []int{}
	for _, s := range strings.Split(line, "-") {
		n, err := strconv.Atoi(s)
		aoc.PanicOnErr(err)

		parts = append(parts, n)
	}

	start := parts[0]
	end := parts[1]

	invalidIdentifierSum := 0
	for i := start; i <= end; i++ {
		s := fmt.Sprintf("%d", i)

		// Skip odd length numbers
		if len(s)%2 != 0 {
			continue
		}

		middle := len(s) / 2
		firstHalf := s[:middle]
		secondHalf := s[middle:]

		if firstHalf != secondHalf {
			continue
		}

		invalidIdentifierSum += i
	}

	return invalidIdentifierSum
}
