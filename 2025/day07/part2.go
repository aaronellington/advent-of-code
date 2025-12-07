package day07

import (
	"strings"
)

func Part2(lines []string) int {
	previousLasers := map[int]int{}
	for y, line := range lines {
		if y == 0 {
			previousLasers[strings.Index(line, "S")] = 1
			continue
		}

		newLaserLocations := previousLasers

		for previousLaserIndex, previousLaserCount := range previousLasers {
			s := string(line[previousLaserIndex])
			switch s {
			case "^":
				if previousLaserIndex <= len(line) {
					newLaserLocations[previousLaserIndex+1] += previousLaserCount
				}

				if previousLaserIndex != 0 {
					newLaserLocations[previousLaserIndex-1] += previousLaserCount

				}
				delete(newLaserLocations, previousLaserIndex)
			case ".":
				newLaserLocations[previousLaserIndex] = previousLaserCount

			}
		}

		previousLasers = newLaserLocations
	}

	sum := 0
	for _, count := range previousLasers {
		sum += count
	}

	return sum
}
