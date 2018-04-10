package main

import "fmt"

func main() {
	p := fmt.Println
	p(trailingZeroes(30))
}

func trailingZeroes(n int) (result int) {
	for i := 5; n/i != 0; {
		result += n / i
		i *= 5
	}
	return result
}
