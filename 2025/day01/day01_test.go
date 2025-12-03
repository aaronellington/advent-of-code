package day01_test

import (
	"testing"

	"github.com/aaronellington/advent-of-code/2025/day01"
	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func testSuite() aoc.TestSuite {
	return aoc.TestSuite{
		SolutionPart1: &day01.Part1{
			Position: 50,
		},
		SolutionPart2: &day01.Part2{
			Position: 50,
		},
		Part1Example: 3,
		Part1:        1074,
		Part2Example: 6,
		Part2:        6254,
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
