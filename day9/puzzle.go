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

// PartOne finds
func PartOne(file string) int {
	rawInput := input.IntSlice(file)
	preamble := 25

	return findTarget(rawInput, preamble)
}

func findTarget(input []int, preamble int) int {
	for i := preamble; i < len(input); i++ {
		if !valid(input[i-preamble:i], input[i]) {
			return input[i]
		}
	}
	return -1
}

func valid(input []int, target int) bool {
	for index, val := range input {
		for j := index + 1; j < len(input); j++ {
			if val+input[j] == target {
				return true
			}
		}
	}

	return false
}

// PartTwo finds
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
	for i := min; i < max; i++ {
		val := rawInput[i]
		if val < smallest {
			smallest = val
		} else if val > largest {
			largest = val
		}
	}

	return smallest + largest
}
