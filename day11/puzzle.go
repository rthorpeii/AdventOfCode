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

type coords struct {
	x, y int
}

// PartOne finds
func PartOne(file string) int {
	input := input.Slice(file)

	changed := true
	empytSeats := 0
	for changed {
		empytSeats = 0
		changed = false
		nextInput := make([]string, len(input))
		for y := 0; y < len(input); y++ {
			for x := 0; x < len(input[y]); x++ {
				pos := coords{x, y}
				value := string(input[y][x])
				switch value {
				case "L":
					nextInput[y] += emptyChange(&input, pos)
				case "#":
					nextInput[y] += occupiedChange(&input, pos)
				case ".":
					nextInput[y] += "."
				}
				if nextInput[y][x] != input[y][x] {
					changed = true
				}
				if string(nextInput[y][x]) == "#" {
					empytSeats++
				}
			}
		}
		input = nextInput
	}
	return empytSeats
}

func validateCoords(seats *[]string, pos coords) bool {
	if pos.y >= len((*seats)) || pos.y < 0 {
		return false
	}
	return (pos.x > -1 && pos.x < len((*seats)[pos.y]))
}

func emptyChange(seats *[]string, pos coords) string {
	for y := -1; y <= 1; y++ {
		for x := -1; x <= +1; x++ {
			target := coords{pos.x + x, pos.y + y}
			if !validateCoords(seats, target) || (target == pos) {
				continue
			}
			if string((*seats)[target.y][target.x]) == "#" {
				return "L"
			}
		}
	}
	return "#"
}

func occupiedChange(seats *[]string, pos coords) string {
	occupied := 0
	for y := -1; y <= 1; y++ {
		for x := -1; x <= +1; x++ {
			target := coords{pos.x + x, pos.y + y}
			if !validateCoords(seats, target) || (target == pos) {
				continue
			}
			if string((*seats)[target.y][target.x]) == "#" {
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

	changed := true
	nextEmpty := 0
	for changed {
		nextEmpty = 0
		changed = false
		nextInput := make([]string, len(input))
		for y := 0; y < len(input); y++ {
			for x := 0; x < len(input[y]); x++ {
				pos := coords{x, y}
				value := string(input[y][x])
				switch value {
				case "L":
					nextInput[y] += emptySightChange(&input, pos)
				case "#":
					nextInput[y] += occupiedSightChange(&input, pos)
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

func emptySightChange(seats *[]string, pos coords) string {
	for y := -1; y <= 1; y++ {
		for x := -1; x <= +1; x++ {
			target := coords{pos.x, pos.y}
		LineOfSight:
			for true {
				target.x += x
				target.y += y
				if !validateCoords(seats, target) || (target == pos) {
					break
				}
				switch string((*seats)[target.y][target.x]) {
				case "#":
					return "L"
				case "L":
					break LineOfSight
				}
			}

		}
	}
	return "#"
}

func occupiedSightChange(seats *[]string, pos coords) string {
	occupied := 0
	for y := -1; y <= 1; y++ {
		for x := -1; x <= +1; x++ {
			target := coords{pos.x, pos.y}
		LineOfSight:
			for true {
				target.x += x
				target.y += y
				if !validateCoords(seats, target) || (target == pos) {
					break
				}
				switch string((*seats)[target.y][target.x]) {
				case "#":
					occupied++
					break LineOfSight
				case "L":
					break LineOfSight
				}
			}
		}
	}
	if occupied >= 5 {
		return "L"
	}
	return "#"
}
