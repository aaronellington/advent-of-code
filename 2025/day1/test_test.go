package day1_test

import (
	"log"
	"testing"

	"github.com/aaronellington/advent-of-code/2025/day1"
)

func TestDay1_Part1_Example(t *testing.T) {
	tester(t, day1.Part1, "data/part1_example.txt", 3)
}

func TestDay1_Part1(t *testing.T) {
	tester(t, day1.Part1, "data/part1.txt", 1074)
}

func tester(t *testing.T, v func(filePath string) (int, error), filePath string, expected int) {
	answer, err := v(filePath)

	if err != nil {
		log.Fatal(err)
	}

	if answer != expected {
		t.Fatalf("Answer was not %d, was %d", expected, answer)
	}
}
