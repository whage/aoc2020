package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)

func isSumOfAny2(l []int, n int) bool {
	for i:=0; i < len(l); i++ {
		for j:=0; j < len(l); j++ {
			if j != i && l[i]+l[j] == n { return true }
		}
	}
	return false
}

func sum(l []int) int {
	sum := 0
	for i:=0; i < len(l); i++ {
		sum += l[i]
	}
	return sum
}

func copySlice(l []int) []int {
	ret := make([]int, len(l))
	for idx, v := range l {
		ret[idx] = v
	}
	return ret
}

func main() {
	data, _ := ioutil.ReadFile("data.txt")
	lines := strings.Split(string(data), "\n")
	numbers := make([]int, len(lines)-1) // omit last empty line
	var invalidOne int
	for idx, _ := range numbers {
		n, _ := strconv.Atoi(lines[idx])
		numbers[idx] = n
		if idx >= 25 {
			window := numbers[idx-25:idx]
			if !isSumOfAny2(window, numbers[idx]) {
				invalidOne = numbers[idx]
			}
		}
	}
	fmt.Println(invalidOne)
	fmt.Println(numbers)
	for windowSize := 2; windowSize <= len(numbers); windowSize++ {
		for i := 0; i <= len(numbers)-windowSize; i++ {
			currentSection := copySlice(numbers[i:i+windowSize])
			sort.Ints(currentSection)
			total := sum(currentSection)
			if total == invalidOne {
				fmt.Println(currentSection[0] + currentSection[len(currentSection)-1])
				return
			}
		}
	}
}
