package day04

func Part2(lines []string) int {
	retrievedRollCount := 0

	grid := buildGrid(lines)
	for {
		accessibleRollVectors := accessibleRolls(grid)
		if len(accessibleRollVectors) == 0 {
			break
		}

		for _, vector := range accessibleRollVectors {
			grid[vector] = false
			retrievedRollCount++
		}
	}

	return retrievedRollCount
}
