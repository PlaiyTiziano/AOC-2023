package chapters

import (
	"slices"
)

func Chapter1(cards [][]rune) int {
	indexOfColon := slices.Index(cards[0], ':')
	indexOfPipe := slices.Index(cards[0], '|')

	points := 0

	for _, card := range cards {
		if len(card) <= 0 {
			break
		}

		winningNumbers := []string{}

		for idx := indexOfColon+2; idx < indexOfPipe-1; idx += 3 {
			winningNumbers = append(winningNumbers, string(card[idx:idx+2]))
		}

		amountOfWinningNumbers := 0

		for idx := indexOfPipe + 2; idx < len(card); idx += 3 {
			if slices.Contains(winningNumbers, string(card[idx:idx+2])) {
				amountOfWinningNumbers++
			}
		}

		if amountOfWinningNumbers > 0 {
			points += (1 << (amountOfWinningNumbers - 1))
		}
	}

	return points
}
