package aoc

import (
	"testing"
)

type TestSuite struct {
	SolutionPart1 Solution
	SolutionPart2 Solution
	Part1Example  int
	Part1         int
	Part2Example  int
	Part2         int
}

func (testingSuite TestSuite) TestPart1(t *testing.T) {
	Test(t, testingSuite.SolutionPart1, "data/values.txt", testingSuite.Part1)
}

func (testingSuite TestSuite) TestPart1Example(t *testing.T) {
	Test(t, testingSuite.SolutionPart1, "data/example.txt", testingSuite.Part1Example)
}

func (testingSuite TestSuite) TestPart2(t *testing.T) {
	Test(t, testingSuite.SolutionPart2, "data/values.txt", testingSuite.Part2)
}

func (testingSuite TestSuite) TestPart2Example(t *testing.T) {
	Test(t, testingSuite.SolutionPart2, "data/example.txt", testingSuite.Part2Example)
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
	actual := lineTest.Solution.SolveLine(lineTest.Line)
	if actual != lineTest.Expected {
		t.Fatalf(
			"Got %v, expected %v",
			actual,
			lineTest.Expected,
		)
	}
}
