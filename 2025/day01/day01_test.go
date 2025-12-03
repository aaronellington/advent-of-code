package day01_test

import (
	"testing"

	"github.com/aaronellington/advent-of-code/2025/day01"
	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func Test_Part1_Example(t *testing.T) {
	aoc.TestCase[int]{
		FilePath: "data/part1_example.txt",
		Expected: 3,
		Func:     day01.Part1,
	}.Test(t)
}

func Test_Part1(t *testing.T) {
	aoc.TestCase[int]{
		FilePath: "data/part1.txt",
		Expected: 1074,
		Func:     day01.Part1,
	}.Test(t)
}

func Test_Part2_Example(t *testing.T) {
	aoc.TestCase[int]{
		FilePath: "data/part1_example.txt",
		Expected: 6,
		Func:     day01.Part2,
	}.Test(t)
}

func Test_Part2(t *testing.T) {
	aoc.TestCase[int]{
		FilePath: "data/part1.txt",
		Expected: 6254,
		Func:     day01.Part2,
	}.Test(t)
}
