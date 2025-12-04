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
	}
}

func Test_Part1_Example(t *testing.T) {
	testSuite().TestPart1Example(t)
}

func Test_Part1(t *testing.T) {
	testSuite().TestPart1(t)
}
