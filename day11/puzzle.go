// Package day11 has solutions for Day 11 of Advent of Code
// https://adventofcode.com/2020/day/11
package day11

import (
	"fmt"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var puzzleInputFile string = "day11/input.txt"
var testFile string = "day11/testInput.txt"

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
	for i := range input {
		input[i] = "." + input[i] + "."
	}
	floorString := ""
	for i := 0; i < len(input[0]); i++ {
		floorString += "."
	}
	floorSlice := make([]string, 0)
	floorSlice = append(floorSlice, floorString)
	input = append(floorSlice, input...)
	input = append(input, floorString)
	changed := true
	nextEmpty := 0
	for changed {
		nextEmpty = 0
		changed = false
		nextInput := make([]string, len(input))
		for y := 0; y < len(input); y++ {

			for x := 0; x < len(input[y]); x++ {
				switch string(input[y][x]) {
				case "L":
					nextInput[y] += emptyChange(input, x, y)
				case "#":
					nextInput[y] += occupiedChange(input, x, y)
				case ".":
					nextInput[y] += "."
				}
				if nextInput[y][x] != input[y][x] {
					changed = true
				}
				if string(nextInput[y][x]) == "#" {
					nextEmpty++
				}
			}
		}
		input = nextInput
	}
	return nextEmpty
}
func emptyChange(seats []string, x1 int, y1 int) string {
	for y := y1 - 1; y <= y1+1; y++ {
		for x := x1 - 1; x <= x1+1; x++ {
			if y1 == y && x1 == x {
				continue
			}
			if string(seats[y][x]) == "#" {
				return "L"
			}
		}
	}
	return "#"
}

func occupiedChange(seats []string, x1 int, y1 int) string {
	occupied := 0
	for y := y1 - 1; y <= y1+1; y++ {
		for x := x1 - 1; x <= x1+1; x++ {
			if y1 == y && x1 == x {
				continue
			}
			if string(seats[y][x]) == "#" {
				occupied++
			}
		}
	}
	if occupied >= 4 {
		return "L"
	}
	return "#"
}

// PartTwo finds
func PartTwo(file string) int {
	input := input.Slice(file)
	for i := range input {
		input[i] = "." + input[i] + "."
	}
	floorString := ""
	for i := 0; i < len(input[0]); i++ {
		floorString += "."
	}
	floorSlice := make([]string, 0)
	floorSlice = append(floorSlice, floorString)
	input = append(floorSlice, input...)
	input = append(input, floorString)
	changed := true
	nextEmpty := 0
	for changed {
		nextEmpty = 0
		changed = false
		nextInput := make([]string, len(input))
		for y := 0; y < len(input); y++ {

			for x := 0; x < len(input[y]); x++ {
				switch string(input[y][x]) {
				case "L":
					nextInput[y] += emptySightChange(input, x, y)
				case "#":
					nextInput[y] += occupiedSightChange(input, x, y)
				case ".":
					nextInput[y] += "."
				}
				if nextInput[y][x] != input[y][x] {
					changed = true
				}
				if string(nextInput[y][x]) == "#" {
					nextEmpty++
				}
			}
		}
		input = nextInput
	}
	return nextEmpty
}

