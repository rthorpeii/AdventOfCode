// Package day12 has solutions for Day 12 of Advent of Code
// https://adventofcode.com/2020/day/12
package day12

import (
	"fmt"
	"strconv"

	"github.com/rthorpeii/AdventOfCode2020/helper"
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

// coords represents the position and facing of an object
type coords struct {
	facing int // How many degrees the object is rotated clockwise, 0 being north
	x      int // Position on the horizontal (East/West) axis
	y      int // Position on the vertical (North/South axis)
}

// PartOne finds the manhattan distance of the ship after following the instructions passed in
func PartOne(file string) int {
	input := input.Slice(file)
	ship := coords{90, 0, 0}
	for _, instr := range input {
		value, _ := strconv.Atoi(instr[1:])
		switch string(instr[0]) {
		case "N":
			ship.y += value
		case "S":
			ship.y -= value
		case "E":
			ship.x += value
		case "W":
			ship.x -= value
		case "L":
			ship.facing -= value
		case "R":
			ship.facing += value
		case "F":
			ship.facing %= 360
			if ship.facing < 0 {
				ship.facing += 360
			}
			switch ship.facing {
			case 0:
				ship.y += value
			case 90:
				ship.x += value
			case 180:
				ship.y -= value
			case 270:
				ship.x -= value
			}
		}
	}

	return helper.AbsInt(ship.x) + helper.AbsInt(ship.y)
}

// PartTwo finds the manhattan distance of the ship after following the instructions passed in
// This time however, most instructions now move a waypoint which the ship will advance towards
func PartTwo(file string) int {
	input := input.Slice(file)
	ship := coords{}
	waypoint := coords{0, 10, 1}
	for _, instr := range input {
		value, _ := strconv.Atoi(instr[1:])
		switch string(instr[0]) {
		case "N":
			waypoint.y += value
		case "S":
			waypoint.y -= value
		case "E":
			waypoint.x += value
		case "W":
			waypoint.x -= value
		case "L":
			rotateWaypoint(&waypoint, 360-value)
		case "R":
			rotateWaypoint(&waypoint, value)
		case "F":
			for i := 0; i < value; i++ {
				ship.x += waypoint.x
				ship.y += waypoint.y
			}
		}
	}

	return helper.AbsInt(ship.x) + helper.AbsInt(ship.y)
}

// Rotates the waypoint clockwise around the ship by a number of degrees
// Assumes degrees is a multiple of 90
func rotateWaypoint(waypoint *coords, degree int) {
	for degree > 0 {
		waypoint.x, waypoint.y = waypoint.y, -waypoint.x
		degree -= 90
	}
}
