package aoc

import (
	"log"
	"testing"
)

type TestCase[T comparable] struct {
	FilePath string
	Expected T
	Func     func(filePath string) (T, error)
}

func (testCase TestCase[T]) Test(t *testing.T) {
	answer, err := testCase.Func(testCase.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	if answer != testCase.Expected {
		t.Fatalf(
			"Got %v, expected %v",
			answer,
			testCase.Expected,
		)
	}
}
