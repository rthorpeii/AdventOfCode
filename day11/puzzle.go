// Package day11 has solutions for Day 11 of Advent of Code
// https://adventofcode.com/2020/day/11
package day11

import (
	"fmt"

	"github.com/rthorpeii/AdventOfCode2020/helper"
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

// PartOne finds how many empty seats there are after the layout stablizes
func PartOne(file string) int {
	input := input.Slice(file)

	changed, emptySeats := true, 0
	for changed {
		changed, emptySeats = updateLayout(&input, false, 3)
	}
	return emptySeats
}

// PartTwo finds how many empty seats there are after the layout stablizes for a new set of rules
func PartTwo(file string) int {
	input := input.Slice(file)

	changed, emptySeats := true, 0
	for changed {
		changed, emptySeats = updateLayout(&input, true, 4)
	}
	return emptySeats
}

// updateLayout applies the rules for people changing seats to the current layout,
// and returns whether any seats have changed, and how many seats in the layout are empty
func updateLayout(layout *[]string, sight bool, threshold int) (changed bool, emptySeats int) {
	changed, emptySeats = false, 0
	updatedLayout := make([]string, len(*layout))
	for y, row := range *layout {
		updatedRow := ""
		for x := range row {
			pos := coords{x, y}
			currentValue := string(row[x])
			var updatedValue string
			switch currentValue {
			case "L":
				updatedValue = updateSeat(layout, pos, sight, 0)
			case "#":
				updatedValue = updateSeat(layout, pos, sight, threshold)
			case ".":
				updatedValue = "."
			}
			if updatedValue != currentValue {
				changed = true
			}
			if updatedValue == "#" {
				emptySeats++
			}
			updatedRow += updatedValue
		}
		updatedLayout[y] = updatedRow
	}
	*layout = updatedLayout
	return changed, emptySeats
}

// updateSeat returns what value the seat at pos should have after updating the layout.
// If sight is true, the function will base the rules off of the first seat within sight
// in a given direction, other wise it will base the rules off of the seats immediately
// adjacent to it. Returns an empty seat after the threshold number of occupied seats is encounterd.
func updateSeat(layout *[]string, pos coords, sight bool, threshold int) string {
	dist := 1
	if sight {
		dist = helper.MaxInt(len(*layout), len((*layout)[pos.y]))
	}

	visible := 0
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			dir := coords{x, y}
			visible += visiblyOccupied(layout, pos, dir, dist)
			if visible > threshold {
				return "L"
			}
		}
	}

	return "#"
}

// visiblyOccupied looks from pos in the direction specified for the first seat within dist
// Returns 1 if that seat is occupied, 0 Otherwise.
func visiblyOccupied(layout *[]string, pos coords, dir coords, dist int) int {
	target := coords{pos.x, pos.y}
	for i := 0; i < dist; i++ {
		target.x += dir.x
		target.y += dir.y
		if !coordsValid(layout, target) || (target == pos) {
			break
		}
		switch string((*layout)[target.y][target.x]) {
		case "#":
			return 1
		case "L":
			return 0
		}
	}
	return 0
}

// coordsValid returns whether the pos is a valid set of coordinates within the layout
func coordsValid(layout *[]string, pos coords) bool {
	if pos.y >= len(*layout) || pos.y < 0 {
		return false
	}
	return (pos.x > -1 && pos.x < len((*layout)[pos.y]))
}
