package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}

func findWords(words []string) []string {
	alphabet := map[rune]int{
		'q': 0,
		'w': 0,
		'e': 0,
		'r': 0,
		't': 0,
		'y': 0,
		'u': 0,
		'i': 0,
		'o': 0,
		'p': 0,
		'a': 1,
		's': 1,
		'd': 1,
		'f': 1,
		'g': 1,
		'h': 1,
		'j': 1,
		'k': 1,
		'l': 1,
		'z': 2,
		'x': 2,
		'c': 2,
		'v': 2,
		'b': 2,
		'n': 2,
		'm': 2,
	}
	var results []string
	for _, s := range words {
		r := []rune(strings.ToLower(s))
		typable := true
		row := alphabet[r[0]]
		for i := 1; i < len(r); i++ {
			if alphabet[r[i]] != row {
				typable = false
			}
		}
		if typable {
			results = append(results, s)
		}
	}
	return results
}
