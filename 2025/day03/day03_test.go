package day03_test

import (
	"testing"

	"github.com/aaronellington/advent-of-code/2025/day03"
	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func Test_Part1_Example(t *testing.T) {
	aoc.TestCase[int]{
		FilePath: "data/part1_example.txt",
		Expected: 357,
		Func:     day03.Part1,
	}.Test(t)
}

func Test_Part1(t *testing.T) {
	aoc.TestCase[int]{
		FilePath: "data/part1.txt",
		Expected: 17193,
		Func:     day03.Part1,
	}.Test(t)
}
