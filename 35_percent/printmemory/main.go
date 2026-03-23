package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	hexDigit := "0123456789abcdef"

	// 1. Print Hexadecimal values
	for i, b := range arr {
		// Calculate the two hex characters for the byte
		first := hexDigit[b/16]
		second := hexDigit[b%16]

		z01.PrintRune(rune(first))
		z01.PrintRune(rune(second))

		// Formatting spaces and newlines
		if (i+1)%4 == 0 || i == 9 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}

	// 2. Print ASCII graphic characters
	for _, b := range arr {
		if b >= 32 && b <= 126 {
			z01.PrintRune(rune(b))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
