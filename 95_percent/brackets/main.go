package main

import (
	"fmt"
	"os"
)

func isBracketValid(s string) bool {
	var stack []rune

	// Map closing brackets to their corresponding opening brackets
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		// 1. If it's an opening bracket, push to stack
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			// 2. If it's a closing bracket
			if len(stack) == 0 {
				return false // Closing bracket with no opener
			}

			// Pop the last item
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Check if it matches the current closing bracket
			if top != pairs[char] {
				return false
			}
		}
		// Ignore all other characters
	}

	// 3. If stack is empty, all brackets were matched
	return len(stack) == 0
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	for _, arg := range args {
		if isBracketValid(arg) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}
