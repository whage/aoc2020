package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func getConstraints(s []string) map[int]int {
	ret := map[int]int{}
	for idx, c := range s {
		if c != "x" {
			n, _ := strconv.Atoi(c)
			ret[n] = idx
		}
	}
	return ret
}

func getBusIds(s []string) []int {
	ret := []int{}
	for _, v := range s {
		if v != "x" {
			n, _ := strconv.Atoi(v)
			ret = append(ret, n)
		}
	}
	return ret
}

func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int64) int64 {
	return a * b / GCD(a, b)
}

func main() {
	data, _ := ioutil.ReadFile("data-final.txt")
	split := strings.Split(string(data), "\n")
	allLines := split[:len(split)-1]
	items := strings.Split(allLines[1], ",")
	constraints := getConstraints(items)
	busIds := getBusIds(items)

	var base int64 = 0
	var step int64 = int64(busIds[0])

	// Counts up with an increasing step size.
	// Key insight: once a pair of constraints is satisfied,
	// step size can be set to LCM(step, next-number).
	for idx, _ := range busIds {
		if idx == len(busIds) - 1 { break }
		for {
			base += step
			next := base + int64(constraints[busIds[idx+1]])
			divisor := int64(busIds[idx+1])
			if next % divisor == 0 {
				break
			}
		}
		step = LCM(step, int64(busIds[idx+1]))
	}

	fmt.Println(base)
}
