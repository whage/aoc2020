package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"strconv"
)

type Range struct {
	Min int
	Max int
}

type Rule struct {
	Name string
	ValidRanges [2]Range
}

func (r Rule) isWithinRanges(n int) bool {
	isWithinLowerRange := n >= r.ValidRanges[0].Min && n <= r.ValidRanges[0].Max
	isWithinUpperRange := n >= r.ValidRanges[1].Min && n <= r.ValidRanges[1].Max
	return isWithinLowerRange || isWithinUpperRange
}

type Set map[int]bool

func (s *Set) add(c int) {
	(*s)[c] = true
}

func (s *Set) remove(c int) {
	(*s)[c] = false
}

func (s *Set) intersect(other *Set) *Set {
	result := Set{}
	for k, _ := range *s {
		v, ok := (*other)[k]
		if ok { result[k] = v }
	}
	return &result
}

func newSet(numbers []int) Set {
	res := Set{}
	for _, v := range numbers {
		res[v] = true
	}
	return res
}

func main() {
	data, _ := ioutil.ReadFile("data.txt")
	sections := strings.Split(string(data), "\n\n")

	validTickets := []string{}

	rulesSection := sections[0]
	rules := make([]Rule, 0)

	for _, ruleLine := range strings.Split(rulesSection, "\n") {
		r := regexp.MustCompile(`(\w+ ?\w+): (\d+)\-(\d+) or (\d+)\-(\d+)`)
		matches := r.FindStringSubmatch(ruleLine)

		n1, _ := strconv.Atoi(matches[2])
		n2, _ := strconv.Atoi(matches[3])
		n3, _ := strconv.Atoi(matches[4])
		n4, _ := strconv.Atoi(matches[5])

		rules = append(rules, Rule{
			Name: matches[1],
			ValidRanges: [2]Range{
				Range{Min: n1, Max: n2},
				Range{Min: n3, Max: n4},
			},
		})
	}

	nearbyTicketsSection := sections[2]

	for _, ticket := range strings.Split(nearbyTicketsSection, "\n") {
		allAreValid := true
		for _, number := range strings.Split(ticket, ",") {
			n, _ := strconv.Atoi(number)
			valid := false
			for _, r := range rules {
				if  r.isWithinRanges(n) {
					valid = true
					break
				}
			}
			if !valid {
				allAreValid = false
			}
		}
		if allAreValid {
			validTickets = append(validTickets, ticket)
		}
	}

	potentialMappings := make([]Set, len(rules))
	for idx, _ := range rules {
		potentialMappings[idx] = newSet([]int{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19})
	}

	for _, ticket := range validTickets {
		for fieldIdx, field := range strings.Split(ticket, ",") {
			n, _ := strconv.Atoi(field)
			for ruleIdx, rule := range rules {
				if ! rule.isWithinRanges(n) {
					potentialMappings[fieldIdx].remove(ruleIdx)
				}
			}
		}
	}

	for mappingIdx, mapping := range potentialMappings {
		potentialFieldCount := 0
		lastPotentialField := ""
		for key, v := range mapping {
			if v {
				potentialFieldCount++
				lastPotentialField = rules[key].Name
			}
		}
		if potentialFieldCount == 1 {
			fmt.Println("mappintIdx", mappingIdx)
			fmt.Println("mapping", mapping)
			fmt.Printf("Field at idx %d must be %s\n", mappingIdx, lastPotentialField)
		}
	}
}
