// Package day21 has solutions for Day 21 of Advent of Code
// https://adventofcode.com/2020/day/21
package day21

import (
	"fmt"
	"sort"
	"strings"

	"github.com/rthorpeii/AdventOfCode2020/input"
)

var puzzleInputFile string = "day21/input.txt"
var testFile string = "day21/testInput.txt"

// SolvePuzzle prints the output produced by running the input and test files on both parts
func SolvePuzzle() {
	fmt.Printf("Part 1 - Test: %v \n", PartOne(testFile))
	fmt.Printf("Part 1 - Actual: %v \n", PartOne(puzzleInputFile))
	fmt.Printf("Part 2 - Test: %v \n", PartTwo(testFile))
	fmt.Printf("Part 2 - Actual: %v \n", PartTwo(puzzleInputFile))
}

// Menu holds stuff
type Menu struct {
	allergenMap    map[string]map[string]bool // Maps allergen to ingredients seen in
	ingredientMap  map[string]map[string]bool // Maps ingredients to allergens seen with
	ingredientMenu map[string][]string
}

func generateCandidates(file string) Menu {
	rawInput := input.Slice(file)

	allergenMap := make(map[string]map[string]bool)
	ingredientMap := make(map[string]map[string]bool)
	ingredientMenu := make(map[string][]string)
	for _, line := range rawInput {
		parts := strings.Split(line, " (contains ")
		ingredients := strings.Split(parts[0], " ")
		allergens := strings.Split(parts[1], ", ")
		allergens[len(allergens)-1] = allergens[len(allergens)-1][:len(allergens[len(allergens)-1])-1] // Strip the closing paren

		mappedAllergens := sliceToMap(&allergens)
		mappedIngredients := sliceToMap(&ingredients)
		for _, ingredient := range ingredients {
			// add the menu to the ingredient
			ingredientMenu[ingredient] = append(ingredientMenu[ingredient], line)

			_, exists := ingredientMap[ingredient]
			if !exists { // Initialiaze the ingredient mapping
				ingredientMap[ingredient] = sliceToMap(&allergens)
				continue
			}
			for allergen := range ingredientMap[ingredient] {
				if !mappedAllergens[allergen] {
					delete(ingredientMap[ingredient], allergen)
				}
			}
		}
		for _, allergen := range allergens {
			_, exists := allergenMap[allergen]
			if !exists { // Initialiaze the ingredient mapping
				allergenMap[allergen] = sliceToMap(&ingredients)
				continue
			}
			for ingredient := range allergenMap[allergen] {
				if !mappedIngredients[ingredient] {
					delete(allergenMap[allergen], ingredient)
				}
			}
		}

	}

	return Menu{allergenMap, ingredientMap, ingredientMenu}
}

func sliceToMap(slice *[]string) map[string]bool {
	endMap := make(map[string]bool)
	for _, value := range *slice {
		endMap[value] = true
	}

	return endMap
}

// PartOne finds
func PartOne(file string) int {
	menu := generateCandidates(file)
	possibleAllergens := make(map[string]bool)
	for _, ingredients := range menu.allergenMap {
		for ingredient := range ingredients {
			possibleAllergens[ingredient] = true
		}
	}

	count := 0
	for ingredient, menu := range menu.ingredientMenu {
		if !possibleAllergens[ingredient] {
			count += len(menu)
		}
	}

	return count
}

// PartTwo finds
func PartTwo(file string) string {
	menu := generateCandidates(file)
	possibleAllergens := make(map[string]bool)
	for _, ingredients := range menu.allergenMap {
		for ingredient := range ingredients {
			possibleAllergens[ingredient] = true
		}
	}

	knownAllergens := make(map[string]bool)
	allergenMap := menu.allergenMap
	var sortingList []string
	for len(allergenMap) > 0 {
	allergenCircuit:
		for allergen, ingredients := range allergenMap {
			for ingredient := range ingredients {
				if len(ingredients) == 1 {
					knownAllergens[ingredient] = true
					alString := allergen + ":" + ingredient
					sortingList = append(sortingList, alString)
					delete(allergenMap, allergen)
					continue allergenCircuit
				}
				if knownAllergens[ingredient] {
					delete(allergenMap[allergen], ingredient)
				}
			}

		}

	}
	// fmt.Println(knownAllergens)

	// var allergens []string

	// for ingredient := range knownAllergens {
	// 	allergens = append(allergens, ingredient)
	// }

	sort.Strings(sortingList)

	returnList := ""
	for _, ingredient := range sortingList {
		split := strings.Split(ingredient, ":")
		returnList += split[1] + ","
	}
	return returnList
}
