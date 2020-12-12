// Package day12 has solutions for Day 12 of Advent of Code
// https://adventofcode.com/2020/day/12
package day12

import (
	"fmt"
	"strconv"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var puzzleInputFile string = "day12/input.txt"
var testFile string = "day12/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(puzzleInputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(puzzleInputFile))
}

type coord struct {
	facing int
	x      int
	y      int
}

// PartOne finds
func PartOne(file string) int {
	input := input.Slice(file)
	coords := coord{90, 0, 0}
	for _, instr := range input {
		value, _ := strconv.Atoi(instr[1:])
		switch string(instr[0]) {
		case "N":
			coords.y += value
		case "S":
			coords.y -= value
		case "E":
			coords.x += value
		case "W":
			coords.x -= value
		case "L":
			coords.facing -= value
			coords.facing %= 360
		case "R":
			coords.facing += value
			coords.facing %= 360
		case "F":
			for coords.facing < 0 {
				coords.facing += 360
			}
			switch coords.facing {
			case 0:
				coords.y += value
			case 90:
				coords.x += value
			case 180:
				coords.y -= value
			case 270:
				coords.x -= value
			}
		}
	}

	return Abs(coords.x) + Abs(coords.y)
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// PartTwo finds
func PartTwo(file string) int {
	input := input.Slice(file)
	shipCoords := coord{90, 0, 0}
	waypointCoords := coord{0, 10, 1}
	for _, instr := range input {
		value, _ := strconv.Atoi(instr[1:])
		switch string(instr[0]) {
		case "N":
			waypointCoords.y += value
		case "S":
			waypointCoords.y -= value
		case "E":
			waypointCoords.x += value
		case "W":
			waypointCoords.x -= value
		case "L":
			switch value {
			case 90:
				newY := waypointCoords.x
				waypointCoords.x = -waypointCoords.y
				waypointCoords.y = newY
			case 180:
				waypointCoords.x = -waypointCoords.x
				waypointCoords.y = -waypointCoords.y
			case 270:
				newY := -waypointCoords.x
				waypointCoords.x = waypointCoords.y
				waypointCoords.y = newY
			}
		case "R":
			switch value {
			case 90:
				newY := -waypointCoords.x
				waypointCoords.x = waypointCoords.y
				waypointCoords.y = newY
			case 180:
				waypointCoords.x = -waypointCoords.x
				waypointCoords.y = -waypointCoords.y
			case 270:
				newY := waypointCoords.x
				waypointCoords.x = -waypointCoords.y
				waypointCoords.y = newY
			}
		case "F":
			for i := 0; i < value; i++ {
				shipCoords.x += waypointCoords.x
				shipCoords.y += waypointCoords.y
			}
		}
	}

	return Abs(shipCoords.x) + Abs(shipCoords.y)
}
