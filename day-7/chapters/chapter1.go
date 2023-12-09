package chapters

import (
	"slices"
	"strconv"
	"strings"
)

const (
	fiveOfAKind  = 21
	fourOfAKind  = 20
	fullHouse    = 19
	threeOfAKind = 18
	twoPair      = 17
	onePair      = 16
	highCard = 15
)

type Hand struct {
	cards string
	value int
	bet   int
}

func getCardValue(card rune) int {
	mappedCards := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
	}

	value := mappedCards[card]
	if value == 0 {
		value, _ = strconv.Atoi(string(card))
	}

	return value
}

func compareHands(a, b Hand) int {
	if a.value < b.value {
		return -1
	}
	if a.value > b.value {
		return 1
	}

	for idx := 0; idx < len(a.cards); idx++ {
		valA := getCardValue(rune(a.cards[idx]))
		valB := getCardValue(rune(b.cards[idx]))

		if valA < valB {
			return -1
		}
		if valA > valB {
			return 1
		}
	}

	return 0
}

func handsBinarySearch(a, b Hand) int {
	return compareHands(a, b)
}

func calculateHandValue(cards string) int {
	tmp := make(map[rune]int)

	for _, char := range cards {
		tmp[char]++
	}

	switch len(tmp) {
	case 1:
		return fiveOfAKind
	case 2:
		firstChar := cards[0]
		if slices.Contains([]int{1, 4}, tmp[rune(firstChar)]) {
			return fourOfAKind
		} else {
			return fullHouse
		}
	case 3:
		firstChar, secondChar := cards[0], cards[1]
		if slices.Contains([]int{tmp[rune(firstChar)], tmp[rune(secondChar)]}, 2) {
			return twoPair
		}
		return threeOfAKind
	case 4:
		return onePair
	default:
		return highCard

		// The code piece below screwed me over so hard.
		// Apparently the hand value is not determined by its highest card when no combination can be formed
		// Instead the hands get the same value "highCard" and then we loop over all chars

		// var cardValues []int
		// // Probably stupid looping over the cards again here
		// for _, card := range cards {
		// 	cardValues = append(cardValues, getCardValue(card))
		// }
		//
		// return slices.Max(cardValues)
	}
}

func Chapter1(data []byte) int {
	var hands []Hand
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			break
		}

		info := strings.Split(line, " ")

		bet, _ := strconv.Atoi(info[1])
		hand := Hand{
			cards: info[0],
			value: calculateHandValue(info[0]),
			bet:   bet,
		}

		idx, _ := slices.BinarySearchFunc(hands, hand, handsBinarySearch)
		if len(hands) == idx {
			hands = append(hands, hand)
			continue
		}
		hands = append(hands[:idx+1], hands[idx:]...)
		hands[idx] = hand
	}

	var value int
	for idx, hand := range hands {
		value += hand.bet * (idx + 1)
	}

	return value
}
