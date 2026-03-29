package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 1. Check if the number of arguments is exactly 1
	if len(os.Args) != 2 {
		return
	}

	// 2. Parse the argument to an integer
	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 1 {
		return
	}

	divisor := 2
	first := true

	// 3. Find prime factors
	for num > 1 {
		if num%divisor == 0 {
			// Print the separator '*' if it's not the first factor found
			if !first {
				fmt.Print("*")
			}
			fmt.Print(divisor)
			
			// Reduce the number and reset 'first' flag
			num /= divisor
			first = false
		} else {
			// Move to the next potential divisor
			divisor++
		}
	}
	fmt.Println()
}
