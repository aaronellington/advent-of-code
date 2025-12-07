package day07_test

import (
	"testing"

	"github.com/aaronellington/advent-of-code/2025/day07"
	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func testSuite() aoc.TestSuite {
	return aoc.TestSuite{
		Part1: aoc.Part{
			Solution: day07.Part1,
			Example:  21,
			Answer:   1633,
		},
		Part2: aoc.Part{
			Solution: day07.Part2,
			Example:  40,
			Answer:   34339203133559,
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
