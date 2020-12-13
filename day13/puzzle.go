// Package day13 has solutions for Day 13 of Advent of Code
// https://adventofcode.com/2020/day/13
package day13

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var puzzleInputFile string = "day13/input.txt"
var testFile string = "day13/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(puzzleInputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(puzzleInputFile))
}

// PartOne finds
func PartOne(file string) int {
	input := input.Slice(file)

	rawIDs := strings.Split(input[1], ",")
	busIDs := make(map[int]float64)
	for index, id := range rawIDs {
		if id == "x" {
			continue
		}
		id, _ := strconv.ParseFloat(id, 64)
		busIDs[index] = id
	}
	departTime, _ := strconv.ParseFloat(input[0], 64)

	minTime, busID := 100000000.0, -1.0
	for _, id := range busIDs {
		mult := math.Ceil(departTime / id)
		busDepart := id * mult
		if busDepart < minTime {
			minTime = busDepart
			busID = id
		}
	}
	return int((minTime - departTime) * busID)
}

// PartTwo finds
func PartTwo(file string) int {
	input := strings.Split(input.Slice(file)[1], ",")

	busIDs := make(map[int]int)
	for index, id := range input {
		if id == "x" {
			continue
		}
		id, _ := strconv.Atoi(id)
		busIDs[index] = id
	}
	delete(busIDs, 0)

	firstID, _ := strconv.Atoi(input[0])
	targetTime, timeIter := firstID, firstID
	for index, id := range busIDs {
		for true {
			if (targetTime+index)%id == 0 {
				timeIter *= id
				break
			}
			targetTime += timeIter
		}
	}

	return targetTime
}
