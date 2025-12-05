package day05

func Part2(lines []string) int {
	state := linesToState(lines)

	freshIngredientIDCount := 0
	for _, r := range state.FreshIngredientIDRanges {
		freshIngredientIDCount++ // Every range has one
		freshIngredientIDCount += r.End - r.Start
	}

	return freshIngredientIDCount
}
