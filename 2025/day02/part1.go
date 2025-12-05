package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func Part1(lines []string) int {
	invalidIdentifierSum := 0

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

	}

	return invalidIdentifierSum
}
