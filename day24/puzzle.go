// Package day24 has solutions for Day 24 of Advent of Code
// https://adventofcode.com/2020/day/24
package day24

import (
	"fmt"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var puzzleInputFile string = "day24/input.txt"
var testFile string = "day24/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(puzzleInputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(puzzleInputFile))
}

// Coord is a coordinate
type Coord struct {
	x int
	y int
}

func parseDirection(directions string) []string {
	var finalDirections []string

	for i := 0; i < len(directions); i++ {
		char := directions[i]
		switch char {
		case 'e':
			finalDirections = append(finalDirections, "e")
		case 'w':
			finalDirections = append(finalDirections, "w")
		case 's':
			finalDirections = append(finalDirections, directions[i:i+2])
			i++
		case 'n':
			finalDirections = append(finalDirections, directions[i:i+2])
			i++
		}
	}
	return finalDirections
}

func followDirections(directions *[]string) Coord {
	coord := Coord{}

	for _, dir := range *directions {
		switch dir {
		case "e":
			coord.x += 2
		case "w":
			coord.x -= 2
		case "se":
			coord.x++
			coord.y--
		case "sw":
			coord.x--
			coord.y--
		case "nw":
			coord.x--
			coord.y++
		case "ne":
			coord.x++
			coord.y++
		}
	}
	return coord
}

// PartOne finds
func PartOne(file string) int {
	rawInput := input.Slice(file)

	// false is white
	orientations := make(map[Coord]bool)

	for _, line := range rawInput {
		directions := parseDirection(line)
		coord := followDirections(&directions)
		orientations[coord] = !orientations[coord]
	}

	count := 0
	for _, value := range orientations {
		if value {
			count++
		}
	}

	return count
}

// PartTwo finds
func PartTwo(file string) int {
	rawInput := input.Slice(file)

	// false is white
	orientations := make(map[Coord]bool)

	for _, line := range rawInput {
		directions := parseDirection(line)
		coord := followDirections(&directions)
		orientations[coord] = !orientations[coord]
	}

	for day := 0; day < 100; day++ {
		orientations = *flipTiles(&orientations)
	}

	count := 0
	for _, value := range orientations {
		if value {
			count++
		}
	}
	return count
}

func flipTiles(tiles *map[Coord]bool) *map[Coord]bool {
	newTiles := make(map[Coord]bool)
	for coord, value := range *tiles {
		if value {
			adjacent := coord.adjacent()
			adjacentBlackCount := adjacentBlack(tiles, &adjacent)
			if adjacentBlackCount == 1 || adjacentBlackCount == 2 {
				newTiles[coord] = true
			}

			for _, adjacentCoord := range adjacent {
				if !(*tiles)[adjacentCoord] {
					adjacentToWhite := adjacentCoord.adjacent()
					numBlackNextTowhite := adjacentBlack(tiles, &adjacentToWhite)
					if numBlackNextTowhite == 2 {
						newTiles[adjacentCoord] = true
					}
				}
			}
		}
	}
	return &newTiles
}

func adjacentBlack(tiles *map[Coord]bool, adjacentTiles *[]Coord) int {
	count := 0
	for _, coord := range *adjacentTiles {
		if (*tiles)[coord] {
			count++
		}
	}
	return count
}

func (coord *Coord) adjacent() []Coord {
	adjacent := make([]Coord, 6)
	adjacent[0] = Coord{coord.x + 1, coord.y + 1}
	adjacent[1] = Coord{coord.x + 2, coord.y}
	adjacent[2] = Coord{coord.x + 1, coord.y - 1}
	adjacent[3] = Coord{coord.x - 1, coord.y - 1}
	adjacent[4] = Coord{coord.x - 2, coord.y}
	adjacent[5] = Coord{coord.x - 1, coord.y + 1}
	return adjacent
}
