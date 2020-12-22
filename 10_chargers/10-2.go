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
	numbers = append(numbers, numbers[len(numbers)-1] +3)
	fmt.Println(numbers)

	flags := make([]bool, len(numbers))

	for idx, _ := range numbers {
		if idx == 0 && numbers[1] <= 3 {
			flags[idx] = true
		}
		if idx >= 2 {
			if numbers[idx] - numbers[idx-2] <= 3 {
				flags[idx-1] = true
			}
		}
	}

	fmt.Println(flags)

	runLengths := []int{}
	counter := 0

	for _, v := range flags {
		if v {
			counter++
		} else {
			runLengths = append(runLengths, counter)
			counter = 0
		}
	}

	fmt.Println(runLengths)

	product := 1

	for _, rl := range runLengths {
		if rl == 1 { product *= 2 }
		if rl == 2 { product *= 4 }
		if rl == 3 { product *= 7 }
	}

	fmt.Println(product)
}
