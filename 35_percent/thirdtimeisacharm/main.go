package main

import (
	"fmt"
)

func ThirdTimeIsACharm(str string) string {
	// If the string is empty or shorter than 3 characters, return only a newline
	if len(str) < 3 {
		return "\n"
	}

	var res string

	// Start at index 2 (the 3rd character)
	// Increment by 3 each time to hit every 3rd character
	for i := 2; i < len(str); i += 3 {
		res += string(str[i])
	}

	return res + "\n"
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}
