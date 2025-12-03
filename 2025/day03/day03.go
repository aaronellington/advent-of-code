package day03

import (
	"strconv"
	"strings"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func LoopOverFile(filePath string, v func(numbers []int)) {
	for _, entry := range strings.Split(aoc.ReadFile(filePath), "\n") {
		if entry == "" {
			continue
		}

		parts := []int{}
		for _, s := range entry {
			n, err := strconv.Atoi(string(s))
			aoc.PanicOnErr(err)

			parts = append(parts, n)
		}

		v(parts)
	}
}
