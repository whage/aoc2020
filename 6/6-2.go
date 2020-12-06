package main

import (
	"fmt"
	"bufio"
	"os"
)

type Set map[rune]bool

func (s *Set) add(c rune) {
	(*s)[c] = true
}

func (s *Set) intersect(other *Set) *Set {
	result := Set{}
	for k, _ := range *s {
		v, ok := (*other)[k]
		if ok { result[k] = v }
	}
	return &result
}

func newSet(s string) Set {
	res := Set{}
	for _, v := range s {
		res[v] = true
	}
	return res
}

func main() {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalCount := 0
	currentGroup := []Set{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			totalCount += getUniqueCount(currentGroup)
			currentGroup = []Set{}
		} else {
			currentGroup = append(currentGroup, newSet(line))
		}
	}

	totalCount += getUniqueCount(currentGroup)

	fmt.Println(totalCount)
}

func getUniqueCount(sets []Set) int {
	accumulated := &sets[0]

	for idx, s := range sets {
		if idx > 0 {
			accumulated = accumulated.intersect(&s)
		}
	}

	count := 0
	for _, v := range *accumulated {
		if v { count += 1 }
	}

	return count
}
