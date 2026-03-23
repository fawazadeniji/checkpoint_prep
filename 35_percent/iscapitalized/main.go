package main

import (
	"fmt"
)

func IsCapitalized(s string) bool {
	// If the string is empty, return false per instructions
	if s == "" {
		return false
	}

	// We iterate through the string
	for i := 0; i < len(s); i++ {
		// A word starts if:
		// 1. It is the first character (index 0)
		// 2. The character before it was a space
		if i == 0 || s[i-1] == ' ' {
			// Check if the starting character is a lowercase letter
			if s[i] >= 'a' && s[i] <= 'z' {
				return false
			}
		}
	}

	// If we checked all words and found no lowercase starts, return true
	return true
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
