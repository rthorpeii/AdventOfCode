// Package day17 has solutions for Day 17 of Advent of Code
// https://adventofcode.com/2020/day/17
package day17

// import (
// 	"fmt"

// 	"github.com/rthorpeii/AdventOfCode2020/input"
// )

// var puzzleInputFile string = "day17/input.txt"
// var testFile string = "day17/testInput.txt"

// // SolvePuzzle2 prints the output produced by running the input and test files on both parts
// func SolvePuzzle2() {
// 	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
// 	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(puzzleInputFile))
// }

// type cubeMap struct {
// 	cubes  map[int]map[int]map[int]map[int]bool
// 	height int
// 	width  int
// 	growth int
// }

// func updateCube(cubes *map[int]map[int]map[int]map[int]bool, x int, y int, z int, w int, value bool) {

// 	_, wExists := (*cubes)[w]
// 	if !wExists {
// 		(*cubes)[w] = make(map[int]map[int]map[int]bool)
// 		(*cubes)[w][z] = make(map[int]map[int]bool)
// 		(*cubes)[w][z][y] = make(map[int]bool)
// 	} else {
// 		_, zExists := (*cubes)[w][z]
// 		if !zExists {
// 			(*cubes)[w][z] = make(map[int]map[int]bool)
// 			(*cubes)[w][z][y] = make(map[int]bool)
// 		} else {
// 			_, yExists := (*cubes)[w][z][y]
// 			if !yExists {
// 				(*cubes)[w][z][y] = make(map[int]bool)
// 			}
// 		}
// 	}

// 	(*cubes)[w][z][y][x] = value
// }

// // PartOne finds
// func PartTwo(file string) int {
// 	rawInput := input.Slice(file)
// 	z := make(map[int]map[int]map[int]map[int]bool)

// 	for y, line := range rawInput {
// 		for x, value := range line {
// 			if string(value) == "#" {
// 				updateCube(&z, x, y, 0, 0, true)
// 			}
// 		}
// 	}

// 	mapping := cubeMap{z, len(rawInput), len(rawInput[0]), 1}
// 	for i := 0; i < 6; i++ {
// 		// printMap(&mapping)
// 		changeState(&mapping)
// 	}

// 	return countActive(&mapping)
// }

// func printMap(cube *cubeMap) {
// 	for w := -(*cube).growth; w <= (*cube).growth; w++ {
// 		for z := -(*cube).growth; z <= (*cube).growth; z++ {
// 			fmt.Printf("w: %v z: %v\n", w, z)
// 			for y := -(*cube).growth; y <= (*cube).height+(*cube).growth; y++ {
// 				line := ""
// 				for x := -(*cube).growth; x <= (*cube).width+(*cube).growth; x++ {
// 					if (*cube).cubes[w][z][y][x] {
// 						line += "#"
// 					} else {
// 						line += "."
// 					}
// 				}
// 				fmt.Println(line)
// 			}
// 		}
// 	}
// }

// func countActive(cubes *cubeMap) int {
// 	count := 0
// 	for _, z := range (*cubes).cubes {
// 		for _, y := range z {
// 			for _, x := range y {
// 				for _, active := range x {
// 					if active {
// 						count++
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return count
// }
// func changeState(cube *cubeMap) {
// 	newStates := make(map[int]map[int]map[int]map[int]bool)
// 	for w := -(*cube).growth; w <= (*cube).growth; w++ {
// 		for z := -(*cube).growth; z <= (*cube).growth; z++ {
// 			for y := -(*cube).growth; y <= (*cube).height+(*cube).growth; y++ {
// 				for x := -(*cube).growth; x <= (*cube).width+(*cube).growth; x++ {
// 					value := checkChange(&(*cube).cubes, x, y, z, w)
// 					updateCube(&newStates, x, y, z, w, value)
// 				}
// 			}
// 		}
// 	}

// 	(*cube).cubes = newStates
// 	(*cube).growth++
// }

// func checkChange(cubes *map[int]map[int]map[int]map[int]bool, x int, y int, z int, w int) bool {
// 	surrounding := 0
// 	for dw := -1; dw <= 1; dw++ {
// 		for dz := -1; dz <= 1; dz++ {
// 			for dy := -1; dy <= 1; dy++ {
// 				for dx := -1; dx <= 1; dx++ {
// 					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
// 						continue
// 					}
// 					if (*cubes)[w+dw][z+dz][y+dy][x+dx] {
// 						surrounding++
// 					}
// 				}
// 			}
// 		}
// 	}

// 	if (*cubes)[w][z][y][x] {
// 		// fmt.Println((*cubes)[z][y][x], x, y, z, surrounding)
// 	}

// 	if surrounding == 3 {
// 		return true
// 	} else if (*cubes)[w][z][y][x] && surrounding == 2 {
// 		return true
// 	}
// 	return false
// }
