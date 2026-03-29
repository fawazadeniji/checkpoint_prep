package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 1. Exactly one argument required
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}

	// 2. Split by whitespace and ignore empty strings (handles extra spaces)
	tokens := strings.Fields(os.Args[1])
	var stack []int

	for _, token := range tokens {
		// Try to parse as a number
		val, err := strconv.Atoi(token)

		if err == nil {
			// It's a number: Push to stack
			stack = append(stack, val)
		} else {
			// It's an operator: Check if we have 2 operands
			if len(stack) < 2 {
				fmt.Println("Error")
				return
			}

			// Pop the two operands
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			// Perform calculation
			var res int
			switch token {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				if b == 0 { // Division by zero check
					fmt.Println("Error")
					return
				}
				res = a / b
			case "%":
				if b == 0 {
					fmt.Println("Error")
					return
				}
				res = a % b
			default:
				// Invalid token found
				fmt.Println("Error")
				return
			}
			// Push result back
			stack = append(stack, res)
		}
	}

	// 3. Result is valid only if exactly one value remains
	if len(stack) == 1 {
		fmt.Println(stack[0])
	} else {
		fmt.Println("Error")
	}
}
