// Package day10 has solutions for Day 10 of Advent of Code
// https://adventofcode.com/2020/day/10
package day10

import (
	"fmt"
	"sort"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var inputFile string = "day10/input.txt"
var testFile string = "day10/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(inputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(inputFile))
}

// PartOne finds the number of 1-jolt differences multiplied by the number of 3-jolt differences
func PartOne(file string) int {
	// Parse the input, sort it, and add the outlet, and phone values
	input := append(input.IntSlice(file), 0)
	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)

	countOne, countThree := 0, 0
	if input[0] == 1 {
		countOne++
	} else if input[0] == 3 {
		countThree++
	}
	for i := 1; i < len(input); i++ {
		if input[i]-input[i-1] == 1 {
			countOne++
		} else if input[i]-input[i-1] == 3 {
			countThree++
		}
	}
	return countOne * countThree
}

// PartTwo finds the total number of distinct ways you can
// arrange the adapters to connect the charging outlet to your device
func PartTwo(file string) int {
	input := append(input.IntSlice(file), 0)
	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)

	return countValid(input)
}

// Represents how many orderings there were with a particular variable, and without it
type orderingCount struct {
	with    int
	without int
}

// countValid determines the number of distinct ways you can
// arrange the input adapters to connect the charging outlet to your device
func countValid(input []int) int {
	doublePrevCount := orderingCount{1, 0} // How many orderings there were two numbers in the past
	prevCount := orderingCount{}           // How many orderings there were one number ago

	// Resolve the base cases
	if input[2]-input[1] <= 3 {
		prevCount = orderingCount{1, 1}
	} else {
		prevCount = orderingCount{1, 0}
	}

	for i := 2; i < len(input)-1; i++ {
		prevAdapter := input[i-1]
		doublePrevAdapter := input[i-2]
		currentAdapter := orderingCount{}
		// The number of valid iterations so far with the current adapter included is equal to
		// the sum of the number of iterations with and without the previous number in the sequence
		currentAdapter.with = prevCount.with + prevCount.without
		if input[i+1]-prevAdapter <= 3 { // If we can remove the current number
			currentAdapter.without += prevCount.with
		}
		if input[i+1]-doublePrevAdapter <= 3 { // If we can remove the current number and the previous number
			currentAdapter.without += prevCount.without - doublePrevCount.without
		}
		doublePrevCount = prevCount
		prevCount = currentAdapter
	}

	return prevCount.with
}
