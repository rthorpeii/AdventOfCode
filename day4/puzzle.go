// Package day4 has solutions for Day 4 of Advent of Code
// https://adventofcode.com/2020/day/4
package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"../input"
)

var inputFile string = "day4/input.txt"
var testFile string = "day4/testInput.txt"

/*

   byr (Birth Year)
   iyr (Issue Year)
   eyr (Expiration Year)
   hgt (Height)
   hcl (Hair Color)
   ecl (Eye Color)
   pid (Passport ID)
   cid (Country ID)

*/

// type Passport struct {
// 	byr int
// 	iyr int
// 	eyr int
// 	hgt int
// 	hcl string
// 	ecl string
// 	pid int
// 	cid int
// }

var fields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
var eyes = map[string](bool){"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}

func validatePassport(passport map[string](string)) bool {
	for _, field := range fields {
		value := passport[field]
		if value == "" && field != "cid" {
			return false
		}
		if field == "byr" {
			val, err := strconv.ParseFloat(value, 32)
			if err != nil || len(value) != 4 || val < 1920 || val > 2002 {
				return false
			}
		} else if field == "iyr" {
			val, err := strconv.ParseFloat(value, 32)
			if err != nil || len(value) != 4 || val < 2010 || val > 2020 {
				return false
			}
		} else if field == "eyr" {
			val, err := strconv.ParseFloat(value, 32)
			if err != nil || len(value) != 4 || val < 2020 || val > 2030 {
				return false
			}
		} else if field == "hgt" {
			if string(value[len(value)-2:len(value)]) == "cm" {
				val, err := strconv.ParseFloat(value[0:3], 32)
				if err != nil {
					return false
				}
				if len(value) != 5 || val < 150 || val > 193 {
					return false
				}
			} else if string(value[len(value)-2:len(value)]) == "in" {
				val, err := strconv.ParseFloat(value[0:2], 32)
				if err != nil {
					return false
				}
				if len(value) != 4 || val < 59 || val > 76 {
					return false
				}
			} else {
				return false
			}
		} else if field == "hcl" {
			matched, _ := regexp.MatchString(`#[0-9a-f]{6}$`, value)
			if !matched {
				return false
			}
		} else if field == "ecl" {
			if eyes[value] == false {
				return false
			}
		} else if field == "pid" {
			matched, _ := regexp.MatchString(`[0-9]{9}$`, value)
			if !matched {
				return false
			}
		}
	}
	if (len(passport) < 7) || (len(passport) == 7 && passport["cid"] != "") || len(passport) > 8 {
		return false
	}
	return true
}

// PartOne finds
func PartOne() int {
	rawInput := input.Slice(inputFile)

	valid := 0

	passport := make(map[string](string))
	for _, val := range rawInput {
		if val == "" {
			if validatePassport(passport) {
				valid++
			}
			passport = make(map[string](string))
			fmt.Println(passport)
			continue
		}

		fields := strings.Fields(val)
		for _, part := range fields {
			split := strings.Split(part, ":")
			passport[split[0]] = split[1]
		}
	}

	return valid
}

// PartTwo finds
func PartTwo() int {
	// rawInput := input.ReadInput(testFile)

	return -1
}
