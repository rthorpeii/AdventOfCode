// Package day1 has solutions for Day 1 of Advent of Code
// https://adventofcode.com/2020/day/1
package day1

import (
	"../input"
)

var inputFile string = "day1/input.txt"
var testFile string = "day1/testInput.txt"

// PartOne finds the two entries that sum to 2020 and then multiply those two numbers together.
// It does so by placing the entries into a hashmap and iterating through the keys to see if it's 'opposite'
// (the number which it needs to add up to 2020) exists within the map. This allows it to find the solution
// in constant time O(n)
func PartOne() int {
	rawInput := input.ReadInput(inputFile)
	expenses := input.IntMap(rawInput)

	for expense := range expenses {
		// The number needed to sum to 2020 with expense
		opposite := 2020 - expense
		if expenses[opposite] == true {
			answer := opposite * expense
			return answer
		}
	}

	return -1
}

// PartTwo finds the three entries that sum to 2020 and then multiplies those together.
// Does so by checking each expense, adding it with an unchecked expense, and seeing if
// the value needed to sum with it to 2020 exists. Runs in O(n^2) time
func PartTwo() int {
	raw := input.ReadInput(inputFile)
	expenses := input.IntMap(raw)

	// creates a list of the keys in the expenses to iterate through
	keys := make([]int, 0, len(expenses))
	for expense := range expenses {
		keys = append(keys, expense)
	}

	// Iterate through each key, then iterate through then sum with the untried keys
	// before checking if the third value needed to sum to 2020 with the previous two exists
	for i := 0; i < len(keys); i++ {
		first := keys[i]
		for j := i + 1; j < len(keys); j++ {
			second := keys[j]
			third := 2020 - first - second
			if third >= 0 && expenses[third] {
				answer := first * second * third
				return answer
			}
		}
	}

	return -1
}
