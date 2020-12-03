// Package day3 has solutions for Day 3 of Advent of Code
// https://adventofcode.com/2020/day/3
package day3

import (
	"../input"
)

var inputFile string = "day3/input.txt"
var testFile string = "day3/testInput.txt"

// PartOne finds
func PartOne() int {
	rawInput := input.Slice(inputFile)
	x := 0

	trees := 0
	for y := 1; y < len(rawInput); y++ {
		x += 3
		if x >= len(rawInput[y]) {
			x -= len(rawInput[y])
		}

		if string(rawInput[y][x]) == "#" {
			trees++
		}
	}

	return trees
}

// PartTwo finds
func PartTwo() int {
	type Pair struct {
		x, y int
	}

	slopes := []Pair{Pair{1, 1}, Pair{3, 1}, Pair{5, 1}, Pair{7, 1}, Pair{1, 2}}
	rawInput := input.Slice(inputFile)

	answer := 1
	for _, slope := range slopes {
		trees := 0
		x := 0
		for y := slope.y; y < len(rawInput); y += slope.y {
			x += slope.x
			if x >= len(rawInput[y]) {
				x -= len(rawInput[y])
			}

			if string(rawInput[y][x]) == "#" {
				trees++
			}
		}
		answer *= trees
	}

	return answer
}
