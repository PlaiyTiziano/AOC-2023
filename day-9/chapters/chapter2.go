package chapters

func prependPrediction(sequences [][]int) int {
	for idx := len(sequences) - 1; idx > 0; idx-- {
		prevSeq := sequences[idx-1]
		currSeq := sequences[idx]

		next := prevSeq[0] - currSeq[0]
		sequences[idx-1] = append([]int{next}, sequences[idx-1]...)
	}

	// println(sequences[0][0])
	return sequences[0][0]
}

func Chapter2(data []byte) int {
	sequences := parseSequences(data)

	for idx := range sequences {
		sequences[idx] = computeValues(sequences[idx])
	}

	// fmt.Printf("%v\n", sequences)

	var score int
	for _, s := range sequences {
		score += prependPrediction(s)
	}

	return score
}
