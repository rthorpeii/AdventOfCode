// Package day7 has solutions for Day 7 of Advent of Code
// https://adventofcode.com/2020/day/7
package day7

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var inputFile string = "day7/input.txt"
var testFile string = "day7/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(inputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(inputFile))
}

type rule struct {
	color string
	num   int
}

// PartOne finds
func PartOne(file string) int {
	rawInput := input.Slice(file)

	contains := make(map[string][]rule)
	for _, line := range rawInput {
		bags := strings.Split(line, "contain")

		bagColor := strings.Split(bags[0], " bags")
		currentRule := rule{}
		var rules []rule
		for _, word := range strings.Fields(bags[1]) {
			if strings.Contains(word, "bag") {
				rules = append(rules, currentRule)
				currentRule = rule{}
				continue
			} else {
				match, _ := regexp.MatchString(`[0-9]+`, word)
				if match {
					num, _ := strconv.Atoi(word)
					currentRule.num = num
				} else {
					if currentRule.color == "" {
						currentRule.color += word
					} else {
						currentRule.color += " " + word
					}

				}
			}

		}
		contains[bagColor[0]] = rules
	}

	possible := make(map[string]bool)
	checkBags(contains, "shiny gold", &possible)

	return len(possible)
}

func checkBags(rules map[string][]rule, target string, possible *map[string]bool) {
	for color := range rules {
		for _, rule := range rules[color] {
			if rule.color == target {
				(*possible)[color] = true
				checkBags(rules, color, possible)
			}
		}
	}
}

// PartTwo finds
func PartTwo(file string) int {
	rawInput := input.Slice(file)

	contains := make(map[string][]rule)
	for _, line := range rawInput {
		bags := strings.Split(line, "contain")

		bagColor := strings.Split(bags[0], " bags")
		currentRule := rule{}
		var rules []rule
		for _, word := range strings.Fields(bags[1]) {
			if strings.Contains(word, "bag") {
				rules = append(rules, currentRule)
				currentRule = rule{}
				continue
			} else {
				match, _ := regexp.MatchString(`[0-9]+`, word)
				if match {
					num, _ := strconv.Atoi(word)
					currentRule.num = num
				} else {
					if currentRule.color == "" {
						currentRule.color += word
					} else {
						currentRule.color += " " + word
					}

				}
			}

		}
		contains[bagColor[0]] = rules
	}

	// fmt.Println(contains)

	return countSubBags(contains, "shiny gold")
}

func countSubBags(rules map[string][]rule, target string) int {
	count := 0
	if target == "no other" {
		return count
	}

	for _, rule := range rules[target] {
		count += rule.num
		count += rule.num * countSubBags(rules, rule.color)
	}

	return count
}
