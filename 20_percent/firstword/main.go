package main

import (
	"fmt"
)

/* func FirstWord(s string) string {
	for i, char := range s {
		if char == ' ' {
			return s[:i]
		}
	}
	return s
} */

func FirstWord(s string) string {
	// If the string is empty, just return the newline
	if s == "" {
		return "\n"
	}

	runes := []rune(s)
	n := len(runes)
	start := -1
	end := -1

	// 1. Find the start of the first word (skip leading spaces)
	for i := 0; i < n; i++ {
		if runes[i] != ' ' {
			start = i
			break
		}
	}

	// If the string was only spaces, return just a newline
	if start == -1 {
		return "\n"
	}

	// 2. Find the end of the first word
	for i := start; i < n; i++ {
		if runes[i] == ' ' {
			end = i
			break
		}
	}

	// If no space was found after the word, the word goes to the end
	if end == -1 {
		end = n
	}

	// 3. Construct the result string + newline
	return string(runes[start:end]) + "\n"
}

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}
