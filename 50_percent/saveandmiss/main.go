package main

import (
	"fmt"
)

func SaveAndMiss(arg string, num int) string {
	// 1. Check for non-positive num
	if num <= 0 {
		return arg
	}

	result := ""
	save := true // Start with the "Save" phase

	// 2. Iterate through the string in chunks of 'num'
	for i := 0; i < len(arg); i += num {
		if save {
			// Determine the end of the current chunk
			end := i + num
			if end > len(arg) {
				end = len(arg)
			}
			// Add the chunk to our result
			result += arg[i:end]
		}
		
		// 3. Toggle the flag for the next chunk
		save = !save
	}

	return result
}

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}
