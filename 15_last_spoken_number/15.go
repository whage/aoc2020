package main

import (
	"fmt"
)

type Pair struct {
	Last int
	BeforeLast int
}

func (p *Pair) Shift(currentTurn int) Pair {
	p.BeforeLast = p.Last
	p. Last = currentTurn
	return *p
}

func main() {
	numbers := []int{0,3,1,6,7,5}
	tracker := map[int]Pair{}

	for idx, i := range numbers {
		tracker[i] = Pair{Last: idx+1, BeforeLast: 0}
	}

	count := len(numbers)+1
	last := numbers[len(numbers)-1]
	var newNumber int

	for count <= 30000000 {
		pair, _ := tracker[last]
		if pair.BeforeLast == 0 {
			newNumber = 0
		} else {
			newNumber = pair.Last - pair.BeforeLast
		}

		t, ok := tracker[newNumber]
		if !ok {
			tracker[newNumber] = Pair{Last: count, BeforeLast: 0}
		} else {
			tracker[newNumber] = t.Shift(count)
		}

		last = newNumber
		count++
	}

	fmt.Println(last)
}
