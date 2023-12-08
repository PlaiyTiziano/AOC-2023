package chapters

import (
	"bytes"
	"regexp"
	"strconv"
)

type Race struct {
	time   int
	record int
}

func parseRaceInfo(timeInfo [][]byte, records [][]byte) []Race {
	var races []Race

	for idx, t := range timeInfo {
		time, _ := strconv.Atoi(string(t))
		record, _ := strconv.Atoi(string(records[idx]))

		races = append(races, Race{
			time:   time,
			record: record,
		})
	}

	return races
}

func Chapter1(data []byte) int {
	regex := regexp.MustCompile(`\d+`)

	info := bytes.Split(data, []byte("\n"))

	races := parseRaceInfo(
		regex.FindAll(info[0], -1),
		regex.FindAll(info[1], -1),
	)

	value := 1
	for _, race := range races {
		var winners int

		for idx := 1; idx < race.time-1; idx++ {
			if idx*(race.time-idx) > race.record {
				winners++
			}
		}

		if winners > 0 {
			value *= winners
		}
	}

	return value
}
