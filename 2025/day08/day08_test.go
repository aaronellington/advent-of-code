package day08_test

import (
	"testing"

	"github.com/aaronellington/advent-of-code/2025/day08"
	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func Test_Part1_Example(t *testing.T) {
	aoc.Test(
		t,
		day08.Part1{MaxConnections: 10}.Solve,
		"data/example.txt",
		40,
	)
}

func Test_Part1(t *testing.T) {
	aoc.Test(
		t,
		day08.Part1{MaxConnections: 1000}.Solve,
		"data/values.txt",
		66640,
	)
}

func Test_Part2_Example(t *testing.T) {
	aoc.Test(
		t,
		day08.Part2{MaxConnections: 10}.Solve,
		"data/example.txt",
		25272,
	)
}

func Test_Part2(t *testing.T) {
	aoc.Test(
		t,
		day08.Part2{MaxConnections: 1000}.Solve,
		"data/values.txt",
		78894156,
	)
}
