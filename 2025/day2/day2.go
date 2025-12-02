package day2

import (
	"strconv"
	"strings"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func LoopOverFile(filePath string, v func(start int, end int)) {
	Split := func(r rune) bool {
		return r == '\n' || r == ','
	}

	for _, entry := range strings.FieldsFunc(aoc.ReadFile(filePath), Split) {
		if entry == "" {
			continue
		}

		parts := []int{}
		for _, s := range strings.Split(entry, "-") {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}

			parts = append(parts, n)
		}

		v(parts[0], parts[1])
	}
}
