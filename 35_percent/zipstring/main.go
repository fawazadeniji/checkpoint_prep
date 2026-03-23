package main

import (
	"fmt"
	"strconv"
)

func ZipString(s string) string {
	if s == "" {
		return ""
	}

	res := ""
	for i := 0; i < len(s); {
		count := 0
		char := s[i]

		// Check how many times the current character repeats
		j := i
		for j < len(s) && s[j] == char {
			count++
			j++
		}

		// Concatenate the count and the character
		// strconv.Itoa converts the integer count to a string
		res += strconv.Itoa(count) + string(char)

		// Move the outer loop index to the next unique character
		i = j
	}

	return res
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
