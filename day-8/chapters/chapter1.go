package chapters

import (
	"regexp"
	"strings"
)

func traverse(nodes map[string][]string, instructions []string, location string) (string, int) {
	var steps int

	for _, instruction := range instructions {
		if instruction == "L" {
			location = nodes[location][0]
		} else {
			location = nodes[location][1]
		}

		steps++
		if location == "ZZZ" {
			break
		}
	}

	return location, steps
}

func Chapter1(lines [][]byte) int {
	regex := regexp.MustCompile(`\w+`)

	nodes := make(map[string][]string)
	for _, line := range lines[2 : len(lines)-1] {
		words := regex.FindAll(line, -1)
		nodes[string(words[0])] = []string{
			string(words[1]),
			string(words[2]),
		}

	}

	instructions := strings.Split(string(lines[0]), "")

	var location string = "AAA"
	var steps int
	for location != "ZZZ" {
		newLocation, stepsTaken := traverse(nodes, instructions, location)

		location = newLocation
		steps += stepsTaken
	}

	return steps
}
