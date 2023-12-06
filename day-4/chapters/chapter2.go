package chapters


import (
	"slices"
)

func Chapter2(cards [][]rune) int {
	indexOfColon := slices.Index(cards[0], ':')
	indexOfPipe := slices.Index(cards[0], '|')

	copies := make(map[int]int)

	for y, card := range cards {
		if len(card) <= 0 {
			break
		}

		copies[y]++

		winningNumbers := []string{}

		for x := indexOfColon+2; x < indexOfPipe-1; x += 3 {
			winningNumbers = append(winningNumbers, string(card[x:x+2]))
		}

		amountOfWinningNumbers := 0

		for x := indexOfPipe + 2; x < len(card); x += 3 {
			if slices.Contains(winningNumbers, string(card[x:x+2])) {
				amountOfWinningNumbers++
			}
		}

		for idx := 1; idx < amountOfWinningNumbers + 1; idx++ {
			copies[y+idx] += copies[y]
		}
	}

	total := 0

	for _, v := range copies {
		total += v
	}

	return total
}
