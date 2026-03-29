package main

import (
	"fmt"
)

func ItoaBase(value, base int) string {
	// Handle the zero edge case immediately
	if value == 0 {
		return "0"
	}

	// Define the character set for bases up to 16
	charset := "0123456789ABCDEF"
	
	isNegative := false
	var n uint

	// Handle negative numbers
	// We cast to uint to safely handle the minimum integer limit (e.g., -9223372036854775808)
	// without causing an overflow when flipping the sign.
	if value < 0 {
		isNegative = true
		n = uint(-value)
	} else {
		n = uint(value)
	}

	var res []byte
	uBase := uint(base)

	// Extract digits from right to left
	for n > 0 {
		rem := n % uBase
		res = append(res, charset[rem])
		n /= uBase
	}

	// Append the minus sign if the original value was negative
	if isNegative {
		res = append(res, '-')
	}

	// The digits were extracted backwards, so we need to reverse the byte slice
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}

func main() {
	fmt.Println(ItoaBase(10, 2))
	fmt.Println(ItoaBase(255, 16))
	fmt.Println(ItoaBase(-42, 4))
	fmt.Println(ItoaBase(123, 10))
	fmt.Println(ItoaBase(0, 8))
	fmt.Println(ItoaBase(255, 2))
	fmt.Println(ItoaBase(-255, 16))
	fmt.Println(ItoaBase(15, 16))
	fmt.Println(ItoaBase(10, 4))
	fmt.Println(ItoaBase(255, 10))
}
