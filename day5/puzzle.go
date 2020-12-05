// Package day5 has solutions for Day 5 of Advent of Code
// https://adventofcode.com/2020/day/5
package day5

import (
	"fmt"

	"../input"
)

var inputFile string = "day5/input.txt"
var testFile string = "day5/testInput.txt"

// PartOne finds
func PartOne() int {
	rawInput := input.Slice(inputFile)

	max := 0
	for _, seat := range rawInput {
		minRow := 0
		maxRow := 127
		minCol := 0
		maxCol := 7
		for _, char := range seat {
			value := string(char)
			fmt.Println(value)
			if value == "F" {
				maxRow -= ((maxRow - minRow + 1) / 2)
			} else if value == "B" {
				minRow += ((maxRow - minRow + 1) / 2)
			} else if value == "L" {
				maxCol -= ((maxCol - minCol + 1) / 2)
			} else if value == "R" {
				minCol += ((maxCol - minCol + 1) / 2)
			}
		}
		id := minRow*8 + minCol
		fmt.Println(minRow)
		fmt.Println(minCol)
		if id > max {
			max = id
		}
	}
	return max
}

// PartTwo finds
func PartTwo() int {
	rawInput := input.Slice(inputFile)

	seats := make(map[int]bool)
	for _, seat := range rawInput {
		minRow := 0
		maxRow := 127
		minCol := 0
		maxCol := 7
		for _, char := range seat {
			value := string(char)
			if value == "F" {
				maxRow -= ((maxRow - minRow + 1) / 2)
			} else if value == "B" {
				minRow += ((maxRow - minRow + 1) / 2)
			} else if value == "L" {
				maxCol -= ((maxCol - minCol + 1) / 2)
			} else if value == "R" {
				minCol += ((maxCol - minCol + 1) / 2)
			}
		}
		id := minRow*8 + minCol
		seats[id] = true
	}
	for key := range seats {
		if seats[key+2] && !seats[key+1] {
			return key + 1
		} else if seats[key-2] && !seats[key-1] {
			return key - 1
		}
	}
	return -1
}
