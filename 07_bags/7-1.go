package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strings"
	"strconv"
)

type BagCount struct {
	Name string
	Count int
}

func main() {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	m := map[string][]BagCount{}

	for _, line := range lines {
		r := regexp.MustCompile(`(.+) bags contain`)
		parentBagMatches := r.FindStringSubmatch(line)
		if parentBagMatches != nil {
			parentBag := parentBagMatches[1]
			startIdx := strings.Index(line, "contain") + 8
			containedBags := []BagCount{}
			containedBagsParts := strings.Split(line[startIdx:], ",")
			for _, str := range containedBagsParts {
				r := regexp.MustCompile(`(\d+) (.+) bag(s)?`)
				matches := r.FindStringSubmatch(str)
				if matches != nil {
					n, _ := strconv.Atoi(matches[1])
					containedBags = append(containedBags, BagCount{Name: matches[2], Count: n})
				}
			}
			if len(containedBags) > 0 {
				m[parentBag] = containedBags
			}
		}
	}
	fmt.Println(summarize(m, "shiny gold"))
}

func summarize(m map[string][]BagCount, key string) int {
	containedBags, ok := m[key]
	if !ok {
		return 1
	}
	sum := 0
	for _, bagType := range containedBags {
		sum += bagType.Count * summarize(m, bagType.Name)
		_, ok := m[bagType.Name]
		if ok {
			sum += bagType.Count
		}

	}
	return sum
}
