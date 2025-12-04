package day04

func accessibleRolls(grid map[Vector]bool) []Vector {
	answer := []Vector{}

	maxRolls := 4

	for vector, hasRoll := range grid {
		if !hasRoll {
			continue
		}

		columnIndex := vector.X
		rowIndex := vector.Y
		toCheck := []Vector{
			// Up
			{X: columnIndex, Y: rowIndex - 1},
			// Down
			{X: columnIndex, Y: rowIndex + 1},
			// Left
			{X: columnIndex - 1, Y: rowIndex},
			// Right
			{X: columnIndex + 1, Y: rowIndex},
			// Top Right
			{X: columnIndex + 1, Y: rowIndex - 1},
			// Top Left
			{X: columnIndex - 1, Y: rowIndex - 1},
			// bottom Right
			{X: columnIndex + 1, Y: rowIndex + 1},
			// bottom Left
			{X: columnIndex - 1, Y: rowIndex + 1},
		}

		// Check adjacent rolls
		adjacentRolls := 0
		for _, v := range toCheck {
			if grid[v] {
				adjacentRolls++
			}
		}

		// Skip if too many adjacent rolls
		if adjacentRolls >= maxRolls {
			continue
		}

		answer = append(answer, vector)
	}

	return answer
}

func buildGrid(lines []string) map[Vector]bool {
	grid := map[Vector]bool{}
	for rowIndex, line := range lines {
		for columnIndex, x := range line {
			grid[Vector{
				X: columnIndex,
				Y: rowIndex,
			}] = x == '@'
		}
	}

	return grid
}

type Vector struct {
	X int
	Y int
}
