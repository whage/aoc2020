package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Pos struct {
	X, Y, Z, W int
}

type Grid map[Pos]bool

type NeighborCountGrid map[Pos]int

func readGrid(inputData []byte) Grid {
	lines := strings.Split(string(inputData), "\n")

	grid := Grid{}

	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y++ {
			if lines[y][x] == '#' {
				grid[Pos{x, y, 0, 0}] = true
			}
		}
	}

	return grid
}

func getNeighborCounts(grid Grid) NeighborCountGrid {
	gridOfNeighborCounts := NeighborCountGrid{}
	for pos, _ := range grid {
		for dx := -1; dx < 2; dx++ {
			for dy := -1; dy < 2; dy++ {
				for dz := -1; dz < 2; dz++ {
					for dw := -1; dw < 2; dw++ {
						if dx == 0 && dy == 0 && dz == 0 && dw == 0 { continue }
						targetPos := Pos{pos.X+dx, pos.Y+dy, pos.Z+dz, pos.W+dw}
						gridOfNeighborCounts[targetPos] += 1
					}
				}
			}
		}
	}
	return gridOfNeighborCounts
}

func step(grid Grid, counter NeighborCountGrid) Grid {
	newGrid := Grid{}
	for pos, neighborCount := range counter {
		_, isAlive := grid[pos]
		if isAlive {
			if neighborCount == 2 || neighborCount == 3 {
				newGrid[pos] = true
			}
		} else {
			if neighborCount == 3 {
				newGrid[pos] = true
			}
		}
	}
	return newGrid
}

func main() {
	data, _ := ioutil.ReadFile("data-1.txt")

	grid := readGrid(data)
	stepCount := 1

	for stepCount <= 6 {
		grid = step(grid, getNeighborCounts(grid))
		stepCount++
	}

	fmt.Println(len(grid))
}
