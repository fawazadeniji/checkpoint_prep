package main

import (
	"fmt"
	"os"
)

func ReverseStrCap(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		// First, default the character to lowercase if it's a letter
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 32
		}

		// Check if it's the last letter of a word
		// Condition: Current is a letter AND (Next is a space OR Next is end of string)
		isEndOfWord := false
		if (runes[i] >= 'a' && runes[i] <= 'z') || (runes[i] >= 'A' && runes[i] <= 'Z') {
			if i == len(runes)-1 || runes[i+1] == ' ' {
				isEndOfWord = true
			}
		}

		// If it's the end of the word, convert to uppercase
		if isEndOfWord {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] -= 32
			}
		}
	}
	return string(runes)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	for _, arg := range args {
		fmt.Println(ReverseStrCap(arg))
	}
}
