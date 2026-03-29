package main

import (
	"fmt"
)

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	// Find the decimal point
	dotIdx := -1
	for i, r := range dec {
		if r == '.' {
			dotIdx = i
			break
		}
	}

	// Case 1: No decimal point
	if dotIdx == -1 {
		return dec + "\n"
	}

	// Case 2: Only a zero after the decimal point (e.g., "15.0")
	// Or check if the decimal part is strictly numeric
	before := dec[:dotIdx]
	after := dec[dotIdx+1:]

	if after == "0" || after == "" { 
		return before + "\n"
	}

	// Case 3: Check if the string is a valid "float" format (numbers and optional minus)
	// If it contains non-digits in the 'after' or 'before' part, return original
	fullNoDot := before + after
	for i, r := range fullNoDot {
		if (r < '0' || r > '9') && !(i == 0 && r == '-') {
			return dec + "\n"
		}
	}

	// Clean up leading zeros for positive numbers, but keep one if it's "0"
	// and handle negative signs correctly.
	res := ""
	isNegative := false
	if fullNoDot[0] == '-' {
		isNegative = true
		fullNoDot = fullNoDot[1:]
	}

	// Remove leading zeros
	start := 0
	for start < len(fullNoDot)-1 && fullNoDot[start] == '0' {
		start++
	}
	
	if isNegative {
		res = "-"
	}
	res += fullNoDot[start:]

	return res + "\n"
}

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
