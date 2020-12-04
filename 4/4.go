package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

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

	validCount := 0

	for _, pp := range passports {
		requiredSegments := map[string]bool{
			"byr": false,
			"iyr": false,
			"eyr": false,
			"hgt": false,
			"hcl": false,
			"ecl": false,
			"pid": false,
			//"cid": false,
		}
		foundAll := true
		for _, seg := range strings.Split(pp, " ") {
			requiredSegments[seg[:3]] = true
		}
		fmt.Println(requiredSegments)
		for _, reqSeg := range requiredSegments {
			if !reqSeg { foundAll = false }
		}
		if foundAll {
			fmt.Println("VALID")
			validCount++
		}
	}

	//fmt.Println(passports)
	fmt.Println(validCount)
}
