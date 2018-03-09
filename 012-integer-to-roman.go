package main

import "fmt"

func main() {
	fmt.Println(intToRoman(1))
	fmt.Println(intToRoman(1986))
	fmt.Println(intToRoman(3999))
}

var (
	romanDigits = map[int](map[int]string){
		0: {1: "I", 5: "V"},
		1: {1: "X", 5: "L"},
		2: {1: "C", 5: "D"},
		3: {1: "M", 5: ""},
	}
)

func intToRoman(num int) string {
	var roman string
	for i := 0; num > 0; i++ {
		n := num % 10
		num /= 10
		if n < 4 {
			roman = repeatNTimes(romanDigits[i][1], n) + roman
			//TODO
		} else if n == 4 {
			roman = romanDigits[i][1] + romanDigits[i][5] + roman
		} else if n >= 5 && n < 9 {
			roman = romanDigits[i][5] + repeatNTimes(romanDigits[i][1], n-5) + roman
		} else if n == 9 {
			roman = romanDigits[i][1] + romanDigits[i+1][1] + roman
		}
	}
	return roman
}

func repeatNTimes(s string, n int) (result string) {
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}
