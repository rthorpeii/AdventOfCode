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

	turn, new := len(nums), true
	for turn < turns {
		if new {
			nums[lastNum] = turn - 1
			new = false
			lastNum = 0
		} else {
			nextNum := turn - 1 - nums[lastNum]
			nums[lastNum] = turn - 1
			_, exists := nums[nextNum]
			new = !exists
			lastNum = nextNum
		}
		turn++
	}

	return lastNum
}
