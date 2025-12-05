package day04

func Part1(lines []string) int {
	grid := buildGrid(lines)

	return len(accessibleRolls(grid))
}
