package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// 1. Check argument count
	if len(os.Args) != 2 {
		if len(os.Args) == 1 { // If no args, just a newline
			z01.PrintRune('\n')
		}
		return
	}

	str := os.Args[1]
	var words []string
	currentWord := ""

	// 2. Extract words (ignoring multiple spaces)
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
	// Add the last word if exists
	if currentWord != "" {
		words = append(words, currentWord)
	}

	// 3. Check if we actually found any words
	if len(words) == 0 {
		z01.PrintRune('\n')
		return
	}

	// 4. Rotate: Print from the second word to the end, then the first word
	for i := 1; i < len(words); i++ {
		printString(words[i])
		z01.PrintRune(' ')
	}

	// Print the first word last
	printString(words[0])
	z01.PrintRune('\n')
}

// Helper to print strings using PrintRune
func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
