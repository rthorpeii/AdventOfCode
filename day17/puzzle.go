// Package day17 has solutions for Day 17 of Advent of Code
// https://adventofcode.com/2020/day/17
package day17

import (
	"fmt"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var puzzleInputFile string = "day17/input.txt"
var testFile string = "day17/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(puzzleInputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(puzzleInputFile))
}

type cubeMap struct {
	subMap map[int]cubeMap
	values map[int]bool
	size   int
	depth  int
}

func newCubeMap(size int, depth int) cubeMap {
	return cubeMap{
		size:   size,
		values: make(map[int]bool),
		subMap: make(map[int]cubeMap),
		depth:  depth}
}

func (dimension *cubeMap) cycleCube(coords []int, origin *cubeMap, growth int) cubeMap {
	nextCubes := newCubeMap(dimension.size, dimension.depth)
	if dimension.depth == 0 {
		newValues := make(map[int]bool)
		for coord := -growth; coord <= dimension.size+growth; coord++ {
			currentCoords := append(coords, coord)
			value := dimension.values[coord]
			nextValue := origin.findNextState(currentCoords, value)
			newValues[coord] = nextValue
		}
		nextCubes.values = newValues
	} else {
		for coord := -growth; coord <= dimension.size+growth; coord++ {
			nextDimension := dimension.getSubMap(coord)
			currentCoords := append(coords, coord)
			nextCubes.subMap[coord] = nextDimension.cycleCube(currentCoords, origin, growth)
		}
	}

	return nextCubes
}

func (dimension cubeMap) findNextState(coords []int, value bool) bool {
	neighbors := dimension.countNeighbors(coords, false)
	return (neighbors == 3 || (neighbors == 2 && value))
}

func (dimension *cubeMap) countNeighbors(coords []int, changed bool) int {
	neighborCount := 0
	if len(coords) == 1 {
		for delta := -1; delta <= 1; delta++ {
			coord := coords[0] + delta
			if dimension.values[coord] && (changed || delta != 0) {
				neighborCount++
			}
		}
	} else {
		for delta := -1; delta <= 1; delta++ {
			coord := coords[0] + delta
			nextChanged := changed || (delta != 0)

			nextDim := dimension.getSubMap(coord)
			neighborCount += nextDim.countNeighbors(coords[1:], nextChanged)
		}
	}

	return neighborCount
}

func (dimension *cubeMap) getSubMap(coord int) *cubeMap {
	_, exists := (*dimension).subMap[coord]
	if !exists {
		newMap := newCubeMap(dimension.size, dimension.depth-1)
		(*dimension).subMap[coord] = newMap
	}
	subMap := dimension.subMap[coord]
	return &subMap
}

func (dimension cubeMap) getValue(coords []int) bool {
	if len(coords) == 1 {
		return dimension.values[coords[0]]
	}
	return dimension.subMap[coords[0]].getValue(coords[1:])
}

func (dimension *cubeMap) setValue(coords []int, value bool) {
	if len(coords) == 1 {
		(*dimension).values[coords[0]] = value
		return
	}
	dimension.getSubMap(coords[0]).setValue(coords[1:], value)
}

func (dimension cubeMap) countActive() int {
	count := 0

	if len(dimension.values) > 0 {
		for _, value := range dimension.values {
			if value {
				count++
			}
		}
	} else {
		for _, cube := range dimension.subMap {
			count += cube.countActive()
		}
	}
	return count
}

func (dimension *cubeMap) printMap(top bool, growth int) {
	if dimension.depth != 0 {
		for coord := -growth; coord <= dimension.size+growth; coord++ {
			if top {
				fmt.Printf("Depth:%v \n", coord)
			}
			dimension.getSubMap(coord).printMap(false, growth)
		}
	}

	printString := ""
	for coord := -growth; coord <= dimension.size+growth; coord++ {
		if dimension.values[coord] {
			printString += "#"
		} else {
			printString += "."
		}
	}
	fmt.Println(printString)
}

// PartOne finds
func PartOne(file string) int {
	dimensions := generateStart(file, 3)
	for i := 1; i <= 6; i++ {
		var coords []int
		dimensions = dimensions.cycleCube(coords, &dimensions, i)
	}
	return dimensions.countActive()
}

// PartTwo finds
func PartTwo(file string) int {
	dimensions := generateStart(file, 4)
	for i := 1; i <= 6; i++ {
		var coords []int
		dimensions = dimensions.cycleCube(coords, &dimensions, i)
	}
	return dimensions.countActive()
}

func generateStart(file string, numDimensions int) cubeMap {
	rawInput := input.Slice(file)
	dimensions := newCubeMap(len(rawInput), numDimensions-1)
	for y, line := range rawInput {
		for x, initialValue := range line {
			if string(initialValue) == "#" { // An active cube
				coords := make([]int, numDimensions)
				coords[len(coords)-1], coords[len(coords)-2] = x, y
				dimensions.setValue(coords, true)
			}
		}
	}
	return dimensions
}
