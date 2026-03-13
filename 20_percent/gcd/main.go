package main

import (
	"fmt"
)

func Gcd(a, b uint) uint {
	// If any input is 0, return 0 as per instructions
	if a == 0 || b == 0 {
		return 0
	}

	// Euclidean Algorithm
	for b != 0 {
		// b becomes the remainder of a divided by b
		// a becomes the old value of b
		a, b = b, a%b
	}

	return a
}

func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}
