package day03

import (
	"log"
)

func Part1(filePath string) (int, error) {
	answer := 0

	LoopOverFile(filePath, func(bank string) {
		bankVoltage := CalculateVoltage(bank, 2)

		log.Printf("Bank %s = %d", bank, bankVoltage)

		answer += bankVoltage
	})

	return answer, nil
}
