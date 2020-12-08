// Package day4 has solutions for Day 4 of Advent of Code
// https://adventofcode.com/2020/day/4
package day4

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var inputFile string = "day4/input.txt"
var testFile string = "day4/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(inputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(inputFile))
}

// parsePassport parses an input file of passports into a slice of passports
// where each passport is a map from the passport field to its value
func parsePassports(file string) []map[string]string {
	input := input.ReadInput(file)
	splitInput := strings.Split(input, "\n\n")

	passports := make([]map[string]string, len(splitInput))
	for _, val := range splitInput {
		passport := make(map[string](string))
		for _, pair := range strings.Fields(val) {
			split := strings.Split(pair, ":")
			passport[split[0]] = split[1]
		}
		passports = append(passports, passport)
	}
	return passports
}

// PartOne finds how many passports have all required fields ('cid' is optional)
func PartOne(file string) int {
	passports := parsePassports(file)
	valid := 0
	for _, passport := range passports {
		if len(passport) == 8 || (len(passport) == 7 && passport["cid"] == "") {
			valid++
		}
	}

	return valid
}

// PartTwo finds how many passports are valid based on the requirements for each field
func PartTwo(file string) int {
	passports := parsePassports(file)
	valid := 0
	for _, passport := range passports {
		if validatePassport(passport) {
			valid++
		}
	}
	return valid
}

var fields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}

// validatePassport confirms that the passport has all fields needed for a passport (defined
// in the fields global, with 'cid' being optional), and that these fields contain valid
// values as defined by the problem spec
func validatePassport(passport map[string](string)) bool {
	for _, field := range fields {
		value := passport[field]
		if value == "" && field != "cid" {
			return false
		}
		var matched bool
		switch field {
		case "cid":
			continue
		case "byr": // byr (Birth Year) - four digits; at least 1920 and at most 2002.
			matched, _ = regexp.MatchString(`^(19[2-9][0-9]|200[0-2])$`, value)
		case "iyr": // iyr (Issue Year) - four digits; at least 2010 and at most 2020.
			matched, _ = regexp.MatchString(`^(201[0-9]|2020)$`, value)
		case "eyr": // eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
			matched, _ = regexp.MatchString(`^(202[0-9]|2030)$`, value)
		case "hgt":
			// hgt (Height) - a number followed by either cm or in:
			// If cm, the number must be at least 150 and at most 193.
			// If in, the number must be at least 59 and at most 76.
			matchedCm, _ := regexp.MatchString(`^(1[5-8][0-9]|19[0-3])cm$`, value)
			matchedIn, _ := regexp.MatchString(`^(59|6[0-9]|7[0-6])in$`, value)
			matched = matchedCm || matchedIn
		case "hcl": // hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
			matched, _ = regexp.MatchString(`^#[0-9a-f]{6}$`, value)
		case "ecl": // ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
			matched, _ = regexp.MatchString(`^amb|blu|brn|gry|grn|hzl|oth$`, value)
		case "pid": // pid (Passport ID) - a nine-digit number, including leading zeroes.
			matched, _ = regexp.MatchString(`^[0-9]{9}$`, value)
		}
		if !matched {
			return false
		}
	}
	return true
}
