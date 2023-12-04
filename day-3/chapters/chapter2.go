package chapters

import (
	"fmt"
	"unicode"
)

func Chapter2(lines [][]rune) int {
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

	ratio := 0

	for y := range lines {
		for x, char := range lines[y] {
			if char == '*' {
				lt := numberLocations[fmt.Sprintf("%d.%d", y-1, x-1)]
				t := numberLocations[fmt.Sprintf("%d.%d", y-1, x)]
				rt := numberLocations[fmt.Sprintf("%d.%d", y-1, x+1)]
				l := numberLocations[fmt.Sprintf("%d.%d", y, x-1)]
				r := numberLocations[fmt.Sprintf("%d.%d", y, x+1)]
				lb := numberLocations[fmt.Sprintf("%d.%d", y+1, x-1)]
				b := numberLocations[fmt.Sprintf("%d.%d", y+1, x)]
				rb := numberLocations[fmt.Sprintf("%d.%d", y+1, x+1)]

				values := []int{}

				if t != 0 {
					values = append(values, t)
				} else {
					if lt != 0 {
						values = append(values, lt)
					}
					if rt != 0 {
						values = append(values, rt)
					}
				}

				if l != 0 {
					values = append(values, l)
				}

				if r != 0 {
					values = append(values, r)
				}

				if b != 0 {
					values = append(values, b)
				} else {
					if lb != 0 {
						values = append(values, lb)
					}
					if rb != 0 {
						values = append(values, rb)
					}
				}

				if len(values) == 2 {
					ratio += values[0] * values[1]
				}
			}
		}
	}

	return ratio
}
