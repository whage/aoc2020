package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"regexp"
)

func xTo(c, all string) uint64 {
	transformed := strings.ReplaceAll(all, "X", c)
	n, _ := strconv.ParseInt(transformed, 2, 64)
	return uint64(n)
}

func main() {
	data, _ := ioutil.ReadFile("data-1-2.txt")
	split := strings.Split(string(data), "\n")
	allLines := split[:len(split)-1]

	memory := map[uint64]uint64{}
	var m0 uint64
	var m1 uint64

	for _, line := range allLines {
		if line[:4] == "mask" {
			m0 = xTo("0", line[7:])
			m1 = xTo("1", line[7:])
		} else {
			r := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
			matches := r.FindStringSubmatch(line)
			addr, _ := strconv.Atoi(matches[1])
			data, _ := strconv.Atoi(matches[2])
			memory[uint64(addr)] = m0 | (uint64(data) & m1)
		}
	}

	var sum uint64
	for _, v := range memory {
		sum += v
	}
	fmt.Println(sum)
}
