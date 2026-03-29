package main

import (
	"fmt"
)

func ConcatSlice(slice1, slice2 []int) []int {
	// We append all elements of slice2 to slice1.
	// The '...' suffix is essential to treat slice2 as a list of arguments.
	return append(slice1, slice2...)
}

func ConcatSlice2(slice1, slice2 []int) []int {
	// Create a new slice with the exact combined length
	newSize := len(slice1) + len(slice2)
	res := make([]int, newSize)

	// Copy the first slice into the start of the result
	copy(res, slice1)
	
	// Copy the second slice into the result starting after slice1
	copy(res[len(slice1):], slice2)

	return res
}

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}
