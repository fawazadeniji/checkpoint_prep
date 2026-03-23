package main

import (
	"fmt"
)

func FindPrevPrime(nb int) int {
	// Primes are greater than 1
	if nb < 2 {
		return 0
	}

	// Check each number starting from nb downwards
	for i := nb; i >= 2; i-- {
		if isPrime(i) {
			return i
		}
	}

	return 0
}

// Helper function to check if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	// We only need to check up to the square root of n for efficiency
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}
