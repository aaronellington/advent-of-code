package day03

import (
	"log"
)

func Part2(filePath string) (int, error) {
	answer := 0

	LoopOverFile(filePath, func(bank string) {
		bankVoltage := CalculateVoltage(bank, 12)

		log.Printf("Bank %s = %d", bank, bankVoltage)

		answer += bankVoltage
	})

	return answer, nil
}
