package main

import (
	"fmt"
)

func FishAndChips(n int) string {
	switch {
	case n < 0:
		return "error: number is negative"
	case n%2 == 0 && n%3 == 0:
		return "fish and chips"
	case n%2 == 0:
		return "fish"
	case n%3 == 0:
		return "chips"
	default:
		return "error: non divisible"
	}
}

func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
}
