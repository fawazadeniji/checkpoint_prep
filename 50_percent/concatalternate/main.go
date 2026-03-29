package main

import (
	"fmt"
)

func ConcatAlternate(slice1, slice2 []int) []int {
	var first, second []int

	// Determine which slice to start with based on length
	if len(slice2) > len(slice1) {
		first = slice2
		second = slice1
	} else {
		first = slice1
		second = slice2
	}

	res := make([]int, 0, len(slice1)+len(slice2))
	
	// Find the minimum length to know how many pairs we can make
	minLen := len(second)

	// Alternating part
	for i := 0; i < minLen; i++ {
		res = append(res, first[i])
		res = append(res, second[i])
	}

	// Add the remaining elements from the longer slice
	res = append(res, first[minLen:]...)

	return res
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
