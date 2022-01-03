// Package day5 has solutions for Day 5 of Advent of Code
// https://adventofcode.com/2020/day/5
package day5

import (
	"fmt"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var inputFile string = "day5/input.txt"
var testFile string = "day5/testInput.txt"

// SolvePuzzle runs the inputs
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(inputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(inputFile))
}

// findIDs converts the seat string to its id. Each seat is written
// as a binary partition of the plane e.g. "FBFBBFFRLR"
// where F means "front", B means "back", L means "left", and R means "right".
func findIDs(seats *[]string) map[int]bool {
	ids := make(map[int]bool)
	for _, seat := range *seats {
		minRow, maxRow := 0, 128
		minCol, maxCol := 0, 8

		for _, char := range seat {
			value := string(char)
			switch value {
			case "F":
				maxRow -= ((maxRow - minRow) / 2)
			case "B":
				minRow += ((maxRow - minRow) / 2)
			case "L":
				maxCol -= ((maxCol - minCol) / 2)
			case "R":
				minCol += ((maxCol - minCol) / 2)
			}
		}
		id := minRow*8 + minCol
		ids[id] = true
	}
	return ids
}

// PartOne finds the max id in the input of seats
func PartOne(file string) int {
	rawInput := input.Slice(file)
	ids := findIDs(&rawInput)

	maxID := -1
	for id := range ids {
		if id > maxID {
			maxID = id
		}
	}

	return maxID
}

// PartTwo finds which id is missing from the file, but has both neighboring ids
func PartTwo(file string) int {
	rawInput := input.Slice(file)
	ids := findIDs(&rawInput)

	for key := range ids {
		if !ids[key+1] && ids[key+2] {
			return key + 1
		} else if !ids[key-1] && ids[key-2] {
			return key - 1
		}
	}
	return -1
}
