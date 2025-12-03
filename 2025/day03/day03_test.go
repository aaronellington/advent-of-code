package day03_test

import (
	"testing"

	"github.com/aaronellington/advent-of-code/2025/day03"
	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func testSuite() aoc.TestSuite {
	return aoc.TestSuite{
		SolutionPart1: day03.Part1,
		SolutionPart2: day03.Part2,
		Part1Example:  357,
		Part1:         17193,
		Part2Example:  3121910778619,
		Part2:         171297349921310,
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

func Test_Line_01(t *testing.T) {
	aoc.LineTest{
		Line:     "57489376",
		Solution: day03.Solution{TargetBatteryCount: 4},
		Expected: 9376,
	}.Test(t)
}
