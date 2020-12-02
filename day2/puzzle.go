// Package day2 has solutions for Day 2 of Advent of Code
// https://adventofcode.com/2020/day/2
package day2

import (
	"log"
	"strconv"
	"strings"

	"../input"
)

var inputFile string = "day2/input.txt"
var testFile string = "day2/testInput.txt"

type password struct {
	min          int
	max          int
	policyLetter string
	password     string
}

func parseInput(inputList []string) []password {
	var passwords []password
	for _, value := range inputList {
		split := strings.Fields(value)
		policyRange := strings.Split(split[0], "-")

		min, err := strconv.Atoi(policyRange[0])
		if err != nil {
			log.Fatal("Failed to convert min to num", err)
		}

		max, err := strconv.Atoi(policyRange[1])
		if err != nil {
			log.Fatal("Failed to convert max to num", err)
		}
		pass := password{min: min}
		pass.max = max
		pass.policyLetter = string(rune(split[1][0]))
		pass.password = split[2]

		passwords = append(passwords, pass)
	}

	return passwords

}

// PartOne finds
func PartOne() int {
	rawInput := input.ReadInput(inputFile)
	inputList := input.Slice(rawInput)
	passwords := parseInput(inputList)

	numValid := 0

	for _, pass := range passwords {
		count := strings.Count(pass.password, pass.policyLetter)
		if count >= pass.min && count <= pass.max {
			numValid++
		}
	}

	return numValid
}

// PartTwo finds
func PartTwo() int {
	rawInput := input.ReadInput(inputFile)
	inputList := input.Slice(rawInput)
	passwords := parseInput(inputList)

	numValid := 0

	for _, pass := range passwords {
		first := string(rune(pass.password[pass.min-1]))
		second := string(rune(pass.password[pass.max-1]))

		if (first != second) && (first == pass.policyLetter || second == pass.policyLetter) {
			numValid++
		}
	}

	return numValid
}
