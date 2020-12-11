package input

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// ReadInput takes an input file and converts it into a string
func ReadInput(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Failed to convert input to string", err)
	}
	return string(data)
}

// Slice parses the input into a slice
func Slice(file string) []string {
	return strings.Split(ReadInput(file), "\n")
}

// IntSlice turns the input into a map
func IntSlice(input string) []int {
	inputSlice := Slice(input)

	output := make([]int, len(inputSlice))

	for pos, num := range inputSlice {
		number, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal("Failed to convert input to map", err)
		}
		output[pos] = number
	}

	return output
}

// IntMap turns the input into a map
func IntMap(input string) map[int]bool {
	inputSlice := Slice(input)

	inputMap := make(map[int]bool)

	for _, num := range inputSlice {
		number, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal("Failed to convert input to map", err)
		}
		inputMap[number] = true
	}

	return inputMap
}
