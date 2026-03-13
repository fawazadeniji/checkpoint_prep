package main

import (
	"fmt"
)

func RectPerimeter(length, width int) int {
	if length < 0 || width < 0 {
		return -1
	}
	return 2 * (length + width)
}

func main() {
	fmt.Println(RectPerimeter(10, 2))
	fmt.Println(RectPerimeter(434343, 898989))
	fmt.Println(RectPerimeter(10, -2))
}
