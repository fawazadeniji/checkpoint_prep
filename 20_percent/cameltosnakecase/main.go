package main

import (
	"fmt"
)

func IsUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func IsLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	runes := []rune(s)
	n := len(runes)

	// --- 1. Validation Logic ---
	for i := 0; i < n; i++ {
		// Rule: No numbers or punctuation
		if !IsUpper(runes[i]) && !IsLower(runes[i]) {
			return s
		}
		// Rule: No two capitalized letters follow each other
		if i > 0 && IsUpper(runes[i]) && IsUpper(runes[i-1]) {
			return s
		}
	}
	// Rule: Word does not end on a capitalized letter
	if IsUpper(runes[n-1]) {
		return s
	}

	// --- 2. Transformation Logic ---
	var result []rune
	for i := 0; i < n; i++ {
		// Insert underscore if current is Upper and not the first character
		if i > 0 && IsUpper(runes[i]) {
			result = append(result, '_')
		}
		result = append(result, runes[i])
	}

	return string(result)
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
