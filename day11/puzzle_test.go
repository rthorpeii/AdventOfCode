// Package day11 has solutions for Day 11 of Advent of Code
// https://adventofcode.com/2020/day/11
package day11

import "testing"

func TestPartOne(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Using the default test file",
			args: args{"./testInput.txt"},
			want: 37,
		},
		{
			name: "Using the input file",
			args: args{"./input.txt"},
			want: 2281,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartOne(tt.args.file); got != tt.want {
				t.Errorf("PartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
