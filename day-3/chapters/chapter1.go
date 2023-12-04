package chapters

import (
	"fmt"
	"slices"
	"strconv"
	"unicode"
)

func parseAndSetNumberLocations(numberLocations map[string]int, lastFoundNumbers []rune, y, x int) {
	number, _ := strconv.Atoi(string(lastFoundNumbers))

	for idx := x - len(lastFoundNumbers); idx < x; idx++ {
		numberLocations[fmt.Sprintf("%d.%d", y, idx)] = number
	}
}

func Chapter1(lines [][]rune) int {
	numberLocations := make(map[string]int)

	for y, line := range lines {
		lastFoundNumbers := []rune{}

		for x := 0; x < len(line); x++ {
			char := line[x]

			if !unicode.IsDigit(char) {
				if len(lastFoundNumbers) > 0 {
					parseAndSetNumberLocations(numberLocations, lastFoundNumbers, y, x)

					lastFoundNumbers = nil
				}

				continue
			}

			lastFoundNumbers = append(lastFoundNumbers, char)
		}

		parseAndSetNumberLocations(numberLocations, lastFoundNumbers, y, len(line))
	}

	symbolCodes := []rune("!#$%&*+,/=@-_")

	value := 0

	for y := range lines {
		for x, char := range lines[y] {
			if slices.Contains(symbolCodes, char) {
				lt := numberLocations[fmt.Sprintf("%d.%d", y-1, x-1)]
				t := numberLocations[fmt.Sprintf("%d.%d", y-1, x)]
				rt := numberLocations[fmt.Sprintf("%d.%d", y-1, x+1)]
				l := numberLocations[fmt.Sprintf("%d.%d", y, x-1)]
				r := numberLocations[fmt.Sprintf("%d.%d", y, x+1)]
				lb := numberLocations[fmt.Sprintf("%d.%d", y+1, x-1)]
				b := numberLocations[fmt.Sprintf("%d.%d", y+1, x)]
				rb := numberLocations[fmt.Sprintf("%d.%d", y+1, x+1)]

				// println(y, x, string(char))
				// println(lt, t, rt)
				// println(l, r)
				// println(lb, b, rb)

				value += t + r + b + l

				if t == 0 {
					// println("adding lt and rt")
					value += lt + rt
				}

				if b == 0 {
					// println("adding lb and rb")
					value += lb + rb
				}
			}
		}
	}

	// fmt.Printf("%q\n", numberLocations)

	return value
}

// for x, char := range line {
// 	for unicode.IsDigit(char) {
//
// 	}
// 	// if slices.Contains(symbolCodes, char) {
// 	// 	// Perform a adjacent checks
// 	// 	if y != 0 {
// 	// 		if x != 0 {
// 	//
// 	// 		}
// 	// 	}
// 	// 	// if y != 0 && unicode.IsDigit(lines[y-1][x]) {
// 	// 	// 	indexesOfAdjacingNumbers = append(indexesOfAdjacingNumbers, []int{y,x})
// 	// 	// }
// 	// }
// }
