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

// Tile is tile at a specific (x,y) coordinate
type Tile struct {
	x int
	y int
}

func generateStart(file string) *map[Tile]bool {
	// false is white
	tiles := make(map[Tile]bool)

	for _, line := range input.Slice(file) {
		directions := parseDirection(line)
		coord := followDirections(directions)
		tiles[coord] = !tiles[coord]
	}
	return &tiles
}

func parseDirection(directions string) *[]string {
	var finalDirections []string
	index := 0
	for index < len(directions) {
		dir := string(directions[index])
		if dir == "s" || dir == "n" {
			index++
			dir += string(directions[index])
		}
		finalDirections = append(finalDirections, dir)
		index++
	}
	return &finalDirections
}

func followDirections(directions *[]string) Tile {
	coord := Tile{}
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

func countBlackTiles(tiles *map[Tile]bool) int {
	count := 0
	for _, value := range *tiles {
		if value {
			count++
		}
	}
	return count
}

// PartOne finds
func PartOne(file string) int {
	tiles := generateStart(file)
	return countBlackTiles(tiles)
}

// PartTwo finds
func PartTwo(file string) int {
	tiles := generateStart(file)
	for day := 0; day < 100; day++ {
		flipTiles(tiles)
	}

	return countBlackTiles(tiles)
}

func flipTiles(tiles *map[Tile]bool) {
	newTiles := make(map[Tile]bool)
	for coord, value := range *tiles {
		if value {
			adjacent := coord.adjacent()
			adjacentBlack := countAdjacentBlack(tiles, &adjacent)
			if adjacentBlack == 1 || adjacentBlack == 2 {
				newTiles[coord] = true
			}

			for _, adjacentCoord := range adjacent {
				adjacentCoord.setNextState(tiles, &newTiles)
			}
		}
	}
	*tiles = newTiles
}

func (tile *Tile) adjacent() []Tile {
	adjacent := make([]Tile, 6)
	adjacent[0] = Tile{tile.x + 1, tile.y + 1}
	adjacent[1] = Tile{tile.x + 2, tile.y}
	adjacent[2] = Tile{tile.x + 1, tile.y - 1}
	adjacent[3] = Tile{tile.x - 1, tile.y - 1}
	adjacent[4] = Tile{tile.x - 2, tile.y}
	adjacent[5] = Tile{tile.x - 1, tile.y + 1}
	return adjacent
}

func countAdjacentBlack(tiles *map[Tile]bool, adjacentTiles *[]Tile) int {
	count := 0
	for _, coord := range *adjacentTiles {
		if (*tiles)[coord] {
			count++
		}
	}
	return count
}

func (tile Tile) setNextState(currentTiles *map[Tile]bool, newTiles *map[Tile]bool) {
	if !(*currentTiles)[tile] {
		adjacent := tile.adjacent()
		if countAdjacentBlack(currentTiles, &adjacent) == 2 {
			(*newTiles)[tile] = true
		}
	}
}
