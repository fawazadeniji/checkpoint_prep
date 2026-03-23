package main

import (
	"fmt"
)

func Itoa(n int) string {
	// Handle 0 explicitly
	if n == 0 {
		return "0"
	}

	var res string
	isNegative := false

	// Handle negative numbers
	if n < 0 {
		isNegative = true
		// Note: For the smallest int (minInt), 
		// direct negation can overflow, but for basic 
		// exercises this standard approach is usually expected.
		n = -n
	}

	// Extract digits from right to left
	for n > 0 {
		digit := n % 10
		// Convert digit to its ASCII character
		res = string(rune(digit+'0')) + res
		n /= 10
	}

	// Prepend minus sign if necessary
	if isNegative {
		res = "-" + res
	}

	return res
}

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
