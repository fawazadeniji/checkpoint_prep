package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. Check for exactly 2 arguments
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	// Handle empty s1 (technically hidden in any string)
	if s1 == "" {
		fmt.Println()
		return
	}

	i := 0 // Pointer for s1
	
	// 2. Iterate through s2 to find characters of s1 in order
	for j := 0; j < len(s2) && i < len(s1); j++ {
		if s1[i] == s2[j] {
			i++
		}
	}

	// 3. If i matches the length of s1, the match is successful
	if i == len(s1) {
		fmt.Println(s1)
	}
}
