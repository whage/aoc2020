package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)

func main() {
	data, _ := ioutil.ReadFile("data.txt")
	allLines := strings.Split(string(data), "\n")
	numbers := []int{}

	for _, l := range allLines[:len(allLines)-1] {
		n, _ := strconv.Atoi(l)
		numbers = append(numbers, n)
	}

	sort.Ints(numbers)
	numbers = append(numbers, numbers[len(numbers)-1] +3)
	fmt.Println(numbers)

	flags := make([]bool, len(numbers))

	numberOfCombinations := 1

	for idx, _ := range numbers {
		if idx == 0 && numbers[1] <= 3 {
			flags[idx] = true
		}
		if idx >= 2 {
			if numbers[idx] - numbers[idx-2] <= 3 {
				flags[idx-1] = true
			}
		}
	}

	for _, v := range flags {
		
	}

	/*fmt.Println(indexesOfUselessOnes)
	fmt.Println("len", len(indexesOfUselessOnes))

	for idx, _ := range indexesOfUselessOnes {
		if idx > 0 {
			if indexesOfUselessOnes[idx] - indexesOfUselessOnes[idx-1] == 1 {
				numberOfCombinations += 2
			} else {
				numberOfCombinations *= 2
			}
		}
	}*/
	
	fmt.Println("numberOfCombinations", numberOfCombinations)
}
