package main

import (
	"fmt"
)

type Pair struct {
	Last int
	BeforeLast int
}

func (p *Pair) Shift(currentTurn int) {
	p.BeforeLast = p.Last
	p. Last = currentTurn
}

func main() {
	numbers := []int{0,3,6}
	turnMap := map[int]Pair{
		0: Pair{Last: 1, BeforeLast: 0},
		3: Pair{Last: 2, BeforeLast: 0},
		6: Pair{Last: 3, BeforeLast: 0},
	}

	turnCount := 4
	lastNumberSpoken := numbers[len(numbers)-1]

	//for turnCount <= 2020 {
	for turnCount <= 6 {
		fmt.Println("lastNumberSpoken", lastNumberSpoken)
		fmt.Println("turnCount",turnCount)
		pair, ok := turnMap[lastNumberSpoken]
		fmt.Println("pair at beginngin", pair)
		if !ok {
			lastNumberSpoken = 0
			turnMap[turnCount] = Pair{Last: turnCount, BeforeLast: 0}
		} else {
			if pair.BeforeLast == 0 {
				lastNumberSpoken = 0
			} else {
				lastNumberSpoken = pair.Last - pair.BeforeLast
			}
			fmt.Println("pair after shifting", pair)
		}
		fmt.Println("next lastNumberSpoken", lastNumberSpoken)
		fmt.Println()
		
		turnCount++
	}

	fmt.Println(lastNumberSpoken)
}
