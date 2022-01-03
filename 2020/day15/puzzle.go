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

// PartOne finds the number said on the 2020th turn
func PartOne(file string) int {
	return playGame(file, 2020)
}

// PartTwo finds the number said on the 30000000th turn
func PartTwo(file string) int {
	return playGame(file, 30000000)
}

// playGame runs the rules of the game for the specified number of
// turns, starting with the numbers in the input file
func playGame(file string, turns int) int {
	starters := strings.Split(input.ReadInput(file), ",")

	nums, lastNum := make(map[int]int), 0
	for index, value := range starters {
		num, _ := strconv.Atoi(value)
		nums[num] = index
		lastNum = num
	}

	target := turns - 1
	turn, seenBefore := len(nums)-1, true
	for turn < target {
		if seenBefore {
			nextNum := turn - nums[lastNum]
			nums[lastNum] = turn
			_, exists := nums[nextNum]
			seenBefore = exists
			lastNum = nextNum
		} else {
			nums[lastNum] = turn
			seenBefore = true
			lastNum = 0
		}
		turn++
	}

	return lastNum
}
