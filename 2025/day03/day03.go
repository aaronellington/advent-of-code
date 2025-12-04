package day03

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

type Solution struct {
	TargetBatteryCount int
}

func (s Solution) SolveFile() aoc.FileSolver {
	return nil
}

func (s Solution) SolveLine() aoc.LineSolver {
	return func(lineIndex int, lines []string) int {
		line := lines[lineIndex]

		// Convert the string to an array of numbers
		bank := parseBank(line)

		// Default to selecting the first X number of indexes
		selectedIndexes := fillNumbers(s.TargetBatteryCount, 0)

		for bankIndex, bankNumber := range bank {
			for answerIndexIndex, answerIndex := range selectedIndexes {
				// If the answer index is already past this point, move on
				if bankIndex <= answerIndex {
					continue
				}

				// Move on if we we can't select this bank for this answer location because
				// we wouldn't have enough remaining options to fill out the rest of the selection
				enoughNumbersLeft := len(bank)-bankIndex >= s.TargetBatteryCount-answerIndexIndex
				if !enoughNumbersLeft {
					continue
				}

				// If the number is smaller than the current bank number, move on
				if bankNumber <= bank[answerIndex] {
					continue
				}

				// Select this location and auto fill the rest of the sequential indexes
				selectedIndexes = append(selectedIndexes[0:answerIndexIndex], fillNumbers(s.TargetBatteryCount-answerIndexIndex, bankIndex)...)
			}
		}

		bankVoltage := toVoltage(bank, selectedIndexes)
		log.Printf("Bank %s = %d", line, bankVoltage)

		return bankVoltage
	}
}

func fillNumbers(n int, startingPoint int) []int {
	numbers := make([]int, n)
	for i := range numbers {
		numbers[i] = i + startingPoint
	}

	return numbers
}

func parseBank(line string) []int {
	bank := []int{}
	for _, s := range line {
		n, err := strconv.Atoi(string(s))
		aoc.PanicOnErr(err)
		bank = append(bank, n)
	}

	return bank
}

func toVoltage(bank []int, answer []int) int {
	bankVoltageString := ""
	for _, index := range answer {
		bankVoltageString += fmt.Sprintf("%d", bank[index])
	}

	bankVoltage, err := strconv.Atoi(bankVoltageString)
	aoc.PanicOnErr(err)

	return bankVoltage
}
