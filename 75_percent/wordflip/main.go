package main

import (
	"fmt"
)

func WordFlip(str string) string {
	// 1. Initial check for empty string
	if str == "" {
		return "Invalid Output\n"
	}

	var words []string
	currentWord := ""

	// 2. Extract words and ignore extra spaces
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			currentWord += string(str[i])
		} else {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		}
	}
	// Catch the last word if it didn't end with a space
	if currentWord != "" {
		words = append(words, currentWord)
	}

	// 3. Handle cases that were only spaces (no words found)
	if len(words) == 0 {
		return "Invalid Output\n"
	}

	// 4. Build the reversed string
	result := ""
	for i := len(words) - 1; i >= 0; i-- {
		result += words[i]
		if i > 0 {
			result += " " // Add space between words, but not after the last one
		}
	}

	return result + "\n"
}

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}
