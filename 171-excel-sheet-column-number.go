package main

import "fmt"

func main() {
	fmt.Println(titleToNumber("AA"))
	fmt.Println(titleToNumber("AB"))
}

func titleToNumber(s string) (number int) {
	for i, weight := len(s)-1, 1; i >= 0; i, weight = i-1, weight*26 {
		letter := s[i]
		number += weight * letterToNumber(letter)
	}
	return number
}

func letterToNumber(s byte) int {
	return int(s) - int('A') + 1
}
