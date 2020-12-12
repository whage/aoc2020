package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)

func main() {
	data, _ := ioutil.ReadFile("data.txt")
	allLines := strings.Split(string(data), "\n")
	numbers := []int{}

	for _, l := range allLines[:len(allLines)-1] {
		n, _ := strconv.Atoi(l)
		numbers = append(numbers, n)
	}

	sort.Ints(numbers)
	fmt.Println(numbers)

	sumOfDiffs := 0
	diffs := map[int]int{}

	for idx, n := range numbers {
		if idx > 0 {
			sumOfDiffs += n - numbers[idx-1]
			diffs[n - numbers[idx-1]] += 1
		} else {
			sumOfDiffs += n
			diffs[n] += 1
		}
	}

	sumOfDiffs += 3
	diffs[3] += 1

	fmt.Println(sumOfDiffs)
	fmt.Println(diffs)

	fmt.Println(diffs[3] * diffs[1])
}
