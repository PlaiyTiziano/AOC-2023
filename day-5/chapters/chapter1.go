package chapters

import (
	"slices"
	"strconv"
	"strings"
)

type Ranges struct {
	dest    int
	destMax int
	src     int
	srcMax  int
}

func Chapter1(content []byte) int {
	instructions := strings.Split(string(content), "\n\n")

	seeds := []int{}
	for _, seed := range strings.Split(strings.Replace(instructions[0], "seeds: ", "", -1), " ") {
		number, _ := strconv.Atoi(seed)
		seeds = append(seeds, number)
	}

	for _, instruction := range instructions {
		ranges := []Ranges{}

		for idx, line := range strings.Split(instruction, "\n") {
			if idx == 0 || len(line) == 0 {
				continue
			}

			info := []int{}

			for _, s := range strings.Split(line, " ") {
				number, _ := strconv.Atoi(s)
				info = append(info, number)
			}

			dest, src, length := info[0], info[1], info[2]
			destMax := dest + length - 1
			srcMax := src + length - 1
			ranges = append(
				ranges,
				Ranges{
					dest:    dest,
					destMax: destMax,
					src:     src,
					srcMax:  srcMax,
				},
			)

		}
		for idx, seed := range seeds {
			for _, rangeInfo := range ranges {
				if seed >= rangeInfo.src && seed <= rangeInfo.srcMax {
					difference := seed - rangeInfo.src
					seeds[idx] = rangeInfo.dest + difference
					break
				}
			}
		}
	}

	return slices.Min(seeds)
}
