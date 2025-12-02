package day2_test

import (
	"testing"

	"github.com/aaronellington/advent-of-code/2025/day2"
	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func Test_Part1_Example(t *testing.T) {
	aoc.TestCase[int]{
		FilePath: "data/part1_example.txt",
		Expected: 1227775554,
		Func:     day2.Part1,
	}.Test(t)
}

func Test_Part1(t *testing.T) {
	aoc.TestCase[int]{
		FilePath: "data/part1.txt",
		Expected: 21139440284,
		Func:     day2.Part1,
	}.Test(t)
}
