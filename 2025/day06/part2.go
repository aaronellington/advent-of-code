package day06

import (
	"strconv"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func Part2(lines []string) int {
	grandTotal := 0

	for operationIndex, operation := range lines[len(lines)-1] {
		if operation == ' ' {
			continue
		}

		numbers := []int{}
		indexToCheck := operationIndex
		for { // Keep checking until we get to a blank column
			fullNumberString := ""
			for lineI := 0; lineI <= len(lines)-2; lineI++ {
				// Skip if the line isn't long enough
				if len(lines[lineI])-1 < indexToCheck {
					continue
				}

				numberString := string(lines[lineI][indexToCheck])
				if numberString == " " {
					continue
				}

				fullNumberString += numberString
			}

			indexToCheck++

			// break on full empty column
			if fullNumberString == "" {
				break
			}

			number, err := strconv.Atoi(fullNumberString)
			aoc.PanicOnErr(err)

			numbers = append(numbers, number)
		}

		operationResult := numbers[0]
		for i, number := range numbers {
			if i == 0 {
				continue
			}

			switch operation {
			case '+':
				operationResult = operationResult + number
			case '*':
				operationResult = operationResult * number
			}
		}

		grandTotal += operationResult
	}

	return grandTotal
}
