package day02

import (
	"fmt"
	"strings"
)

func Part2(fileName string) (int, error) {
	invalidIdentifierSum := 0

	isValidID := func(id int) bool {
		s := fmt.Sprintf("%d", id)

		middle := len(s) / 2
		for charsToCheck := 1; charsToCheck <= middle; charsToCheck++ {
			part := s[0:charsToCheck]
			targetString := strings.Repeat(part, len(s)/charsToCheck)
			if targetString == s {
				return false
			}
		}

		return true
	}

	LoopOverFile(fileName, func(start int, end int) {
		for i := start; i <= end; i++ {
			if isValidID(i) {
				continue
			}

			invalidIdentifierSum += i
		}
	})

	return invalidIdentifierSum, nil
}
