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

// rule represents a rule for how many of a certain color of bag must be contained
// within a larger bag.
type rule struct {
	color string
	num   int
}

// parseRules parses an input file containing rules for how different colored bags are nested
// a line within the input file is formatted similar to:
// "muted yellow bags contain 2 shiny gold bags, 9 faded blue bags."
//
// The function outputs a map of bag color to rules for sub bag colors.
func parseRules(file string) map[string][]rule {
	input := input.Slice(file)
	rules := make(map[string][]rule, len(input))
	for _, line := range input {
		// First item should be the color of the outer bag, second item should be the
		// info on the color and number of bags that can fit within the outer bag.
		splitLine := strings.Split(line, "contain")

		var currentRules []rule
		currentRule := rule{}
		for _, word := range strings.Fields(splitLine[1]) {
			if strings.Contains(word, "bag") { // We've reached the end of a sub-rule
				currentRule.color = strings.TrimSpace(currentRule.color)
				currentRules = append(currentRules, currentRule)
				currentRule = rule{}
				continue
			} else {
				match, _ := regexp.MatchString(`[0-9]+`, word)
				if match {
					num, _ := strconv.Atoi(word)
					currentRule.num = num
				} else {
					currentRule.color += word + " "
				}
			}
		}
		bagColor := strings.Split(splitLine[0], " bags")
		rules[bagColor[0]] = currentRules
	}
	return rules
}

// PartOne finds how many bag colors can eventually contain at least one shiny gold bag
func PartOne(file string) int {
	rules := parseRules(file)
	possible := make(map[string]bool)
	checkBags(rules, "shiny gold", &possible)

	return len(possible)
}

// checkBags determines how many bags can contain the target color of bag based on
// the set of rules passed in for how bags can nest. The set of bags that can contain
// the target bag is returned within possible
func checkBags(rules map[string][]rule, targetColor string, possible *map[string]bool) {
	for color := range rules {
		for _, rule := range rules[color] {
			if rule.color == targetColor {
				(*possible)[color] = true
				checkBags(rules, color, possible)
			}
		}
	}
}

// PartTwo finds how many individual bags are required inside a single shiny gold bag
func PartTwo(file string) int {
	rules := parseRules(file)
	return countSubBags(&rules, "shiny gold")
}

// countSubBags determines how many individual bags are contained within the target color of bag
// based on the set of rules passed in for how bags are nested
func countSubBags(rules *map[string][]rule, target string) int {
	count := 0
	for _, rule := range (*rules)[target] {
		count += rule.num * (1 + countSubBags(rules, rule.color))
	}

	return count
}
