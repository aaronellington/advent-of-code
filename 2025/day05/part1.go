package day05

func Part1(lines []string) int {
	freshIngredientCount := 0

	state := linesToState(lines)

	for _, ingredientID := range state.AvailableIngredientIDs {
		isFresh := false
		for _, r := range state.FreshIngredientIDRanges {
			if ingredientID >= r.Start && ingredientID <= r.End {
				isFresh = true
				break
			}
		}

		if isFresh {
			freshIngredientCount++
		}
	}

	return freshIngredientCount
}
