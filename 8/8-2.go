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

	lastSwitched := -1
	
	ip := 0
	acc := 0

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
			instructions = switchNext(instructions, &lastSwitched)
			ip = 0
			acc = 0
		}
	}

	fmt.Println(acc)
}

func switchNext(instructions []Instruction, lastSwitched *int) []Instruction {
	ret := make([]Instruction, len(instructions))
	for idx, ins := range instructions {
		ret[idx] = ins
		ret[idx].Visited = false // reset

		if idx > *lastSwitched {
			fmt.Println(ret[idx].Op)
			if ret[idx].Op[:2] == "jmp" {
				ret[idx].Op = strings.ReplaceAll(ret[idx].Op, "nop", "jmp")
			}

			if ret[idx].Op[:3] == "nop" {
				ret[idx].Op = strings.ReplaceAll(ret[idx].Op, "jpm", "nop")
			}
		}

		if *lastSwitched >= 0 {
			if ret[*lastSwitched].Op[:3] == "jmp" {
				ret[*lastSwitched].Op = strings.ReplaceAll(ret[*lastSwitched].Op, "nop", "jmp")
			}

			if ret[*lastSwitched].Op[:3] == "nop" {
				ret[*lastSwitched].Op = strings.ReplaceAll(ret[*lastSwitched].Op, "jpm", "nop")
			}
		}
	}
	return ret
}
