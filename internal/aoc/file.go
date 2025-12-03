package aoc

import (
	"os"
)

func readFile(filePath string) string {
	fileBytes, err := os.ReadFile(filePath)
	PanicOnErr(err)

	return string(fileBytes)
}
