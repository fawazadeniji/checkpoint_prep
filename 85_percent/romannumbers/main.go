package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxRoman = 3999

type Roman struct {
	Value  int
	Symbol string
	Calc   string
}

var romanMap = []Roman{
	{1000, "M", "M"},
	{900, "CM", "(M-C)"},
	{500, "D", "D"},
	{400, "CD", "(D-C)"},
	{100, "C", "C"},
	{90, "XC", "(C-X)"},
	{50, "L", "L"},
	{40, "XL", "(L-X)"},
	{10, "X", "X"},
	{9, "IX", "(X-I)"},
	{5, "V", "V"},
	{4, "IV", "(V-I)"},
	{1, "I", "I"},
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num == 0 || num > maxRoman {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	roman, calc := toRoman(num)

	fmt.Println(calc)
	fmt.Println(roman)
}

// toRoman converts integer to Roman numeral + calc string
func toRoman(num int) (string, string) {
	var romanBuilder strings.Builder
	var calcParts []string

	for _, r := range romanMap {
		for num >= r.Value {
			romanBuilder.WriteString(r.Symbol)
			calcParts = append(calcParts, r.Calc)
			num -= r.Value
		}
	}

	return romanBuilder.String(), strings.Join(calcParts, "+")
}
