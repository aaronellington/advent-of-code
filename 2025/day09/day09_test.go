package day09_test

import (
	"testing"

	"github.com/aaronellington/advent-of-code/2025/day09"
	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func Test_Part1_Example(t *testing.T) {
	aoc.Test(
		t,
		day09.Part1,
		"data/example.txt",
		50,
	)
}

func Test_Part1(t *testing.T) {
	aoc.Test(
		t,
		day09.Part1,
		"data/values.txt",
		4771532800,
	)
}
