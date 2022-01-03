// Package day20 has solutions for Day 20 of Advent of Code
// https://adventofcode.com/2020/day/20
package day20

import (
	"fmt"
	"strings"
)

var puzzleInputFile string = "day20/input.txt"
var testFile string = "day20/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(puzzleInputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(puzzleInputFile))
}

// PartOne finds the corners of a corrupted image and multiplies their IDs together
func PartOne(file string) int {
	image := newImage(file)
	corners := image.findCorners()

	value := 1
	for _, tile := range corners {
		value *= tile.id
	}
	return value
}

// seaMonster is an ASCI sea monster that we need to match
var seaMonster = []string{
	"                  # ",
	"#    ##    ##    ###",
	" #  #  #  #  #  #   "}

// PartTwo finds how many ASCI sea monsters are in the image
func PartTwo(file string) int {
	image := newImage(file)
	image.assemble()
	image.createFinalImage()
	image.findMonsters(&seaMonster)

	count := 0
	for _, line := range image.finalImage {
		count += strings.Count(line, "#")
	}

	return count
}
