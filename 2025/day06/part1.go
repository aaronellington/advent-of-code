package day06

import (
	"strconv"
	"strings"
)

func Part1(lines []string) int {

	values := [][]int{}
	for _, line := range lines {
		lineNumbers := []int{}
		for _, numberString := range strings.Fields(line) {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				break
			}

			lineNumbers = append(lineNumbers, number)
		}

		if len(lineNumbers) > 0 {
			values = append(values, lineNumbers)
		}
	}

	grandTotal := 0
	for i, operation := range strings.Fields(lines[len(lines)-1]) {
		columnResult := values[0][i]
		for _, row := range values[1:] {

			switch operation {
			case "+":
				columnResult = columnResult + row[i]
			case "*":
				columnResult = columnResult * row[i]
			}
		}

		grandTotal += columnResult
	}

	return grandTotal
}
