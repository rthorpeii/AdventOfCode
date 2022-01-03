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

// MinInt calculates the max of two ints
func MinInt(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

//Mod takes the mod of two nums
func Mod(x int, y int) int {
	for x <= 0 {
		x += y
	}
	if x == y {
		return y
	}
	return x % y
}
