package main

import (
	"fmt"
)


func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	// Filter out all spaces from the original string first
	var filtered []rune
	for _, r := range str {
		if r != ' ' {
			filtered = append(filtered, r)
		}
	}

	// Check the length of the filtered content
	if len(filtered) < 5 {
		return "Invalid Input\n"
	}

	var result []rune
	count := 0

	for i := 0; i < len(filtered); i++ {
		// Add the character to our result
		result = append(result, filtered[i])
		count++

		// If we've reached 5 characters
		if count == 5 {
			// If there are still characters left in the filtered slice
			if i+1 < len(filtered) {
				// Add a space to separate the block
				result = append(result, ' ')
				// Skip the 6th character (the "skip" part of the function)
				i++
			}
			// Reset counter for the next block
			count = 0
		}
	}

	return string(result) + "\n"
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}
