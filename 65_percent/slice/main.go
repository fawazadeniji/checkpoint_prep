package main

import (
	"fmt"
)

func Slice(a []string, nbrs ...int) []string {
	l := len(a)
	if l == 0 {
		return nil
	}

	start := 0
	end := l

	// Assign start index
	if len(nbrs) >= 1 {
		start = nbrs[0]
		if start < 0 {
			start = l + start
		}
	}

	// Assign end index
	if len(nbrs) >= 2 {
		end = nbrs[1]
		if end < 0 {
			end = l + end
		}
	}

	// Boundary Guardrails
	if start < 0 {
		start = 0
	}
	if end > l {
		end = l
	}

	// If start is greater than end, or start exceeds length, return nil
	if start >= end || start >= l {
		return nil
	}

	return a[start:end]
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}
