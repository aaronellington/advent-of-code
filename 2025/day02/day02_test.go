package day02_test

import (
	"testing"

	"github.com/aaronellington/advent-of-code/2025/day02"
	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func testSuite() aoc.TestSuite {
	return aoc.TestSuite{
		Part1: aoc.Part{
			Solution: day02.Part1,
			Example:  1227775554,
			Answer:   21139440284,
		},
		Part2: aoc.Part{
			Solution: day02.Part2,
			Example:  4174379265,
			Answer:   38731915928,
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
