package main

import (
	"fmt"
	"regexp"
	"strconv"
	"bufio"
	"os"
)

func main() {
	r := regexp.MustCompile(`(\d+)-(\d+) (.): (.+)`)

	validCount := 0

	file, _ := os.Open("2-data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		matches := r.FindStringSubmatch(line)
			fmt.Printf("%#v\n", matches)
		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		character := matches[3][0]
		password := matches[4]

		count := 0
		for _, c := range password {
			if c == rune(character) { count++}
		}
		if count >= min && count <= max {
			validCount++
		}
	}

	fmt.Println(validCount)
}
