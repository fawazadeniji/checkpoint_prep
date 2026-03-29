package main

import (
	"fmt"
	"os"
)

// Helper function to check if a character is a vowel
func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}

func main() {
	// If the number of arguments is not exactly 1, print nothing and exit
	if len(os.Args) != 2 {
		return
	}

	word := os.Args[1]
	vowelIndex := -1

	// Iterate through the word to find the index of the first vowel
	for i := 0; i < len(word); i++ {
		if isVowel(word[i]) {
			vowelIndex = i
			break
		}
	}

	// Rule 3: No vowels found
	if vowelIndex == -1 {
		fmt.Println("No vowels")
		return
	}

	// Rule 1: Starts with a vowel
	if vowelIndex == 0 {
		fmt.Println(word + "ay")
	} else {
		// Rule 2: Starts with a consonant
		// Slice from the vowel to the end + slice from the start to the vowel + "ay"
		fmt.Println(word[vowelIndex:] + word[:vowelIndex] + "ay")
	}
}
