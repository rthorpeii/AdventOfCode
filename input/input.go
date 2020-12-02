package input

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// ReadInput takes an input file and converts it into a string
func ReadInput(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}
	return string(data)
}

// Slice parses the input into a slice
func Slice(input string) []string {
	return strings.Split(input, "\n")
}

// Map turns the input into a map
func Map(input string) map[int]bool {
	inputSlice := Slice(input)

	var inputMap map[int]bool
	inputMap = make(map[int]bool)

	for _, num := range inputSlice {
		number, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal("Failed to convert input to map", err)
		}
		inputMap[number] = true
	}

	return inputMap
}
