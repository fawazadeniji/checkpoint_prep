package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	str := os.Args[1]
	if str == "" {
		return
	}

	// 1. Manually Split: Find words and store them
	var words []string
	word := ""
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(str[i])
		}
	}
	// Don't forget the last word!
	if word != "" {
		words = append(words, word)
	}

	// 2. Reverse the slice of words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// 3. Print the result word by word
	for i, w := range words {
		// Print characters of the word
		for _, r := range w {
			z01.PrintRune(r)
		}
		// Print a space between words, but not after the last one
		if i < len(words)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
