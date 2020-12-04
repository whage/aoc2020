package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strings"
	"strconv"
)

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func main() {
	file, _ := os.Open("passports.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var passports []string

	currentPassport := ""

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passports = append(passports, strings.Trim(currentPassport, " "))
			currentPassport = ""
		} else {
			currentPassport += " " + line
		}
	}

	if currentPassport != "" { passports = append(passports, strings.Trim(currentPassport, " ")) }

	validCount := 0

	rules := map[string]func(string)bool{
		"byr": func(s string) bool { return atoi(s) >= 1920 && atoi(s) <= 2002 },
		"iyr": func(s string) bool { return atoi(s) >= 2010 && atoi(s) <= 2020 },
		"eyr": func(s string) bool { return atoi(s) >= 2020 && atoi(s) <= 2030 },
		"hgt": func(s string) bool {
			r := regexp.MustCompile(`(\d+)(cm|in)`)
			matches := r.FindStringSubmatch(s)
			if matches == nil { return false }
			if matches[2] == "cm" { return atoi(matches[1]) >= 150 && atoi(matches[1]) <= 193 }
			if matches[2] == "in" { return atoi(matches[1]) >= 59 && atoi(matches[1]) <= 76 }
			return false
		},
		"hcl": func(s string) bool {
			return regexp.MustCompile(`#([0-9a-fA-F]{6})`).FindStringSubmatch(s) != nil
		},
		"ecl": func(s string) bool {
			return regexp.MustCompile(`(amb|blu|brn|gry|grn|hzl|oth)`).FindStringSubmatch(s) != nil
		},
		"pid": func(s string) bool {
			return regexp.MustCompile(`^[0-9]{9}$`).FindStringSubmatch(s) != nil
		},
	}

	for _, pp := range passports {
		requiredSegments := map[string]bool{
			"byr": false,
			"iyr": false,
			"eyr": false,
			"hgt": false,
			"hcl": false,
			"ecl": false,
			"pid": false,
		}
		foundAll := true
		for _, seg := range strings.Split(pp, " ") {
			if seg[:3] == "cid" { continue }
			ruleOk := rules[seg[:3]](seg[4:])
			if ruleOk {
				requiredSegments[seg[:3]] = true
			}
		}
		for _, reqSeg := range requiredSegments {
			if !reqSeg { foundAll = false }
		}
		if foundAll {
			validCount++
		}
	}

	fmt.Println(validCount)
}
