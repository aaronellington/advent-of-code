package day02_test

import (
	"testing"

	"github.com/aaronellington/advent-of-code/2025/day02"
	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func Test_Part1_Example(t *testing.T) {
	aoc.TestCase[int]{
		FilePath: "data/part1_example.txt",
		Expected: 1227775554,
		Func:     day02.Part1,
	}.Test(t)
}

func Test_Part1(t *testing.T) {
	aoc.TestCase[int]{
		FilePath: "data/part1.txt",
		Expected: 21139440284,
		Func:     day02.Part1,
	}.Test(t)
}

func Test_Part2_Example(t *testing.T) {
	aoc.TestCase[int]{
		FilePath: "data/part1_example.txt",
		Expected: 4174379265,
		Func:     day02.Part2,
	}.Test(t)
}

func Test_Part2(t *testing.T) {
	aoc.TestCase[int]{
		FilePath: "data/part1.txt",
		Expected: 38731915928,
		Func:     day02.Part2,
	}.Test(t)
}
