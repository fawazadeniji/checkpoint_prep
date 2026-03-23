package main

import (
	"fmt"
)

func getUniques(s string) string { // "ahelloa"
	res := "" // ""
	for _, char := range s { // char -> 'o'
		found := false
		for _, r := range res { // r -> 'l'
			if r == char {
				found = true
				break
			}
		}
		if !found {
			res += string(char) // "helo"
		}
	}
	return res // "ahHe lo"
}

func WeAreUnique(str1, str2 string) int { // "hello there" "whats up"
	if str1 == "" && str2 == "" {
		return -1
	}

	count := 0

	u1 := getUniques(str1) // "helo tr"
	u2 := getUniques(str2) // "whats up"

	// Check characters in str1 not in str2
	for _, char := range u1 { // char -> 'h'
		isInStr2 := false
		for _, r := range str2 {
			if r == char {
				isInStr2 = true
				break
			}
		}
		if !isInStr2 {
			count++
		}
	}

	// Check characters in str2 not in str1
	for _, char := range u2 {
		isInStr1 := false
		for _, r := range str1 {
			if r == char {
				isInStr1 = true
				break
			}
		}
		if !isInStr1 {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(getUniques("whats up"))
	// fmt.Println(WeAreUnique("foo", "boo"))
	// fmt.Println(WeAreUnique("", ""))
	// fmt.Println(WeAreUnique("abc", "def"))
}
