package day04_test

import (
	"testing"

	"github.com/aaronellington/advent-of-code/2025/day04"
	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func testSuite() aoc.TestSuite {
	return aoc.TestSuite{
		SolutionPart1: day04.Part1{},
		Part1Example:  13,
		Part1:         1457,
		SolutionPart2: day04.Part2{},
		Part2Example:  43,
		Part2:         8310,
	}
}

func Test_Part1_Example(t *testing.T) {
	testSuite().TestPart1Example(t)
}

func Test_Part1(t *testing.T) {
	testSuite().TestPart1(t)
}

func Test_Part2_Example(t *testing.T) {
	testSuite().TestPart2Example(t)
}

func Test_Part2(t *testing.T) {
	testSuite().TestPart2(t)
}
