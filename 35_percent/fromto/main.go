package main

import (
	"fmt"
)

func FromTo(from int, to int) string {
	// 1. Validation
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	var result string
	current := from

	for {
		// 2. Convert integer to two-digit string manually
		// Tens digit: (n / 10) + '0'
		// Units digit: (n % 10) + '0'
		result += string(rune(current/10 + '0'))
		result += string(rune(current%10 + '0'))

		// 3. Exit condition
		if current == to {
			break
		}

		// 4. Add separator
		result += ", "

		// 5. Increment or Decrement
		if from < to {
			current++
		} else {
			current--
		}
	}

	return result + "\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}
