// Package day18 has solutions for Day 18 of Advent of Code
// https://adventofcode.com/2020/day/18
package day18

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var puzzleInputFile string = "day18/input.txt"
var testFile string = "day18/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(puzzleInputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(puzzleInputFile))
}

func parseLine(line string) int {
	split := strings.Fields(line)

	depthValue := make(map[int]int)
	operandDepth := make(map[int]string)
	depthValue[0] = 0
	depth := 0
	index := 0
	for index < len(split) {
		// fmt.Println(split, depthValue)
		value := string(split[index][0])
		// fmt.Println(value)
		op := operandDepth[depth]
		prevValue := depthValue[depth]
		// fmt.Println("prevValue: ", prevValue)

		if value == "(" {
			depth++
			depthValue[depth] = 0
			operandDepth[depth] = ""
		} else if value == ")" {
			depth--
			if operandDepth[depth] != "" {
				depthValue[depth] = runOp(operandDepth[depth], depthValue[depth], depthValue[depth+1])
			} else {
				depthValue[depth] = depthValue[depth+1]
			}
		} else {
			isNum, _ := regexp.MatchString(`[0-9]{1}`, value)
			if isNum {
				num, _ := strconv.Atoi(value)
				if op == "" {
					depthValue[depth] = num
				} else {
					depthValue[depth] = runOp(op, prevValue, num)
				}
			} else {
				operandDepth[depth] = value
			}
		}
		if len(split[index]) > 1 {
			split[index] = split[index][1:]
			continue
		}
		index++
	}
	// fmt.Println(depthValue)
	return depthValue[0]
}

func runOp(op string, a int, b int) int {
	switch op {
	case "*":
		return a * b
	case "+":
		return a + b
	}
	log.Fatal("Should have seen an op")
	return -1
}

// PartOne finds
func PartOne(file string) int {
	rawInput := input.Slice(file)

	count := 0
	for _, line := range rawInput {
		count += parseLine(line)
	}

	return count
}

// PartTwo finds
func PartTwo(file string) int {
	// rawInput := input.ReadInput(file string)

	return -1
}
