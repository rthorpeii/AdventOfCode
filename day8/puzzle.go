// Package day8 has solutions for Day 8 of Advent of Code
// https://adventofcode.com/2020/day/8
package day8

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var inputFile string = "day8/input.txt"
var testFile string = "day8/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(inputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(inputFile))
}

// PartOne finds
func PartOne(file string) int {
	rawInput := input.Slice(file)
	visited := make(map[int]bool)
	pos := 0
	acc := 0
	for pos < len(rawInput) {
		if visited[pos] {
			return acc
		}
		visited[pos] = true
		instr := rawInput[pos]
		parts := strings.Split(instr, " ")
		value, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "nop":
			pos++
		case "acc":
			acc += value
			pos++
		case "jmp":
			pos += value
		}
	}

	return -1
}

// PartTwo finds
func PartTwo(file string) int {
	rawInput := input.Slice(file)
	for index, line := range rawInput {
		parts := strings.Split(line, " ")

		if parts[0] == "jmp" {
			newProg := make([]string, len(rawInput))
			copy(newProg, rawInput)
			newProg[index] = "nop " + parts[1]
			visited := make(map[int]bool)
			pos := 0
			acc := 0
			value := runProg(newProg, pos, acc, &visited)
			if value != -1 {
				return value
			}
		} else if parts[0] == "nop" {
			newProg := make([]string, len(rawInput))
			copy(newProg, rawInput)
			newProg[index] = "jmp " + parts[1]
			visited := make(map[int]bool)
			pos := 0
			acc := 0
			value := runProg(newProg, pos, acc, &visited)
			if value != -1 {
				return value
			}
		}
	}
	return -1

}

func runProg(prog []string, pos int, acc int, visited *map[int]bool) int {
	for pos < len(prog) {
		if (*visited)[pos] {
			return -1
		}
		(*visited)[pos] = true
		instr := prog[pos]
		parts := strings.Split(instr, " ")
		value, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "nop":
			pos++
		case "acc":
			acc += value
			pos++
		case "jmp":
			pos += value
		}
	}
	return acc
}
