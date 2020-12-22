package main

import (
	"fmt"
)

type Coord struct {
	X, Y int
}

type Cell struct {
	NeighborCount int
	IsAlive bool
}

func main() {
	m := map[Coord]int{
		Coord{1, 1}: 1,
		Coord{2, 1}: 2,
		Coord{3, 1}: 1,
	}
	fmt.Println(m)
}
