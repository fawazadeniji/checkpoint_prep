package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. Check if the number of arguments is exactly 2
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	// 2. Handle the empty string case (always hidden)
	if s1 == "" {
		fmt.Println("1")
		return
	}

	i := 0 // index for s1
	j := 0 // index for s2

	// 3. Iterate through s2 to find characters of s1
	for j < len(s2) && i < len(s1) {
		if s1[i] == s2[j] {
			i++ // Found the current character, move to the next one in s1
		}
		j++ // Always move forward in s2
	}

	// 4. If i reached the end of s1, all characters were found in order
	if i == len(s1) {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}