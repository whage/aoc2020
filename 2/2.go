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
		pos1, _ := strconv.Atoi(matches[1])
		pos2, _ := strconv.Atoi(matches[2])
		character := matches[3][0]
		password := matches[4]

		pos1_ok := password[pos1-1] == character
		pos2_ok := password[pos2-1] == character
		if pos1_ok != pos2_ok {
			validCount++
		}
	}

	fmt.Println(validCount)
}
