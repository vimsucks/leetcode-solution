package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MDLXX"))
	fmt.Println("================")
	fmt.Println(romanToInt("MCMLXXXVI"))
	fmt.Println("================")
	fmt.Println(romanToInt("MMMCMXCIX"))
}

var romanDigits = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) (result int) {
	for n := 10; len(s) > 0; n *= 10 {
		i := 0
		for i < len(s) && romanDigits[s[i]] >= n {
			i++
		}
		if i == len(s) {
			continue
		}
		var digits []int
		for _, d := range s[i:] {
			digits = append(digits, romanDigits[byte(d)])
		}
		if len(digits) >= 3 || len(digits) == 1 {
			for _, d := range digits {
				result += d
			}
		} else {
			if digits[0] >= digits[1] {
				result += digits[0] + digits[1]
			} else {
				result += digits[1] - digits[0]
			}
		}
		fmt.Println(digits)
		s = s[:i]
	}
	return
}
