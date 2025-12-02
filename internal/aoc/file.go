package aoc

import "os"

func ReadFile(filePath string) string {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(fileBytes)
}
