// Package day16 has solutions for Day 16 of Advent of Code
// https://adventofcode.com/2020/day/16
package day16

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var puzzleInputFile string = "day16/input.txt"
var testFile string = "day16/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(puzzleInputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(puzzleInputFile))
}

type rulePair struct {
	low  int
	high int
}

// PartOne finds
func PartOne(file string) int {
	rawInput := input.ReadInput(file)
	parts := strings.Split(rawInput, "\n\n")
	rawRules := parts[0]
	// myticket := parts[1]
	nearbyTickets := strings.Split(parts[2], "\n")

	rules := make(map[string][]rulePair)
	for _, line := range strings.Split(rawRules, "\n") {
		ruleParts := strings.Split(line, ": ")
		name := ruleParts[0]
		ruleVals := strings.Split(ruleParts[1], " or ")

		var pairs []rulePair
		for _, rule := range ruleVals {
			nums := strings.Split(rule, "-")
			low, _ := strconv.Atoi(nums[0])
			high, _ := strconv.Atoi(nums[1])
			currentPair := rulePair{low, high}
			pairs = append(pairs, currentPair)
		}

		rules[name] = pairs
	}

	errorRate := 0
	rulePos := make(map[int]map[string]bool)
	for i := 1; i < len(nearbyTickets); i++ {
		_, error := validateTicket(&rules, nearbyTickets[i], &rulePos)
		errorRate += error
	}

	return errorRate
}

func validateTicket(rules *map[string][]rulePair, ticket string, rulePos *map[int]map[string]bool) (bool, int) {
	errorRate := 0
	ruleUpdates := make(map[int]map[string]bool)
	ticketFields := strings.Split(ticket, ",")
	for index, num := range ticketFields {
		value, _ := strconv.Atoi(num)
		valid := false
		possibleRules := make(map[string]bool)
	RuleCheck:
		for name, rule := range *rules {
			for _, pair := range rule {
				if pair.low <= value && value <= pair.high {
					valid = true
					possibleRules[name] = true
					continue RuleCheck
				}
			}
		}
		if !valid {
			errorRate += value
		} else {
			ruleUpdates[index] = possibleRules
		}
	}
	if errorRate == 0 {
		for index := 0; index < len(ticketFields); index++ {
			_, exists := (*rulePos)[index]
			if !exists {
				(*rulePos)[index] = ruleUpdates[index]
			} else {
				for rule := range (*rulePos)[index] {
					if !(ruleUpdates[index][rule]) {
						delete((*rulePos)[index], rule)
					}
				}
			}
		}

		return true, errorRate
	}
	return false, errorRate
}

// PartTwo finds
func PartTwo(file string) int {
	rawInput := input.ReadInput(file)
	parts := strings.Split(rawInput, "\n\n")
	rawRules := parts[0]
	myTicket := strings.Split(parts[1], "\n")
	nearbyTickets := strings.Split(parts[2], "\n")

	rules := make(map[string][]rulePair)
	for _, line := range strings.Split(rawRules, "\n") {
		ruleParts := strings.Split(line, ": ")
		name := ruleParts[0]
		ruleVals := strings.Split(ruleParts[1], " or ")

		var pairs []rulePair
		for _, rule := range ruleVals {
			nums := strings.Split(rule, "-")
			low, _ := strconv.Atoi(nums[0])
			high, _ := strconv.Atoi(nums[1])
			currentPair := rulePair{low, high}
			pairs = append(pairs, currentPair)
		}

		rules[name] = pairs
	}

	var validTickets []string
	rulePos := make(map[int]map[string]bool)
	for i := 1; i < len(nearbyTickets); i++ {
		valid, _ := validateTicket(&rules, nearbyTickets[i], &rulePos)
		if valid {
			validTickets = append(validTickets, nearbyTickets[i])
		}
	}

	myticketFields := strings.Split(myTicket[1], ",")
	finalPos := make(map[string]int)
	for len(finalPos) < len(myticketFields) {
	OuterLoop:
		for index, rules := range rulePos {
			for name := range rules {
				if len(rules) == 1 {
					finalPos[name] = index
					delete(rulePos, index)
					break OuterLoop
				}
				_, valid := finalPos[name]
				if valid {
					delete(rules, name)
					continue
				}
			}

		}
	}

	// fmt.Println(finalPos)
	answer := 1
	for name, index := range finalPos {
		matches, _ := regexp.MatchString(`^departure`, name)
		if matches {
			num, _ := strconv.Atoi(myticketFields[index])
			answer *= num
		}
	}

	return answer
}
