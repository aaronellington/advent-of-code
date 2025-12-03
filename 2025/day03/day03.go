package day03

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func LoopOverFile(filePath string, v func(bank string)) {
	for _, entry := range strings.Split(aoc.ReadFile(filePath), "\n") {
		if entry == "" {
			continue
		}

		v(entry)
	}
}

func CalculateVoltage(bankString string, targetCount int) int {
	bank := []int{}
	// Convert the string to an array of numbers
	for _, s := range bankString {
		n, err := strconv.Atoi(string(s))
		aoc.PanicOnErr(err)
		bank = append(bank, n)
	}

	answers := fillNumbers(targetCount, 0)

	for bankIndex, bankNumber := range bank {
		for answerIndexIndex, answerIndex := range answers {
			// If the answer index is already past this point, move on
			if bankIndex <= answerIndex {
				continue
			}

			// Move on if we we can't select this bank for this answer location because
			// we wouldn't have enough remaining options to fill out the rest of the selection
			enoughNumbersLeft := len(bank)-bankIndex >= targetCount-answerIndexIndex
			if !enoughNumbersLeft {
				continue
			}

			// If the number is smaller than the current bank number, move on
			if bankNumber <= bank[answerIndex] {
				continue
			}

			// Select this location and auto fill the rest of the array
			x := answers[0:answerIndexIndex]
			y := fillNumbers(targetCount-answerIndexIndex, bankIndex)
			answers = append(x, y...)
		}
	}

	bankVoltageString := ""
	for _, index := range answers {
		bankVoltageString += fmt.Sprintf("%d", bank[index])
	}

	bankVoltage, err := strconv.Atoi(bankVoltageString)
	aoc.PanicOnErr(err)

	return bankVoltage

}

func fillNumbers(n int, startingPoint int) []int {
	numbers := make([]int, n)
	for i := range numbers {
		numbers[i] = i + startingPoint
	}

	return numbers
}
