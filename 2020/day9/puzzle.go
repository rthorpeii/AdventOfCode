// Package day9 has solutions for Day 9 of Advent of Code
// https://adventofcode.com/2020/day/9
package day9

import (
	"fmt"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var inputFile string = "day9/input.txt"
var testFile string = "day9/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(inputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(inputFile))
}

// What size should the preamble f
var preamble = 25

// PartOne finds what number in the input is not the sum of two of the previous 25 numbers
func PartOne(file string) int {
	input := input.IntSlice(file)
	return findTarget(input, 25)
}

// Finds what number in the input is not the sum of the previous x numbers
// where x is preamble
func findTarget(input []int, preamble int) int {
	for i := preamble; i < len(input); i++ {
		if !valid(input[i-preamble:i], input[i]) {
			return input[i]
		}
	}
	return -1
}

// valid determines if the target can be created by adding two numbers within the input
func valid(input []int, target int) bool {
	for first, val := range input {
		for second := first + 1; second < len(input); second++ {
			if val+input[second] == target {
				return true
			}
		}
	}

	return false
}

// PartTwo finds what consecutive sequence of at least 2 numbers in the input sum to the
// target number which is the number in the input that is not the sum of two of the previous 25 numbers
func PartTwo(file string) int {
	rawInput := input.IntSlice(file)
	preamble := 25
	target := findTarget(rawInput, preamble)

	min, max := 0, 1
	sum := rawInput[min]
	for max < len(rawInput)-1 {
		sum += rawInput[max]

		for sum > target {
			if max-min == 0 {
				break
			}
			sum -= rawInput[min]
			min++
		}

		if sum == target {
			break
		}

		max++
	}

	smallest, largest := rawInput[max], rawInput[max]
	for index := min; index < max; index++ {
		val := rawInput[index]
		if val < smallest {
			smallest = val
		} else if val > largest {
			largest = val
		}
	}

	return smallest + largest
}
