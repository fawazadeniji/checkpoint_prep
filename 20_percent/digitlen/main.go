package main

import (
	"fmt"
)

func DigitLen(n, base int) int {
	// Check if base is within the allowed range [2, 36]
	if base < 2 || base > 36 {
		return -1
	}

	// Handle negative numbers by reversing the sign
	if n < 0 {
		// Note: Be careful with math.MinInt64, but for standard 
		// int exercises, n = -n is usually sufficient.
		n = -n
	}

	// Counter for divisions
	count := 0

	// Special case for 0: it is already "at zero", 
	// but usually occupies 1 digit space.
	if n == 0 {
		return 1
	}

	// Divide until n reaches zero
	for n > 0 {
		n = n / base
		count++
	}

	return count
}


func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}
