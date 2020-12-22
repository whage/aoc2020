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

func main() {
	data, _ := ioutil.ReadFile("data.txt")
	sections := strings.Split(string(data), "\n\n")

	validTickets := []string{}

	rulesSection := sections[0]
	rules := make([]Range, 0)

	for _, ruleLine := range strings.Split(rulesSection, "\n") {
		r := regexp.MustCompile(`\w+: (\d+)\-(\d+) or (\d+)\-(\d+)`)
		matches := r.FindStringSubmatch(ruleLine)

		n1, _ := strconv.Atoi(matches[1])
		n2, _ := strconv.Atoi(matches[2])
		n3, _ := strconv.Atoi(matches[3])
		n4, _ := strconv.Atoi(matches[4])

		rules = append(rules, Range{Min: n1, Max: n2})
		rules = append(rules, Range{Min: n3, Max: n4})
	}

	nearbyTicketsSection := sections[2]

	for _, ticket := range strings.Split(nearbyTicketsSection, "\n") {
		allAreValid := true
		for _, number := range strings.Split(ticket, ",") {
			n, _ := strconv.Atoi(number)
			valid := false
			for _, r := range rules {
				if n >= r.Min && n <= r.Max {
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

	//fmt.Println(len(validTickets))
}
