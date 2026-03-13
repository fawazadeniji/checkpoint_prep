package main

import (
	"fmt"
)

func RepeatAlpha(s string) string {
	var result []rune

	for _, char := range s {
		repeatCount := 1

		if char >= 'a' && char <= 'z' {
			// Calculate index: 'a' is 0, so we add 1 to get 1 repeat
			repeatCount = int(char - 'a' + 1)
		} else if char >= 'A' && char <= 'Z' {
			// Calculate index: 'A' is 0, so we add 1 to get 1 repeat
			repeatCount = int(char - 'A' + 1)
		}

		// Append the character to the result slice N times
		for i := 0; i < repeatCount; i++ {
			result = append(result, char)
		}
	}

	return string(result)
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
