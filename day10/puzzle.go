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

// PartOne finds
func PartOne(file string) int {
	input := input.IntSlice(file)
	sort.Ints(input)

	countOne := 0
	countThree := 0
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
	fmt.Printf("One: %v\n Three: %v\n", countOne, countThree)
	return countOne * (countThree + 1)
}

// PartTwo finds
func PartTwo(file string) int {
	input := input.IntSlice(file)
	input = append(input, 0)
	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)

	return countValid(input)
}

type valid struct {
	with    int
	without int
}

// 0,  1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19,  22
func countValid(input []int) int {
	validMap := make(map[int]valid, len(input))
	count := 0
	validMap[0] = valid{1, 0}
	if input[2]-input[1] <= 3 {
		validMap[1] = valid{1, 1}
	} else {
		validMap[1] = valid{1, 0}
	}
	for i := 2; i < len(input)-1; i++ {
		value := input[i]
		prev := input[i-1]
		doublePrev := input[i-2]
		current := valid{}
		current.with = validMap[prev].with + validMap[prev].without
		if input[i+1]-prev <= 3 {
			current.without += validMap[prev].with
		}
		if input[i+1]-doublePrev <= 3 {
			current.without += validMap[prev].without - validMap[doublePrev].without
		}
		validMap[value] = current
	}

	fmt.Println(validMap)

	return count

}
