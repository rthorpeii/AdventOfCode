// Package day6 has solutions for Day 6 of Advent of Code
// https://adventofcode.com/2020/day/6
package day6

import (
	"fmt"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var inputFile string = "day6/input.txt"
var testFile string = "day6/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(inputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(inputFile))
}

// PartOne finds
func PartOne(file string) int {
	rawInput := input.Slice(file)
	count := 0
	answer := make(map[string]bool)

	for _, line := range rawInput {
		if line == "" {
			count += len(answer)
			answer = make(map[string]bool)
		}
		for _, char := range line {
			answer[string(char)] = true
		}
	}

	return count
}

// PartTwo finds
func PartTwo(file string) int {
	rawInput := input.Slice(file)
	count := 0
	people := 0
	answer := make(map[string]int)

	for _, line := range rawInput {
		if line == "" {
			for val := range answer {
				if answer[val] == people {
					count++
				}
			}
			answer = make(map[string]int)
			people = 0
			continue
		}
		people++
		for _, char := range line {
			answer[string(char)]++
		}
	}

	return count
}
