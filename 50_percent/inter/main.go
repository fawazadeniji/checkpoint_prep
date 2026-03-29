package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. Validate arguments
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	// 2. Map characters in s2 for O(1) lookup
	inS2 := make(map[rune]bool)
	for _, char := range s2 {
		inS2[char] = true
	}

	// 3. Keep track of what we've already printed
	alreadyPrinted := make(map[rune]bool)

	// 4. Iterate through s1 to maintain the required order
	for _, char := range s1 {
		// If it exists in s2 AND we haven't printed it yet
		if inS2[char] && !alreadyPrinted[char] {
			fmt.Print(string(char))
			alreadyPrinted[char] = true
		}
	}

	// 5. Final newline
	fmt.Println()
}
