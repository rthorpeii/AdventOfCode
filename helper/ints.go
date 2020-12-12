// Package helper includes helper methods for solving puzzles faster
package helper

// AbsInt returns the absolute value of x.
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// MaxInt calculates the max of two ints
func MaxInt(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
