package utils

import (
	"bytes"
	"fmt"
	"os"
	"strings"
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

func ReadFileSplitLinesAndChars(filepath string) [][]rune {
	content := ReadFile(filepath)

	chars := [][]rune{}
	for _, line := range strings.Split(string(content), "\n") {
		chars = append(chars, []rune(line))
	}

	return chars
}
