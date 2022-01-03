// Package day19 has solutions for Day 19 of Advent of Code
// https://adventofcode.com/2020/day/19
package day19

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var puzzleInputFile string = "day19/input.txt"
var testFile string = "day19/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	// fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	// fmt.Printf("Part 1 - Actual: %v \n", PartOne(puzzleInputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(puzzleInputFile))
}

// Subrule is a
type Subrule struct {
	rules []int
}

// Rule is a
type Rule struct {
	subrules  []Subrule
	character string
}

func generateRules(inputRules string) []Rule {
	rawRules := strings.Split(inputRules, "\n")

	finalRules := make([]Rule, len(rawRules)*2)
	for _, line := range rawRules {
		split := strings.Fields(line)

		currentRule := Rule{}

		curSubrule := Subrule{}
		for i := 1; i < len(split); i++ {
			value := split[i]
			if string(value[0]) == "\"" {
				currentRule.character = string(value[1])
			} else if value == "|" {
				currentRule.subrules = append(currentRule.subrules, curSubrule)
				curSubrule = Subrule{}
			} else {
				num, _ := strconv.Atoi(value)
				curSubrule.rules = append(curSubrule.rules, num)
			}
		}
		currentRule.subrules = append(currentRule.subrules, curSubrule)
		ruleNum, _ := strconv.Atoi(split[0][:len(split[0])-1])
		finalRules[ruleNum] = currentRule
	}

	return finalRules
}

var seen = make(map[int]map[string]bool)

func (rule Rule) generateMatches(rules *[]Rule, self int) map[string]bool {

	if len(seen[self]) != 0 {
		return seen[self]
	}

	mapping := make(map[string]bool)
	if rule.character != "" {
		mapping[rule.character] = true
		return mapping
	}

	for _, subrule := range rule.subrules {
		subRuleString := make(map[string]bool)
		for _, num := range subrule.rules {
			subRulesComb := (*rules)[num].generateMatches(rules, num)

			if len(subRuleString) == 0 {
				subRuleString = subRulesComb
			} else {
				newSubStrings := make(map[string]bool)
				for value1 := range subRuleString {
					for value2 := range subRulesComb {
						newSubStrings[value1+value2] = true
					}
				}
				subRuleString = newSubStrings
			}
		}
		for value := range subRuleString {
			mapping[value] = true
		}
	}
	seen[self] = mapping
	return mapping
}

// PartOne finds
func PartOne(file string) int {
	seen = make(map[int]map[string]bool)
	rawInput := strings.Split(input.ReadInput(file), "\n\n")
	rules := generateRules(rawInput[0])
	mapping := rules[0].generateMatches(&rules, 0)

	count := 0
	for _, value := range strings.Split(rawInput[1], "\n") {
		if mapping[value] {
			count++
		}
	}

	return count
}

// PartTwo finds
func PartTwo(file string) int {
	seen = make(map[int]map[string]bool)
	rawInput := strings.Split(input.ReadInput(file), "\n\n")
	rules := generateRules(rawInput[0])
	map42 := rules[42].generateMatches(&rules, 42)
	max42 := maxString(map42)
	map31 := rules[31].generateMatches(&rules, 31)
	max31 := maxString(map31)

	count := 0
	fmt.Println(map31)
	for _, value := range strings.Split(rawInput[1], "\n") {
		if messageValid(value, map42, max42, map31, max31, 0, false, 0) {
			count++
			fmt.Println("Matched: ", value)
		}
	}

	return count
}

func messageValid(substring string, map42 map[string]bool, max42 int, map31 map[string]bool, max31 int, valid42 int, started31 bool, valid31 int) bool {
	fmt.Println("Validating", substring, valid42, started31, valid31)
	if len(substring) == 0 {
		if valid31 > 0 && valid31 < valid42 {
			return true
		}
		return false
	}
	start := 0
	end := 2
	for true {
		if !started31 {
			if end-start > max42 || end == len(substring)+1 {
				if valid42 < 2 {
					return false
				}
				started31 = true
				end = start + 2
				continue
			} else {
				if map42[substring[start:end]] {
					if messageValid(substring[end:], map42, max42, map31, max31, valid42+1, started31, valid31) {
						return true
					}
				}
				end++
			}
		} else {
			if end-start > max31 || end == len(substring)+1 {
				return false
			}
			if map31[substring[start:end]] {
				if messageValid(substring[end:], map42, max42, map31, max31, valid42, started31, valid31+1) {
					return true
				}
			}
			end++
		}
	}
	return false
}

func maxString(mapping map[string]bool) int {
	maxValue := 0
	for value := range mapping {
		if len(value) > maxValue {
			maxValue = len(value)
		}
	}
	return maxValue
}
