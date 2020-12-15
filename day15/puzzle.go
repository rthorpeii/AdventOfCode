// Package day15 has solutions for Day 15 of Advent of Code
// https://adventofcode.com/2020/day/15
package day15

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var puzzleInputFile string = "day15/input.txt"
var testFile string = "day15/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(puzzleInputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(puzzleInputFile))
}

// PartOne finds
func PartOne(file string) int {
	rawInput := input.ReadInput(file)
	starters := strings.Split(rawInput, ",")

	turn := 0
	nums := make(map[int]int)
	lastNum := 0
	for _, value := range starters {
		num, _ := strconv.Atoi(value)
		nums[num] = turn
		lastNum = num
		turn++
	}

	new := true
	// fmt.Println(nums)
	for turn < 30000000 {
		// fmt.Println(lastNum, nums, turn)
		nextNum := 0
		if new {
			nums[lastNum] = turn - 1
			nextNum = 0
			new = false
		} else {
			nextNum = turn - 1 - nums[lastNum]
			nums[lastNum] = turn - 1
			_, ok := nums[nextNum]
			new = !ok
		}
		lastNum = nextNum
		turn++
	}

	return lastNum
}

// PartTwo finds
func PartTwo(file string) int {
	// rawInput := input.ReadInput(file string)

	return -1
}
