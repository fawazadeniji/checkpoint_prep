package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. Validate argument count
	if len(os.Args) != 3 {
		// Instructions specify displaying a newline if args != 2
		if len(os.Args) != 1 {
			fmt.Println()
		}
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	// 2. Track characters to avoid duplicates
	seen := make(map[rune]bool)

	// 3. Process the first string
	for _, char := range s1 {
		if !seen[char] {
			fmt.Print(string(char))
			seen[char] = true
		}
	}

	// 4. Process the second string
	for _, char := range s2 {
		if !seen[char] {
			fmt.Print(string(char))
			seen[char] = true
		}
	}

	// 5. Final newline
	fmt.Println()
}
