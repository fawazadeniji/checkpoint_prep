package main

import (
	"fmt"
)

func HashCode(dec string) string {
	n := len(dec)
	result := make([]rune, n)

	for i, char := range dec {
		// 1. Calculate the raw hash: (ASCII + size) % 127
		// We use int for the calculation to avoid overflow issues
		hashedVal := (int(char) + n) % 127

		// 2. Check for "unprintable" status. 
		// ASCII values 0-31 and 127 are typically unprintable.
		if hashedVal < 32 || hashedVal == 127 {
			hashedVal += 33
		}

		result[i] = rune(hashedVal)
	}

	return string(result)
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
