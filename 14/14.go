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

func getPowerSet(items []int, acc [][]int, level int) [][]int {
	if len(items) == 0 { return acc }
	withCurrentOne := append(acc, []int{items[0]})
	withoutCurrentOne := acc
	return append(acc, getPowerSet(items[1:], withCurrentOne), getPowerSet(items[1:], withoutCurrentOne))
}

func solve14_1() {
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


func solve14_2() {
	data, _ := ioutil.ReadFile("data-1-2.txt")
	split := strings.Split(string(data), "\n")
	allLines := split[:len(split)-1]

	memory := map[uint64]uint64{}
	//var m0 uint64
	//var m1 uint64
	var mask string

	for _, line := range allLines {
		if line[:4] == "mask" {
			mask = line[7:]
			//m1 = xTo("1", mask)
		} else {
			r := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
			matches := r.FindStringSubmatch(line)
			baseAddress, _ := strconv.Atoi(matches[1])
			fmt.Println("baseAddress", baseAddress)
			//partiallyMaskedAddress := m0 | m1
			data, _ := strconv.Atoi(matches[2])
			fmt.Println("data", data)
			
			xPlaces := make([]int, 0)
			for idx, c := range mask {
				if c == 'X' { xPlaces = append(xPlaces, len(mask)-1-idx) }
			}

			//fmt.Println("xPlaces", xPlaces)
			/*for _, addr := range getPowerSet(xPlaces) {

				//memory[uint64(addr)] = m0 | (uint64(data) & m1)
			}*/
		}
	}

	var sum uint64
	for _, v := range memory {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	solve14_1()
	solve14_2()
}
