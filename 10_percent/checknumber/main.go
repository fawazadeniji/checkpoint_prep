package main

import (
	"fmt"
)

func CheckNumber(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}
