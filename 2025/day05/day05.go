package day05

import (
	"slices"
	"strconv"
	"strings"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

type State struct {
	FreshIngredientIDRanges map[int]Range
	AvailableIngredientIDs  []int
	starts                  []int
}

type Range struct {
	Start int
	End   int
}

func linesToState(lines []string) State {
	state := State{
		FreshIngredientIDRanges: map[int]Range{},
	}

	for _, line := range lines {
		parts := strings.Split(line, "-")

		firstNumber, err := strconv.Atoi(parts[0])
		aoc.PanicOnErr(err)

		if len(parts) == 2 {
			secondNumber, err := strconv.Atoi(parts[1])
			aoc.PanicOnErr(err)

			existing, alreadyExisting := state.FreshIngredientIDRanges[firstNumber]
			// Skip if duplicate start, but smaller end
			if alreadyExisting && existing.End > secondNumber {
				continue
			}

			state.FreshIngredientIDRanges[firstNumber] = Range{
				Start: firstNumber,
				End:   secondNumber,
			}

			// Don't append the same number twice
			if !alreadyExisting {
				state.starts = append(state.starts, firstNumber)
			}
		} else {
			state.AvailableIngredientIDs = append(state.AvailableIngredientIDs, firstNumber)
		}
	}

	slices.Sort(state.starts)

	// Fix the ranges so they don't overlap
	previousStart := state.starts[0]
	for i, start := range state.starts {
		// skip the first start
		if i == 0 {
			continue
		}

		rangeCurrent := state.FreshIngredientIDRanges[start]
		rangePrevious := state.FreshIngredientIDRanges[previousStart]

		// Remove inclusive ranges
		if rangeCurrent.Start > rangePrevious.Start && rangeCurrent.End <= rangePrevious.End {
			delete(state.FreshIngredientIDRanges, start)
			continue
		}

		// Modify the end values so they don't overlap
		if rangeCurrent.Start <= rangePrevious.End {
			state.FreshIngredientIDRanges[previousStart] = Range{
				Start: rangePrevious.Start,
				End:   start - 1,
			}
		}

		previousStart = start
	}

	return state
}
