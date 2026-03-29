package main

import (
	"fmt"
)

func RevConcatAlternate(slice1, slice2 []int) []int {
	var res []int
	i := len(slice1) - 1
	j := len(slice2) - 1

	// Phase 1: Exhaust the longer slice until lengths are equal
	for i > j {
		res = append(res, slice1[i])
		i--
	}
	for j > i {
		res = append(res, slice2[j])
		j--
	}

	// Phase 2: Slices are now "equal" in terms of remaining elements (i == j)
	// We alternate starting with slice1 as per the instructions
	for i >= 0 && j >= 0 {
		res = append(res, slice1[i])
		res = append(res, slice2[j])
		i--
		j--
	}

	return res
}

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
