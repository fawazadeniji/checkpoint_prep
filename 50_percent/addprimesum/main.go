package main

import (
	"fmt"
	"os"
)

func main() { // 
	args := os.Args[1:]// "7"
	if len(args) != 1 {
		fmt.Println("0")
		return
	}
	num := atoi(args[0])// 7
	if num < 0 {
		fmt.Println("0")
		return
	}

	sum := 0 // 17
	for i := 0; i <= num; i++ {
		sum += isPrime(i) // 7
	}
	fmt.Println(sum)
}

func isPrime(n int) int {
	if n < 2 {
		return 0
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return 0
		}
	}
	return n
}

func atoi(s string) int {
	var num int
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		num = num*10 + int(s[i]-'0')
	}
	return num
}
