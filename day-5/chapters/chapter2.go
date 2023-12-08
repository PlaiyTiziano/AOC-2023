package chapters

import (
	"slices"
	"strconv"
	"strings"
	"sync"
)

func Chapter2(content []byte) int {
	instructions := strings.Split(string(content), "\n\n")

	seedRanges := [][]int{}
	for idx, seed := range strings.Split(strings.Replace(instructions[0], "seeds: ", "", -1), " ") {
		number, _ := strconv.Atoi(seed)

		if idx%2 == 0 {
			seedRanges = append(seedRanges, []int{number})
		} else {
			seedRanges[len(seedRanges)-1] = append(seedRanges[len(seedRanges)-1], number)
		}
	}

	var wg sync.WaitGroup
	var results []int
	ch := make(chan int)
	done := make(chan bool)

	for seedRangesIdx, seedRange := range seedRanges {
		wg.Add(1)

		go func(seedRange []int) {
			defer wg.Done()

			seeds := []int{}

			for idx := 0; idx < seedRange[1]; idx++ {
				seeds = append(seeds, seedRange[0]+idx)
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

			ch <- slices.Min(seeds)
			if seedRangesIdx == len(seedRanges)-1 {
				done <- true
			}
		}(seedRange)
	}

L:
	for {
		select {
		case v := <-ch:
			results = append(results, v)
			if len(results) == len(seedRanges) {
				close(ch)
				break L
			}
		case _ = <-done:
			break
		}
	}
	return slices.Min(results)
}
