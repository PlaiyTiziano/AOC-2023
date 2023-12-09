package chapters

import (
	"slices"
	"strconv"
	"strings"
)

func getCardValueV2(card rune) int {
	mappedCards := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 1,
		'T': 10,
	}

	value := mappedCards[card]
	if value == 0 {
		value, _ = strconv.Atoi(string(card))
	}

	return value
}

func compareHandsV2(a, b Hand) int {
	if a.value < b.value {
		return -1
	}
	if a.value > b.value {
		return 1
	}

	for idx := 0; idx < len(a.cards); idx++ {
		valA := getCardValueV2(rune(a.cards[idx]))
		valB := getCardValueV2(rune(b.cards[idx]))

		if valA < valB {
			return -1
		}
		if valA > valB {
			return 1
		}
	}

	return 0
}

func handsBinarySearchV2(a, b Hand) int {
	return compareHandsV2(a, b)
}

func replaceJokers(hand *Hand) {
	if !strings.Contains(hand.cards, "J") {
		return
	}

	cardTypes := []string{
		"A",
		"K",
		"Q",
		"T",
		"9",
		"8",
		"7",
		"6",
		"5",
		"4",
		"3",
		"2",
	}
	var prevHighestValue int

	// Very naive way of finding the best possible hand value
	for _, ct := range cardTypes {
		s := strings.Replace(hand.cards, "J", ct, -1)
		val := calculateHandValue(s)
		if val > prevHighestValue {
			prevHighestValue = val
		}
	}

	hand.value = prevHighestValue
}

func Chapter2(data []byte) int {
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

		replaceJokers(&hand)

		idx, _ := slices.BinarySearchFunc(hands, hand, handsBinarySearchV2)
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
