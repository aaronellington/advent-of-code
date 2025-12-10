package day09_test

import (
	"testing"

	"github.com/aaronellington/advent-of-code/2025/day09"
	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func Test_Part1_Example(t *testing.T) {
	aoc.Test(
		t,
		day09.Part1,
		"data/example.txt",
		50,
	)
}

func Test_Part1(t *testing.T) {
	aoc.Test(
		t,
		day09.Part1,
		"data/values.txt",
		4771532800,
	)
}

func Test_Part2_Example(t *testing.T) {
	aoc.Test(
		t,
		day09.Part2,
		"data/example.txt",
		24,
	)
}

func Test_Part2(t *testing.T) {
	aoc.Test(
		t,
		day09.Part2,
		"data/values.txt",
		1544362560,
	)
}

func Test_Part2_Case_C(t *testing.T) {
	aoc.Test(
		t,
		func(lines []string) int {
			state := day09.NewState(lines)
			state.Print()

			for y := 0; y <= state.Height; y++ {
				for x := 0; x <= state.Width; x++ {
					state.IsVectorInPolygon(day09.Vector{
						X: x,
						Y: y,
					})
				}
			}
			state.Print()
			return 0
		},
		"data/case_c.txt",
		0,
	)
}
