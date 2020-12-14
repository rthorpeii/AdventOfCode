// Package day14 has solutions for Day 14 of Advent of Code
// https://adventofcode.com/2020/day/14
package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var puzzleInputFile string = "day14/input.txt"
var testFile string = "day14/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(puzzleInputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(puzzleInputFile))
}

// PartOne finds
func PartOne(file string) int {
	input := input.Slice(file)
	memory := make(map[uint64]map[uint]bool)

	var mask string
	for _, line := range input {
		split := strings.Split(line, " = ")
		if split[0] == "mask" {
			mask = split[1]
			continue
		}
		// fmt.Println(split[0])
		strAddress := split[0][4 : len(split[0])-1]
		// fmt.Println(strAddress)
		address, _ := strconv.ParseUint(strAddress, 10, 64)
		base10Num, _ := strconv.ParseInt(split[1], 10, 64)
		num := strconv.FormatInt(base10Num, 2)
		setAddress(&memory, address, num, mask)
	}

	returnVal := 0

	for _, val := range memory {
		strVal := bitToString(val)
		base10Val, _ := strconv.ParseInt(strVal, 2, 64)
		returnVal += int(base10Val)
	}
	// fmt.Println(memory)
	return returnVal
}

func bitToString(bitmap map[uint]bool) string {
	var value string
	for i := 35; i > -1; i-- {
		if bitmap[uint(i)] {
			value += "1"
		} else {
			value += "0"
		}
	}
	return value
}

func padNum(num string) string {
	for len(num) < 36 {
		num = "0" + num
	}
	return num
}

func setAddress(addresses *map[uint64]map[uint]bool, address uint64, num string, mask string) {
	newValue := make(map[uint]bool)
	num = padNum(num)
	fmt.Println(num)
	for index, value := range num {
		bit := uint((len(mask) - 1) - index)
		if string(value) == "1" {
			newValue[bit] = true
		}
	}
	fmt.Printf("%v: %v Before mask\n", address, newValue)
	applyMask(mask, &newValue)
	fmt.Printf("%v: %v After mask\n", address, newValue)

	(*addresses)[address] = newValue
}

func applyMask(mask string, num *map[uint]bool) {
	for index, value := range mask {
		bit := uint((len(mask) - 1) - index)

		if string(value) == "0" {
			(*num)[bit] = false
		} else if string(value) == "1" {
			(*num)[bit] = true
		}
	}
}

// PartTwo finds
func PartTwo(file string) int {
	// rawInput := input.ReadInput(file string)

	return -1
}
