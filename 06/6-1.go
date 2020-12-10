package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	totalCount := 0

	currentGroup := map[rune]bool{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			totalCount += getUniqueCount(currentGroup)
			currentGroup = map[rune]bool{}
		} else {
			for _, c := range line {
				currentGroup[c] = true
			}
		}
	}

	totalCount += getUniqueCount(currentGroup)

	fmt.Println(totalCount)
}

func getUniqueCount(m map[rune]bool) int {
	count := 0
	for _, v := range m {
		if v { count += 1 }
	}
	return count
}
