package chapters

import (
	"bytes"
	"strconv"
)

func Chapter2(data []byte) int {
	data = bytes.Replace(data, []byte("Time: "), []byte(""), -1)
	data = bytes.Replace(data, []byte("Distance: "), []byte(""), -1)
	data = bytes.Replace(data, []byte(" "), []byte(""), -1)

	info := bytes.Split(data, []byte("\n"))
	time, _ := strconv.Atoi(string(info[0]))
	record, _ := strconv.Atoi(string(info[1]))

	var value int

	for idx := 1; idx < time-1; idx++ {
		if idx*(time-idx) > record {
			value++
		}
	}

	return value
}