func emptySightChange(seats []string, x1 int, y1 int) string {
	for delta := 1; delta <= y1; delta++ {
		y := y1 - delta
		if string(seats[y][x1]) == "#" { // straight up
			return "L"
		} else if string(seats[y][x1]) == "L" {
			break
		}
	}
	for delta := 1; delta <= y1; delta++ {
		if x1-delta > -1 {
			y := y1 - delta
			if string(seats[y][x1-delta]) == "#" { // diagonal up left
				return "L"
			} else if string(seats[y][x1-delta]) == "L" {
				break
			}
		}
	}
	for delta := 1; delta <= y1; delta++ {
		y := y1 - delta
		if x1+delta < len(seats[y]) {
			y := y1 - delta
			if string(seats[y][x1+delta]) == "#" { // diagonal up right
				return "L"
			} else if string(seats[y][x1+delta]) == "L" {
				break
			}
		}
	}
	for delta := -1; delta > y1-len(seats); delta-- {
		y := y1 - delta
		if string(seats[y][x1]) == "#" { // straight down
			return "L"
		} else if string(seats[y][x1]) == "L" {
			break
		}
	}
	for delta := -1; delta > y1-len(seats); delta-- {
		y := y1 - delta
		if x1-delta < len(seats[y]) {
			if string(seats[y][x1-delta]) == "#" { // diagonal right
				return "L"
			} else if string(seats[y][x1-delta]) == "L" {
				break
			}
		}
	}
	for delta := -1; delta > y1-len(seats); delta-- {
		y := y1 - delta
		if x1+delta > -1 {
			if string(seats[y][x1+delta]) == "#" { // diagonal left
				return "L"
			} else if string(seats[y][x1+delta]) == "L" {
				break
			}
		}
	}
	for delta := 1; delta < len(seats[0])-x1; delta++ {
		x := x1 + delta
		if string(seats[y1][x]) == "#" { // straight right
			return "L"
		} else if string(seats[y1][x]) == "L" {
			break
		}
	}
	for delta := 1; delta <= x1; delta++ {
		x := x1 - delta
		if string(seats[y1][x]) == "#" { // straight left
			return "L"
		} else if string(seats[y1][x]) == "L" {
			break
		}
	}
	return "#"
}

func occupiedSightChange(seats []string, x1 int, y1 int) string {
	occupied := 0
	for delta := 1; delta <= y1; delta++ {
		y := y1 - delta
		if string(seats[y][x1]) == "#" { // straight up
			occupied++
			break
		} else if string(seats[y][x1]) == "L" {
			break
		}
	}
	for delta := 1; delta <= y1; delta++ {
		y := y1 - delta
		if x1-delta > -1 {
			if string(seats[y][x1-delta]) == "#" { // diagonal up left
				occupied++
				break
			} else if string(seats[y][x1-delta]) == "L" {
				break
			}
		}
	}
	for delta := 1; delta <= y1; delta++ {
		y := y1 - delta
		if x1+delta < len(seats[y]) {
			if string(seats[y][x1+delta]) == "#" { // diagonal up right
				occupied++
				break
			} else if string(seats[y][x1+delta]) == "L" {
				break
			}
		}
	}
	for delta := -1; delta > y1-len(seats); delta-- {
		y := y1 - delta
		if string(seats[y][x1]) == "#" { // straight down
			occupied++
			break
		} else if string(seats[y][x1]) == "L" {
			break
		}
	}
	for delta := -1; delta > y1-len(seats); delta-- {
		y := y1 - delta
		if x1-delta < len(seats[y]) {
			if string(seats[y][x1-delta]) == "#" { // diagonal right
				occupied++
				break
			} else if string(seats[y][x1-delta]) == "L" {
				break
			}
		}
	}
	for delta := -1; delta > y1-len(seats); delta-- {
		y := y1 - delta
		if x1+delta > -1 {
			if string(seats[y][x1+delta]) == "#" { // diagonal left
				occupied++
				break
			} else if string(seats[y][x1+delta]) == "L" {
				break
			}
		}
	}
	for delta := 1; delta < len(seats[0])-x1; delta++ {
		x := x1 + delta
		if string(seats[y1][x]) == "#" { // straight right
			occupied++
			break
		} else if string(seats[y1][x]) == "L" {
			break
		}
	}
	for delta := 1; delta <= x1; delta++ {
		x := x1 - delta
		if string(seats[y1][x]) == "#" { // straight left
			occupied++
			break
		} else if string(seats[y1][x]) == "L" {
			break
		}
	}

	if occupied >= 5 {
		return "L"
	}
	return "#"
}
