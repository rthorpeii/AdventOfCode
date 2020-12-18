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
	cubes  map[int]cubeMap
	values map[int]bool
	size   int
	depth  int
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
			_, exists := dimension.cubes[coord]
			if !exists {
				dimension.cubes[coord] = newCubeMap(dimension.size, dimension.depth-1)
			}

			nextDimension := dimension.cubes[coord]
			currentCoords := append(coords, coord)
			nextCubes.cubes[coord] = (&nextDimension).cycleCube(currentCoords, origin, growth)
		}
	}

	return nextCubes
}

func (dimension cubeMap) findNextState(coords []int, value bool) bool {
	neighbors := dimension.countNeighbors(coords, false)
	if neighbors == 3 {
		return true
	}
	if neighbors == 2 && value {
		return true
	}
	return false

}

func (dimension *cubeMap) countNeighbors(coords []int, changed bool) int {
	if len(coords) == 1 {
		neighborCount := 0
		for delta := -1; delta <= 1; delta++ {
			coord := coords[0] + delta

			if dimension.values[coord] && (changed || delta != 0) {
				neighborCount++
			}

		}
		return neighborCount
	}
	neighborCount := 0
	for delta := -1; delta <= 1; delta++ {
		coord := coords[0] + delta

		_, exists := (*dimension).cubes[coord]
		if !exists {
			newMap := newCubeMap(dimension.size, dimension.depth-1)
			(*dimension).cubes[coord] = newMap
		}

		nextChanged := changed || (delta != 0)

		nextDim := (*dimension).cubes[coord]
		neighborCount += nextDim.countNeighbors(coords[1:], nextChanged)
	}

	return neighborCount
}

func (dimension cubeMap) getValue(coords []int) bool {
	if len(coords) == 1 {
		return dimension.values[coords[0]]
	}
	return dimension.cubes[coords[0]].getValue(coords[1:])
}

func (dimension *cubeMap) setValue(coords []int, value bool) {
	if len(coords) == 1 {
		(*dimension).values[coords[0]] = value
		return
	}
	_, exists := (*dimension).cubes[coords[0]]
	if !exists {
		newMap := newCubeMap(dimension.size, dimension.depth-1)
		(*dimension).cubes[coords[0]] = newMap
	}
	subMap := (*dimension).cubes[coords[0]]
	subMap.setValue(coords[1:], value)
}

func newCubeMap(size int, depth int) cubeMap {
	return cubeMap{size: size, values: make(map[int]bool), cubes: make(map[int]cubeMap), depth: depth}
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
		for _, value := range dimension.cubes {
			count += value.countActive()
		}
	}
	return count
}

func (dimension *cubeMap) printMap(top bool, growth int) {
	if dimension.depth == 0 {
		printString := ""
		if dimension.size == 0 {
			fmt.Println("--------size is 0 -----------", dimension)
		}
		for coord := -growth; coord <= dimension.size+growth; coord++ {
			if dimension.values[coord] {
				printString += "#"
			} else {
				printString += "."
			}
		}
		fmt.Println(printString)
		return
	}

	for coord := -growth; coord <= dimension.size+growth; coord++ {
		if top {
			fmt.Printf("Depth:%v \n", coord)
		}
		_, exists := dimension.cubes[coord]
		if !exists {
			dimension.cubes[coord] = newCubeMap(dimension.size, dimension.depth-1)
		}
		nextDim := dimension.cubes[coord]
		nextDim.printMap(false, growth)
	}
	return
}

// PartOne finds
func PartOne(file string) int {
	rawInput := input.Slice(file)
	dimensions := newCubeMap(len(rawInput), 2)

	for y, line := range rawInput {
		for x, value := range line {
			if string(value) == "#" {
				coords := []int{0, y, x}
				dimensions.setValue(coords, true)
			}
		}
	}

	for i := 0; i < 6; i++ {
		var coords []int
		dimensions = dimensions.cycleCube(coords, &dimensions, i+1)
	}

	return dimensions.countActive()
}

// PartTwo finds
func PartTwo(file string) int {
	rawInput := input.Slice(file)
	dimensions := newCubeMap(len(rawInput), 3)

	for y, line := range rawInput {
		for x, value := range line {
			if string(value) == "#" {
				coords := []int{0, 0, y, x}
				dimensions.setValue(coords, true)
			}
		}
	}

	for i := 0; i < 6; i++ {
		var coords []int
		dimensions = dimensions.cycleCube(coords, &dimensions, i+1)
	}

	return dimensions.countActive()
}
