package main

import (
	"fmt"
)

func PrintIf(s string) string {
	if len(s) >= 3 || len(s) == 0 {
		return "G\n"
	}
	return "Invalid Input\n"
}

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}
