package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("map.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	slopes := [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	treeCounts := make([]int, len(slopes))

	for i, slope := range slopes {
		rowIdx := 0
		columnIdx := 0
		treeCount := 0
		for rowIdx < len(lines)-slope[1] {
			nextField := getNextFieldInSlope(lines, rowIdx, columnIdx, slope[0], slope[1])
			if isTree(nextField) { treeCount++ }
			rowIdx += slope[1]
			columnIdx += slope[0]
		}		
		fmt.Printf("Tree count for slope %v: %d\n", slope, treeCount)
		treeCounts[i] = treeCount
	}

	fmt.Printf("%d\n", sum(treeCounts))
}

func sum(numbers []int) int {
	sum := 1
	for _, n := range numbers {
		sum *= n
	}
	return sum
}

func isTree(c string) bool {
	return c == "#"
}

func getNextFieldInSlope(lines []string, rowIdx, columnIdx, dx, dy int) string {
	lineWidth := len(lines[0])
	return string(lines[rowIdx+dy][(columnIdx+dx)%lineWidth])
}
