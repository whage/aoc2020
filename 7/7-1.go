package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
)

func createPattern(s string) string {
	return fmt.Sprintf(`(.+ bag)s contain.+%s`, s)
}

func removeDuplicates(list []string, pattern string) []string {
	results := []string{}
	for _, p := range list {
		if p != pattern { results = append(results, p)}
	}
	return results
}

func main() {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	patterns := []string{
		createPattern("shiny gold bag"),
	}
	parentBags := map[string]bool{}

	for len(patterns) > 0 {
		r := regexp.MustCompile(patterns[0])
		for _, line := range lines {
			matches := r.FindStringSubmatch(line)
			if matches != nil {
				patterns = append(patterns, createPattern(matches[1]))
				parentBags[matches[1]] = true
			}
		}
		patterns = removeDuplicates(patterns, patterns[0])
	}

	fmt.Println(len(parentBags))

}
