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
	_, value := runProg(rawInput)
	return value
}

// PartTwo finds
func PartTwo(file string) int {
	rawInput := input.Slice(file)
	for index, line := range rawInput {
		parts := strings.Split(line, " ")
		if parts[0] == "acc" {
			continue
		}
		newProg := make([]string, len(rawInput))
		copy(newProg, rawInput)
		if parts[0] == "jmp" {
			newProg[index] = "nop " + parts[1]
		} else if parts[0] == "nop" {
			newProg[index] = "jmp " + parts[1]
		}

		valid, value := runProg(newProg)
		if valid {
			return value
		}
	}

	return -1
}

func runProg(prog []string) (bool, int) {
	pos, acc, visited := 0, 0, make(map[int]bool)
	for pos < len(prog) {
		if visited[pos] {
			return false, acc
		}
		visited[pos] = true

		parts := strings.Split(prog[pos], " ")
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
	return true, acc
}
