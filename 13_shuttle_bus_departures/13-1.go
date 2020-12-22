package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"math"
)

func filterAndConvert(s []string, pred func(string)bool) []int {
	ret := []int{}
	for _, c := range s {
		if pred(c) {
			n, _ := strconv.Atoi(c)
			ret = append(ret, n)
		}
	}
	return ret
}

func main() {
	data, _ := ioutil.ReadFile("data-2.txt")
	split := strings.Split(string(data), "\n")
	allLines := split[:len(split)-1]
	earliestDeparture, _ := strconv.Atoi(allLines[0])
	inService := filterAndConvert(strings.Split(allLines[1], ","), func(c string) bool {
		return c != "x"
	})

	busToTake := inService[0]
	minDifference := 59 // TODO
	finalMultiple := 0

	for _, n := range inService {
		if earliestDeparture % n == 0 {
			return
		}

		floorOfRatio := math.Floor(float64(earliestDeparture) / float64(n))
		multiple := (int(floorOfRatio) + 1) * n
		diff := (multiple) - earliestDeparture
		if diff <= minDifference {
			minDifference = diff
			busToTake = n
			finalMultiple = multiple
		}
	}

	fmt.Println((finalMultiple - earliestDeparture)*busToTake)
}
