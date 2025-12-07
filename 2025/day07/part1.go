package day07

import (
	"strings"
)

func Part1(lines []string) int {
	previousRowLasers := map[int]bool{}
	splits := 0
	for y, line := range lines {

		// Starting location
		if y == 0 {
			previousRowLasers[strings.Index(line, "S")] = true
			continue
		}

		newLaserLocations := map[int]bool{}

		// Loop over laser locations on row above
		for previousRowLaser := range previousRowLasers {
			lineChar := string(line[previousRowLaser])

			switch lineChar {
			case "^":
				splits++

				// Place laster to right if there is room
				if previousRowLaser <= len(line) {
					newLaserLocations[previousRowLaser+1] = true
				}

				// Place laster to left if there is room
				if previousRowLaser != 0 {
					newLaserLocations[previousRowLaser-1] = true
				}

			case ".":
				// Allow laser to fall
				newLaserLocations[previousRowLaser] = true
			}
		}

		previousRowLasers = newLaserLocations
	}

	return splits
}
