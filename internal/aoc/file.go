package aoc

import (
	"os"
)

func ReadFile(filePath string) string {
	fileBytes, err := os.ReadFile(filePath)
	PanicOnErr(err)

	return string(fileBytes)
}
