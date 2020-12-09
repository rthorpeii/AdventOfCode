// Package day8 has solutions for Day 8 of Advent of Code
// https://adventofcode.com/2020/day/8
package day8

import (
	"fmt"
	"strings"

	"github.com/rthorpeii/AdventOfCode2020/vm"
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
	vm := vm.NewVM(file)
	vm.Execute()
	return vm.Acc
}

// PartTwo finds
func PartTwo(file string) int {
	vm := vm.NewVM(file)
	for index, line := range vm.Instructions {
		parts := strings.Split(line, " ")
		if parts[0] == "acc" {
			continue
		}
		copyVM := vm.Copy()
		if parts[0] == "jmp" {
			copyVM.Instructions[index] = "nop " + parts[1]
		} else if parts[0] == "nop" {
			copyVM.Instructions[index] = "jmp " + parts[1]
		}

		valid := copyVM.Execute()
		if valid {
			return copyVM.Acc
		}
	}

	return -1
}
