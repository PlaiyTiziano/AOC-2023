package chapters

import (
	"regexp"
)

func gcd (a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a % b)
}

func Chapter2(lines [][]byte) int {
	regex := regexp.MustCompile(`\w+`)

	var startingNodes []string
	nodes := make(map[string][]string)
	for _, line := range lines[2 : len(lines)-1] {
		words := regex.FindAll(line, -1)

		if string(words[0][2]) == "A" {
			startingNodes = append(startingNodes, string(words[0]))
		}

		nodes[string(words[0])] = []string{
			string(words[1]),
			string(words[2]),
		}

	}

	tmp := make(map[string]int)
	var steps int
L:
	for {
		for _, instruction := range string(lines[0]) {
			direction := 0
			if instruction == 'R' {
				direction = 1
			}

			steps++

			for idx, location := range startingNodes {
				startingNodes[idx] = nodes[location][direction]

				if string(startingNodes[idx][2]) == "Z" {
					tmp[startingNodes[idx]] = steps
				}

				if len(tmp) == len(startingNodes) {
					break L
				}
			}
		}
	}

	// Would've never thought of the LCM approach without reading about it on
	// reddit.
	val := 1
	for _, v := range tmp {
		val = (v * val) / (gcd(v, val))
	}

	return val
}
