package utils

import (
	"bytes"
	"fmt"
	"os"
)

func BuildInputFilepath(currentDayNumber int) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/day-%d/input.txt", wd, currentDayNumber)
}

func ReadFile(filepath string) []byte {
	content, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	return content
}

func ReadFileAndSplitLines(filepath string) [][]byte {
	content := ReadFile(filepath)

	return bytes.Split(content, []byte("\n"))
}
