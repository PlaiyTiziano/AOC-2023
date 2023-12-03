package chapters

import (
	"regexp"
	"strconv"
	"strings"
)

func splitOnCommaAndSemicolon(r rune) bool {
	return r == ',' || r == ';'
}

func checkCubesValidity(cubeInfo []string, cubeLimits map[string]int) bool {
	amount, err := strconv.Atoi(cubeInfo[0])
	if err != nil {
		panic(err)
	}

	return amount <= cubeLimits[cubeInfo[1]]
}

func Chapter1(games [][]byte) int {
	cubeLimits := map[string]int{
		"blue":  14,
		"green": 13,
		"red":   12,
	}
	gameRe := regexp.MustCompile(`Game \d+: `)

	value := 0

	for idx, game := range games {
		game = gameRe.ReplaceAll(game, []byte(""))

		cubeStrings := strings.FieldsFunc(string(game), splitOnCommaAndSemicolon)

		for cubeIdx, s := range cubeStrings {
			s = strings.Trim(s, " ")
			cubeInfo := strings.Split(s, " ")

			if !checkCubesValidity(cubeInfo, cubeLimits) {
				break
			}

			if cubeIdx == len(cubeStrings)-1 {
				value += idx + 1
			}
		}
	}

	return value
}
