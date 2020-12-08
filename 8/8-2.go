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

func copyInsturctions(i []Instruction) []Instruction {
	ret := make([]Instruction, len(i))
	for idx, v := range i {
		ret[idx] = v
	}
	return ret
}

func main() {
	data, _ := ioutil.ReadFile("data.txt")

	lines := strings.Split(string(data), "\n")

	instructions := []Instruction{}
	for _, ins := range lines {
		if ins != "" {
			instructions = append(instructions, Instruction{Op: ins, Visited: false})
		}
	}

	originalInstructions := copyInsturctions(instructions)
	
	ip := 0
	acc := 0
	lastSwitchedIdx := -1

	for {
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

		if ip == len(instructions) {
			fmt.Println("Terminated! Accumulator at %d", acc)
			break;
		}

		// if getting into a loop
		if instructions[ip].Visited {
			instructions = switchNext(copyInsturctions(originalInstructions), &lastSwitchedIdx)
			ip = 0
			acc = 0
		}
	}

	fmt.Println(acc)
}

func switchOp(ins Instruction) string {
	if ins.Op[:3] == "jmp" {
		return strings.ReplaceAll(ins.Op, "jmp", "nop")
	}

	if ins.Op[:3] == "nop" {
		return strings.ReplaceAll(ins.Op, "nop", "jmp")
	}

	return ins.Op
}

func switchNext(baseInstructions []Instruction, lastSwitchedIdx *int) []Instruction {
	for idx, ins := range baseInstructions {
		if ins.Op[:3] == "jmp" || ins.Op[:3] == "nop" {
			if idx > *lastSwitchedIdx {
				baseInstructions[idx].Op = switchOp(ins)
				*lastSwitchedIdx = idx
				break;
			}
		}
	}
	return baseInstructions
}
