package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(convertToBase7(100))
	fmt.Println(convertToBase7(-7))
	fmt.Println(convertToBase7(0))
	fmt.Println(convertToBase7(-1))
}

func convertToBase7(num int) (result string) {
	if num == 0 {
		return "0"
	}
	negative := false
	if num < 0 {
		num = -num
		negative = true
	}
	for i := 7; num > 0; i *= 7 {
		digit := num % i / (i / 7)
		result = strconv.Itoa(digit) + result
		num -= digit * i / 7
	}
	if negative {
		result = "-" + result
	}
	return
}
