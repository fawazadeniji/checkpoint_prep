package main

import (
	"fmt"
)


func LastWord(s string) string {
	runes := []rune(s)
	n := len(runes)

	// Step 1: Find the end of the last word (ignore trailing spaces)
	end := -1
	for i := n - 1; i >= 0; i-- {
		if runes[i] != ' ' {
			end = i + 1 // end is exclusive in slicing
			break
		}
	}

	// If the string was empty or only spaces
	if end == -1 {
		return "\n"
	}

	// Step 2: Find the start of the last word
	start := 0
	for i := end - 1; i >= 0; i-- {
		if runes[i] == ' ' {
			start = i + 1
			break
		}
	}

	// Step 3: Extract and add newline
	return string(runes[start:end]) + "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
