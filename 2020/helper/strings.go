package helper

// Reverse reverses a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// RotateSlice rotates a slice of strings 90 degrees clockwise
func RotateSlice(slice *[]string) []string {
	newSlice := make([]string, len(*slice))
	for x := 0; x < len((*slice)[0]); x++ {
		newLine := ""
		for y := len(*slice) - 1; y >= 0; y-- {
			newLine += string((*slice)[y][x])
		}
		newSlice[x] = newLine
	}
	return newSlice
}

// FlipSlice flips a slice along its Horizontal axis
func FlipSlice(slice *[]string) []string {
	newSlice := make([]string, len(*slice))
	for i, j := 0, len(*slice)-1; i < j; i, j = i+1, j-1 {
		newSlice[i], newSlice[j] = (*slice)[j], (*slice)[i]
	}
	return newSlice
}
