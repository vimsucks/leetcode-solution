package main

import (
	"fmt"
	"strings"
)

var digit2Letter = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func main() {
	fmt.Println(len(letterCombinations("")))
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("234"))
}

func letterCombinations(digits string) (combinations []string) {
	if len(digits) == 0 {
		return
	}
	var letters []string
	for _, digit := range strings.Split(digits, "") {
		letters = append(letters, digit2Letter[digit])
	}
	letterCombinationsIter("", letters, 0, &combinations)
	return
}

func letterCombinationsIter(combination string, letters []string, currentIndex int, combinations *[]string) {
	if currentIndex == len(letters) {
		*combinations = append(*combinations, combination)
	} else {
		for _, l := range strings.Split(letters[currentIndex], "") {
			letterCombinationsIter(combination+l, letters, currentIndex+1, combinations)
		}
	}
}
