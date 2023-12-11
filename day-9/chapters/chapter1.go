package chapters

import (
	"strconv"
	"strings"
)

func parseSequences(data []byte) [][][]int {
	var sequences [][][]int

	for _, sequence := range strings.Split(string(data), "\n") {
		if len(sequence) == 0 {
			break
		}

		var tmp []int
		for _, s := range strings.Split(sequence, " ") {
			num, _ := strconv.Atoi(string(s))
			tmp = append(tmp, num)
		}
		sequences = append(sequences, [][]int{tmp})
	}

	return sequences
}

func computeValues(sequences [][]int) [][]int {
	var values []int

	containsOnlyZeros := true

	sequence := sequences[len(sequences)-1]
	for j := 1; j < len(sequence); j++ {
		diff := sequence[j] - sequence[j-1]
		if diff != 0 {
			containsOnlyZeros = false
		}

		values = append(values, diff)
	}

	sequences = append(sequences, values)

	if !containsOnlyZeros {
		return computeValues(sequences)
	}

	return sequences
}

func appendPrediction(sequences [][]int) int {
	for idx := len(sequences) - 1; idx > 0; idx-- {
		prevSeq := sequences[idx-1]
		currSeq := sequences[idx]

		next := prevSeq[len(prevSeq)-1] + currSeq[len(currSeq)-1]
		sequences[idx-1] = append(sequences[idx-1], next)
	}

	return sequences[0][len(sequences[0])-1]
}

func Chapter1(data []byte) int {
	sequences := parseSequences(data)

	for idx := range sequences {
		sequences[idx] = computeValues(sequences[idx])
	}

	var score int
	for _, s := range sequences {
		score += appendPrediction(s)
	}

	return score
}
