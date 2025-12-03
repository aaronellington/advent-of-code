package day03

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func Part1(filePath string) (int, error) {
	answer := 0

	LoopOverFile(filePath, func(numbers []int) {
		largestIndex := 0
		secondLargestIndex := 1
		for index, newNumber := range numbers {
			isLastNumber := len(numbers)-1 == index

			if !isLastNumber && newNumber > numbers[largestIndex] {
				largestIndex = index
				secondLargestIndex = index + 1
				continue
			}

			if secondLargestIndex < index && newNumber > numbers[secondLargestIndex] {
				secondLargestIndex = index
				continue
			}
		}

		bankVoltage, err := strconv.Atoi(fmt.Sprintf("%d%d", numbers[largestIndex], numbers[secondLargestIndex]))
		aoc.PanicOnErr(err)

		log.Printf("Bank %d", bankVoltage)

		answer += bankVoltage
	})

	return answer, nil
}
