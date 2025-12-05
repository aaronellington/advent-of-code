package day03_test

import (
	"testing"

	"github.com/aaronellington/advent-of-code/2025/day03"
	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func testSuite() aoc.TestSuite {
	return aoc.TestSuite{
		Part1: aoc.Part{
			Solution: day03.Part1.Solve,
			Example:  357,
			Answer:   17193,
		},
		Part2: aoc.Part{
			Solution: day03.Part2.Solve,
			Example:  3121910778619,
			Answer:   171297349921310,
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

func Test_Line_01(t *testing.T) {
	aoc.LineTest{
		Line:     "57489376",
		Solution: day03.Solution{TargetBatteryCount: 4}.Solve,
		Expected: 9376,
	}.Test(t)
}
