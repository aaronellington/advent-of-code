package day2

import (
	"fmt"
)

func Part1(fileName string) (int, error) {
	invalidIdentifierSum := 0

	LoopOverFile(fileName, func(start int, end int) {
		for i := start; i <= end; i++ {
			s := fmt.Sprintf("%d", i)

			// Skip odd length numbers
			if len(s)%2 != 0 {
				continue
			}

			middle := len(s) / 2
			firstHalf := s[:middle]
			secondHalf := s[middle:]

			if firstHalf != secondHalf {
				continue
			}

			invalidIdentifierSum += i
		}
	})

	return invalidIdentifierSum, nil
}
