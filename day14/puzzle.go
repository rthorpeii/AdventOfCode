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

// PartOne finds the sum of the memory after running the inputs, while masking the value
func PartOne(file string) int {
	input := input.Slice(file)
	memory := make(map[string]string)

	var mask string
	for _, line := range input {
		split := strings.Split(line, " = ")
		if split[0] == "mask" {
			mask = split[1]
		} else {
			address := split[0][4 : len(split[0])-1]
			num := base10StringToBase2(split[1])
			newValue := applyMask(mask, num, "X")
			memory[address] = newValue
		}
	}

	return sumMemory(memory)
}

// base10StringToBase2 converts a number from base 10 to base 2
func base10StringToBase2(num string) string {
	base10Num, _ := strconv.ParseInt(num, 10, 64)
	base2Num := strconv.FormatInt(base10Num, 2)
	for len(base2Num) < 36 {
		base2Num = "0" + base2Num
	}
	return base2Num
}

// applyMask overwrites the bits in num with the corresponding bit in the mask
// bits that match the value of skip are skipped
func applyMask(mask string, num string, skip string) string {
	for index, value := range mask {
		if string(value) != skip {
			num = num[:index] + string(value) + num[index+1:]
		}
	}
	return num
}

// sumMemory sums the values within the memory addresses
func sumMemory(memory map[string]string) int {
	sum := 0
	for _, val := range memory {
		base10Val, _ := strconv.ParseInt(val, 2, 64)
		sum += int(base10Val)
	}
	return sum
}

// PartTwo finds the sum of the memory after running the inputs, while masking the address
func PartTwo(file string) int {
	input := input.Slice(file)
	memory := make(map[string]string)

	var mask string
	for _, line := range input {
		split := strings.Split(line, " = ")
		if split[0] == "mask" {
			mask = split[1]
		} else {
			address := base10StringToBase2(split[0][4 : len(split[0])-1])
			maskedAddress := applyMask(mask, address, "0")
			value := base10StringToBase2(split[1])
			for _, newAddr := range findAddresses(maskedAddress) {
				memory[newAddr] = value
			}
		}
	}

	return sumMemory(memory)
}

// findAddresses takes a masked address and finds all possible addresses that can be
// created from the masked address
func findAddresses(maskedAddress string) []string {
	var addresses []string
	var prefix string
	for index, char := range maskedAddress {
		if string(char) != "X" {
			prefix += string(char)
		} else {
			subAddresses := findAddresses(maskedAddress[index+1:])
			for _, address := range subAddresses {
				addresses = append(addresses, prefix+"0"+address, prefix+"1"+address)
			}
			return addresses
		}
	}
	return append(addresses, prefix)
}
