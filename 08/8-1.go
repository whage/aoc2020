package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"strconv"
)

type Instruction struct {
	Op string
	Visited bool
}

func main() {
	data, _ := ioutil.ReadFile("data.txt")

	lines := strings.Split(string(data), "\n")

	instructions := []Instruction{}
	for _, ins := range lines {
		instructions = append(instructions, Instruction{Op: ins, Visited: false})
	}
	
	ip := 0
	acc := 0

	for ( ! instructions[ip].Visited) {
		instructions[ip].Visited = true
		r := regexp.MustCompile(`([a-z]{3}) ([+|-]\d+)`)
		matches := r.FindStringSubmatch(instructions[ip].Op)
		op := matches[1]
		arg, _ := strconv.Atoi(matches[2])
		switch op {
		case "acc":
			acc += arg
			ip += 1
		case "nop":
			ip += 1
		case "jmp":
			ip += arg
		}
	}

	fmt.Println(acc)
}
