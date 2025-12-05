package aoc

import (
	"testing"
)

type TestSuite struct {
	Part1 Part
	Part2 Part
}

type Part struct {
	Solution Solution
	Answer   int
	Example  int
}

func (testingSuite TestSuite) TestPart1(t *testing.T) {
	Test(t, testingSuite.Part1.Solution, "data/values.txt", testingSuite.Part1.Answer)
}

func (testingSuite TestSuite) TestPart1Example(t *testing.T) {
	Test(t, testingSuite.Part1.Solution, "data/example.txt", testingSuite.Part1.Example)
}

func (testingSuite TestSuite) TestPart2(t *testing.T) {
	Test(t, testingSuite.Part2.Solution, "data/values.txt", testingSuite.Part2.Answer)
}

func (testingSuite TestSuite) TestPart2Example(t *testing.T) {
	Test(t, testingSuite.Part2.Solution, "data/example.txt", testingSuite.Part2.Example)
}

func Test(t *testing.T, solution Solution, filePath string, expected int) {
	t.Helper()
	actual := Solve(solution, filePath)
	if actual != expected {
		t.Fatalf(
			"Got %v, expected %v",
			actual,
			expected,
		)
	}
}

type LineTest struct {
	Solution Solution
	Line     string
	Expected int
}

func (lineTest LineTest) Test(t *testing.T) {
	actual := lineTest.Solution([]string{lineTest.Line})
	if actual != lineTest.Expected {
		t.Fatalf(
			"Got %v, expected %v",
			actual,
			lineTest.Expected,
		)
	}
}
