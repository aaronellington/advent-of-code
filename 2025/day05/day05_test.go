package day05_test

import (
	"testing"

	"github.com/aaronellington/advent-of-code/2025/day05"
	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func testSuite() aoc.TestSuite {
	return aoc.TestSuite{
		Part1: aoc.Part{
			Solution: day05.Part1,
			Example:  3,
			Answer:   773,
		},
		Part2: aoc.Part{
			Solution: day05.Part2,
			Example:  14,
			Answer:   332067203034711,
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
