package chapters

import (
	"regexp"
	"strconv"
	"strings"
)

func Chapter2(games [][]byte) int {
	gameRe := regexp.MustCompile(`Game \d+: `)

	value := 0

	for _, game := range games {
		minAmountOfCubes := map[string]int{
			"blue":  0,
			"green": 0,
			"red":   0,
		}
		game = gameRe.ReplaceAll(game, []byte(""))

		cubeStrings := strings.FieldsFunc(string(game), splitOnCommaAndSemicolon)

		for _, s := range cubeStrings {
			s = strings.Trim(s, " ")
			cubeInfo := strings.Split(s, " ")

			amount, _ := strconv.Atoi(cubeInfo[0])

			if amount > minAmountOfCubes[cubeInfo[1]] {
				minAmountOfCubes[cubeInfo[1]] = amount
			}
		}

		value += minAmountOfCubes["blue"] * minAmountOfCubes["green"] * minAmountOfCubes["red"]
	}

	return value
}
