package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 1. Basic Argument Count Validation
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	rawArr := os.Args[1]
	rawTarget := os.Args[2]

	// 2. Validate Target Sum
	target, err := strconv.Atoi(rawTarget)
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}

	// 3. Validate Array Format "[...]"
	if len(rawArr) < 2 || rawArr[0] != '[' || rawArr[len(rawArr)-1] != ']' {
		fmt.Println("Invalid input.")
		return
	}

	// 4. Parse the numbers
	content := rawArr[1 : len(rawArr)-1]
	// Handle empty brackets case
	if strings.TrimSpace(content) == "" {
		fmt.Println("No pairs found.")
		return
	}

	parts := strings.Split(content, ",")
	var nums []int
	for _, p := range parts {
		trimmed := strings.TrimSpace(p)
		if trimmed == "" {
			// This handles cases like "[1, , 2]" which are usually considered invalid
			fmt.Println("Invalid input.")
			return
		}
		n, err := strconv.Atoi(trimmed)
		if err != nil {
			fmt.Printf("Invalid number: %s\n", trimmed)
			return
		}
		nums = append(nums, n)
	}

	// 5. Find Pairs (Ensuring each index is used only once)
	var pairs [][]int
	used := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		// Skip if this index was already paired
		if used[i] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			// Skip if this second index was already paired
			if used[j] {
				continue
			}

			if nums[i]+nums[j] == target {
				pairs = append(pairs, []int{i, j})
				used[i] = true
				used[j] = true
				// Once i is paired, break to move to the next available i
				break
			}
		}
	}

	// 6. Output Results
	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
	} else {
		fmt.Printf("Pairs with sum %d: %v\n", target, pairs)
	}
}
