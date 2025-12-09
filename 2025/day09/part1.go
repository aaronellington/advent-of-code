package day09

import (
	"strconv"
	"strings"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func Part1(lines []string) int {
	vectors := State(lines, nil)

	largestArea := 0

	for _, v1 := range vectors {
		for _, v2 := range vectors {
			if area := Area(v1, v2); area > largestArea {
				largestArea = area
			}
		}
	}

	return largestArea
}

type Vector struct {
	X int
	Y int
}

func Area(v1 Vector, v2 Vector) int {
	return (v2.X - v1.X + 1) * (v2.Y - v1.Y + 1)
}

func State(lines []string, f func(Vector)) []Vector {
	v := []Vector{}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		vector := Vector{}
		for i, part := range parts {
			number, err := strconv.Atoi(part)
			aoc.PanicOnErr(err)
			if i == 0 {
				vector.X = number
			} else {
				vector.Y = number
			}
		}
		if f != nil {
			f(vector)
		}
		v = append(v, vector)
	}

	return v
}
