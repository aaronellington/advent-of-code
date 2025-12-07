package day06_test

import (
	"testing"

	"github.com/aaronellington/advent-of-code/2025/day06"
	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func testSuite() aoc.TestSuite {
	return aoc.TestSuite{
		Part1: aoc.Part{
			Solution: day06.Part1,
			Example:  4277556,
			Answer:   5873191732773,
		},
		Part2: aoc.Part{
			Solution: day06.Part2,
			Example:  3263827,
			Answer:   11386445308378,
		},
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
