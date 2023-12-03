package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/PlaiyTiziano/AOC-2023/utils"
)

func chapter1(lines [][]byte) int {
	regex := regexp.MustCompile(`\d`)

	value := 0

	for _, line := range lines {
		digits := regex.FindAll(line, -1)
		len_digits := len(digits)

		if len_digits == 0 {
			continue
		}

		lhs, err := strconv.Atoi(string(digits[0]))
		if err != nil {
			panic(err)
		}

		value += lhs * 10

		if len_digits == 1 {
			value += lhs
			continue
		}

		rhs, err := strconv.Atoi(string(digits[len_digits-1]))
		if err != nil {
			panic(err)
		}

		value += rhs
	}

	return value
}

func wordToNumber(word string) (int, bool) {
	number_map := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	number, exists := number_map[word]
	return number, exists
}

func parseNumber(char byte) int {
	number, err := strconv.Atoi(string(char))
	if err != nil {
		panic(err)
	}

	return number
}

func chapter2(lines [][]byte) int {
	words := "one two three four five six seven eight nine"

	value := 0

	for _, line := range lines {
		found_digits := []int{}

		word := ""
		for _, char := range line {
			if unicode.IsDigit(rune(char)) {
				found_digits = append(found_digits, parseNumber(char))
				word = ""
			}

			word += string(char)

			if len(word) < 3 {
				continue
			}

			if strings.Contains(words, word) {
				number, exists := wordToNumber(word)
				if !exists {
					continue
				}

				found_digits = append(found_digits, number)
			}

			word = word[len(word)-2:]
		}

		if len(found_digits) == 0 {
			continue
		}

		lhs := found_digits[0]
		rhs := found_digits[len(found_digits)-1]

		value += lhs*10 + rhs
	}

	return value
}

func main() {
	inputFilepath := utils.BuildInputFilepath(1)
	lines := utils.ReadFileAndSplitLines(inputFilepath)

	fmt.Println("Day 1:")

	value := chapter1(lines)
	fmt.Printf("\t- Chapter 1: %d\n", value)

	value = chapter2(lines)
	fmt.Printf("\t- Chapter 2: %d\n", value)
}
