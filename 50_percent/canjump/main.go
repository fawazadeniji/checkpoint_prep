package main

import (
	"fmt"
)
//        0  1  2  3  4  5  
// []uint{1, 3, 5, 1, 2, 3}
func CanJump(steps []uint) bool {
	// If the array is empty, return false
	if len(steps) == 0 {
		return false
	}

	lastIndex := len(steps) - 1 // 5
	currentIndex := 0           // 6

	// Continue jumping as long as we haven't reached or passed the last index
	for currentIndex < lastIndex {
		jumpSize := int(steps[currentIndex]) // 2

		// If the jump size is 0 and we aren't at the end, we are stuck
		if jumpSize == 0 {
			return false
		}

		// Move exactly 'jumpSize' steps forward
		currentIndex += jumpSize // 6
	}

	// Return true only if we landed exactly on the last index
	return currentIndex == lastIndex// 6 == 5 --> false
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}
