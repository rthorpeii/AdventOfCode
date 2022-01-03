// Package day23 has solutions for Day 23 of Advent of Code
// https://adventofcode.com/2020/day/23
package day23

import (
	"fmt"
	"strconv"

	"github.com/rthorpeii/AdventOfCode2020/helper"
	"github.com/rthorpeii/AdventOfCode2020/input"
)

var puzzleInputFile string = "day23/input.txt"
var testFile string = "day23/testInput.txt"

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

	cups := make(map[int]int, len(rawInput))

	prevNumber, _ := strconv.Atoi(string(rawInput[0]))
	for i := 1; i < len(rawInput); i++ {
		currentNumber, _ := strconv.Atoi(string(rawInput[i]))
		cups[prevNumber] = currentNumber
		prevNumber = currentNumber
	}
	firstNumber, _ := strconv.Atoi(string(rawInput[0]))
	cups[prevNumber] = firstNumber
	fmt.Println(cups)

	currentCup := firstNumber
	for round := 0; round < 100; round++ {
		// fmt.Printf("-- move %v --\n", round+1)
		// fmt.Printf("cups: %v\n", cups)
		removed := pullCups(&cups, currentCup)
		// fmt.Printf("pick up: %v\n", removed)

		insertCups(&cups, &removed, currentCup)
		currentCup = cups[currentCup]
	}

	start := 1
	cupString := ""
	for i := 0; i < len(cups)-1; i++ {
		cupString += strconv.Itoa(cups[start])
		start = cups[start]
	}
	fmt.Println(cupString)

	return -1
}

func insertCups(cups *map[int]int, removed *[]int, currentVal int) {
	max := len(*cups)
	target := helper.Mod(currentVal-1, max)
	for target == (*removed)[0] || target == (*removed)[1] || target == (*removed)[2] {
		target = helper.Mod(target-1, max)
	}
	// fmt.Println("Destination: ", target)
	postTarget := (*cups)[target]
	(*cups)[target] = (*removed)[0]
	(*cups)[(*removed)[2]] = postTarget
}

func pullCups(cups *map[int]int, current int) []int {
	numCups := 3

	removed := make([]int, numCups)

	next := current
	for i := 0; i < 3; i++ {
		removed[i] = (*cups)[next]
		next = removed[i]
	}

	(*cups)[current] = (*cups)[removed[2]]

	return removed
}

// PartTwo finds
func PartTwo(file string) int {
	rawInput := input.ReadInput(file)

	cups := make(map[int]int, len(rawInput))

	prevNumber, _ := strconv.Atoi(string(rawInput[0]))
	for i := 1; i < len(rawInput); i++ {
		currentNumber, _ := strconv.Atoi(string(rawInput[i]))
		cups[prevNumber] = currentNumber
		prevNumber = currentNumber
	}
	fmt.Println(len(cups))
	for i := len(cups) + 2; i <= 1000000; i++ {
		cups[prevNumber] = i
		prevNumber = i
	}

	firstNumber, _ := strconv.Atoi(string(rawInput[0]))
	cups[prevNumber] = firstNumber
	// fmt.Println(cups)

	currentCup := firstNumber
	for round := 0; round < 10000000; round++ {
		// fmt.Printf("-- move %v --\n", round+1)
		// fmt.Printf("cups: %v\n", cups)
		removed := pullCups(&cups, currentCup)
		// fmt.Printf("pick up: %v\n", removed)

		insertCups(&cups, &removed, currentCup)
		currentCup = cups[currentCup]
	}

	fmt.Println(cups[1], cups[cups[1]])
	return cups[1] * cups[cups[1]]
}
