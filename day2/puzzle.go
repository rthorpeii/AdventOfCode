// Package day2 has solutions for Day 2 of Advent of Code
// https://adventofcode.com/2020/day/2
package day2

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"../input"
)

var inputFile string = "day2/input.txt"
var testFile string = "day2/testInput.txt"

type password struct {
	min          int
	max          int
	policyLetter byte
	password     string
}

/*
parseInput parses each line of the input into a password struct.
Every line of the input is of a form similar to '1-3 a: abcde' where '1-3'
refer to the lowest and highest number of times a given letter must appear in the password
'a' is the letter being referenced, and 'abcde' is the password
*/
func parseInput(file string) []password {
	inputList := input.Slice(file)

	var passwords []password
	for _, value := range inputList {
		// Split the password into its constituent components
		splitOn := regexp.MustCompile(`[- ]`)
		split := splitOn.Split(value, -1)

		min, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal("Failed to convert min to num", err)
		}

		max, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal("Failed to convert max to num", err)
		}

		pass := password{
			min:          min,
			max:          max,
			policyLetter: split[2][0],
			password:     split[3]}

		passwords = append(passwords, pass)
	}

	return passwords
}

/*
PartOne finds how many passwords are valid according to their policies
Policy: Each line gives the password policy and then the password.
The password policy indicates the lowest and highest number of times a
given letter must appear for the password to be valid.
Runs in O(n) time
*/
func PartOne() int {
	passwords := parseInput(inputFile)

	numValid := 0
	for _, pass := range passwords {
		count := strings.Count(pass.password, string(pass.policyLetter))
		if count >= pass.min && count <= pass.max {
			numValid++
		}
	}

	return numValid
}

/*
PartTwo finds how many passwords are valid according to their policies
Policy: Each line gives the password policy and then the password.
The password policy indicates two positions in the password, where 1 means the
first character, 2 means the second character, and so on. Exactly one of these positions
must contain the given letter. Other occurrences of the letter are irrelevant.
Runs in O(n) time
*/
func PartTwo() int {
	passwords := parseInput(inputFile)

	numValid := 0
	for _, pass := range passwords {
		first := pass.password[pass.min-1]
		second := pass.password[pass.max-1]

		if (first != second) && (first == pass.policyLetter || second == pass.policyLetter) {
			numValid++
		}
	}

	return numValid
}
