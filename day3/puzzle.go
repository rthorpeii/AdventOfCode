// Package day3 has solutions for Day 3 of Advent of Code
// https://adventofcode.com/2020/day/3
package day3

import (
	"../input"
)

var inputFile string = "day3/input.txt"
var testFile string = "day3/testInput.txt"

// Slope represents a slope to take down the mountain
type Slope struct {
	x, y int
}

// traverseSlope takes the terrain (A slice of strings of multiple "."s and "#"s )
// and determines how many trees ("#"s) would be hit if you moved down the terrain
// following the given slope.
func traverseSlope(terrain *[]string, slope Slope) int {
	var x, trees = 0, 0
	for y := slope.y; y < len(*terrain); y += slope.y {
		x += slope.x
		x %= len((*terrain)[y])

		if string((*terrain)[y][x]) == "#" {
			trees++
		}
	}
	return trees
}

// PartOne finds how many trees would you encounter if your slope was (1,1)
func PartOne() int {
	rawInput := input.Slice(inputFile)
	return traverseSlope(&rawInput, Slope{3, 1})
}

// PartTwo finds how many trees would you encounter across several slopes, and multiplies those together
func PartTwo() int {
	slopes := []Slope{Slope{1, 1}, Slope{3, 1}, Slope{5, 1}, Slope{7, 1}, Slope{1, 2}}
	rawInput := input.Slice(inputFile)

	answer := 1
	for _, slope := range slopes {
		answer *= traverseSlope(&rawInput, slope)
	}

	return answer
}
