package main

import (
	"fmt"
	"bufio"
	"os"
	"math"
	"sort"
)

func main() {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	max := 0
	var seats []int

	for _, l := range lines {
		row := mapToBinary(l[:7])
		column := mapToBinary(l[7:])
		product := row * 8 + column
		seats = append(seats, product)
		if product > max {
			max = product
		}
	}

	sort.Ints(seats)

	for idx, _ := range seats {
		if idx == 0 { continue }
		if seats[idx] == seats[idx-1] + 2 {
			fmt.Println(seats[idx] - 1)
			break
		}
	}

	fmt.Println(max)
}

func mapToBinary(s string) int {
	result := 0
	for idx, l := range s {
		if string(l) == "B" || string(l) == "R" {
			result += int(math.Pow(2, float64(len(s)-1)-float64(idx)))
		}
	}
	return result
}
