package day01

import (
	"strconv"
	"strings"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func LoopOverFile(filePath string, v func(direction string, amount int)) {
	fileString := aoc.ReadFile(filePath)

	for _, lineString := range strings.Split(fileString, "\n") {
		if lineString == "" {
			continue
		}

		direction := string(lineString[0])

		amount, err := strconv.Atoi(lineString[1:])
		aoc.PanicOnErr(err)

		v(direction, amount)
	}
}
