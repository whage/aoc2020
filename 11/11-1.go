package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func showGrid(grid [][]string) {
	for y := 0; y < len(grid[0]); y++ {
		for x := 0; x < len(grid); x++ {
			fmt.Printf("%s", grid[x][y])
		}
		fmt.Println()
	}
	fmt.Println()
}

func copyGrid(grid [][]string) [][]string {
	ret := make([][]string, len(grid))
	for i := 0; i < len(grid); i++ {
		ret[i] = make([]string, len(grid[0]))
	}
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			ret[x][y] = grid[x][y]
		}
	}
	return ret
}

func isEmpty(grid [][]string, x, y int) bool {
	if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
		if grid[x][y] == "L" {
			return true
		}
	}
	return false
}

func isOccupied(grid [][]string, x, y int) bool {
	if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
		if grid[x][y] == "#" {
			return true
		}
	}
	return false
}

func hasFourOrMoreOccupiedNeighbors(grid [][]string, x, y int) bool {
	occupiedCount := 0

	if isOccupied(grid, x-1, y-1) { occupiedCount += 1 }
	if isOccupied(grid, x, y-1)   { occupiedCount += 1 }
	if isOccupied(grid, x+1, y-1) { occupiedCount += 1 }
	if isOccupied(grid, x-1, y)   { occupiedCount += 1 }
	if isOccupied(grid, x+1, y)   { occupiedCount += 1 }
	if isOccupied(grid, x-1, y+1) { occupiedCount += 1 }
	if isOccupied(grid, x, y+1)   { occupiedCount += 1 }
	if isOccupied(grid, x+1, y+1) { occupiedCount += 1 }

	return occupiedCount >= 4
}

func changeCell(grid [][]string, x, y int) (bool, string) {
	if isEmpty(grid, x, y) &&
		!isOccupied(grid, x-1, y-1) &&
		!isOccupied(grid, x, y-1) &&
		!isOccupied(grid, x+1, y-1) &&
		!isOccupied(grid, x-1, y) &&
		!isOccupied(grid, x+1, y) &&
		!isOccupied(grid, x-1, y+1) &&
		!isOccupied(grid, x, y+1) &&
		!isOccupied(grid, x+1, y+1) {
			return true, "#"
		}

	if isOccupied(grid, x, y) && hasFourOrMoreOccupiedNeighbors(grid, x, y) {
		return true, "L"
	}

	return false, ""
}

func countOccupied(grid [][]string) int {
	count := 0
	for y := 0; y < len(grid[0]); y++ {
		for x := 0; x < len(grid); x++ {
			if grid[x][y] == "#" { count++ }
		}
	}
	return count
}

func main() {
	data, _ := ioutil.ReadFile("data-1.txt")
	allLines := strings.Split(string(data), "\n")
	grid := make([][]string, len(allLines[0]))

	for i := 0; i < len(allLines[0]); i++ {
		grid[i] = make([]string, len(allLines)-1)
	}

	for y_idx, l := range allLines[:len(allLines)-1] {
		for x_idx, c := range l {
			grid[x_idx][y_idx] = string(c)
		}
	}

	stepCount := 0
	numChangedInLastRound := 1

	for numChangedInLastRound > 0 {
		numChangedInLastRound = 0
		copyOfGrid := copyGrid(grid)
		for x := 0; x < len(grid); x++ {
			for y := 0; y < len(grid[0]); y++ {
				//fmt.Println("x", x, "y", y)
				changed, newValue := changeCell(grid, x, y)
				//fmt.Println("changed", changed, "newValue", newValue)
				if changed {
					copyOfGrid[x][y] = newValue
					numChangedInLastRound++
				}
			}
		}
		//showGrid(copyOfGrid)
		fmt.Println(numChangedInLastRound)
		stepCount += 1
		grid = copyOfGrid
	}

	showGrid(grid)

	fmt.Println(countOccupied(grid))

}
