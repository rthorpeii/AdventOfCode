// Package day25 has solutions for Day 25 of Advent of Code
// https://adventofcode.com/2020/day/25
package day25

import (
	"fmt"
	"strconv"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var puzzleInputFile string = "day25/input.txt"
var testFile string = "day25/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(puzzleInputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(puzzleInputFile))
}

// PartOne finds
func PartOne(file string) int {
	rawInput := input.Slice(file)

	cardPubKey, _ := strconv.Atoi(rawInput[0])
	doorPubKey, _ := strconv.Atoi(rawInput[1])

	cardLoopSize := findLoopSize(7, cardPubKey)
	doorLoopSize := findLoopSize(7, doorPubKey)

	fmt.Println(cardLoopSize, doorLoopSize)

	return transform(cardLoopSize, doorPubKey)
}

// PartTwo finds
func PartTwo(file string) int {
	// rawInput := input.ReadInput(file string)

	return -1
}

func findLoopSize(subject int, target int) int {
	loopSize := 0
	value := 1
	for value != target {
		value *= subject
		value %= 20201227
		loopSize++
	}

	return loopSize
}

func transform(loopSize int, subject int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value *= subject
		value %= 20201227
	}
	return value
}
