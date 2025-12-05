package day05

import (
	"strconv"
	"strings"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

type State struct {
	FreshIngredientIDRanges []Range
	AvailableIngredientIDs  []int
}

type Range struct {
	Start int
	End   int
}

func linesToState(lines []string) State {
	state := State{}
	for _, line := range lines {
		parts := strings.Split(line, "-")

		firstNumber, err := strconv.Atoi(parts[0])
		aoc.PanicOnErr(err)
		if len(parts) == 2 {
			secondNumber, err := strconv.Atoi(parts[1])
			aoc.PanicOnErr(err)
			state.FreshIngredientIDRanges = append(state.FreshIngredientIDRanges, Range{
				Start: firstNumber,
				End:   secondNumber,
			})
		} else {
			state.AvailableIngredientIDs = append(state.AvailableIngredientIDs, firstNumber)
		}
	}

	return state
}
