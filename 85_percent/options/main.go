package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	// 1. If no arguments are provided, print the valid options and exit
	if len(args) == 0 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}

	// 2. First Pass: Check for the -h flag priority.
	// If any argument starts with '-h', it takes precedence over everything else.
	for _, arg := range args {
		if len(arg) >= 2 && arg[0] == '-' && arg[1] == 'h' {
			fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
			return
		}
	}

	var options uint32 = 0

	// 3. Second Pass: Validate arguments and set bits.
	for _, arg := range args {
		// If an argument doesn't start with '-' or is just a '-' string
		if len(arg) < 2 || arg[0] != '-' {
			fmt.Println("Invalid Option")
			return
		}

		// Iterate through the characters of the option
		for i := 1; i < len(arg); i++ {
			char := arg[i]
			// If it's not a lowercase letter, it's invalid
			if char < 'a' || char > 'z' {
				fmt.Println("Invalid Option")
				return
			}
			// Shift 1 by the alphabetical index (a=0, b=1, ... z=25) and apply bitwise OR
			options |= 1 << (char - 'a')
		}
	}

	// 4. Print the 32-bit integer in blocks of 8 bits
	for i := 31; i >= 0; i-- {
		// Check if the bit at the current index is 1 or 0
		if (options & (1 << i)) != 0 {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}

		// Add a space after every 8 bits, but not at the very end
		if i%8 == 0 && i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
