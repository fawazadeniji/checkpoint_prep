package main

import (
	"fmt"
)

func Chunk(slice []int, size int) {
	// 1. Handle size 0 as per instructions
	if size == 0 {
		fmt.Println()
		return
	}

	// 2. Handle empty slice
	if len(slice) == 0 {
		fmt.Println("[]")
		return
	}

	var res [][]int

	// 3. Loop through the slice with a step of 'size'
	for i := 0; i < len(slice); i += size {
		end := i + size

		// 4. Ensure we don't go out of bounds for the last chunk
		if end > len(slice) {
			end = len(slice)
		}

		// Append the sub-slice to our result
		res = append(res, slice[i:end])
	}

	// 5. Print the final 2D slice
	fmt.Println(res)
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
