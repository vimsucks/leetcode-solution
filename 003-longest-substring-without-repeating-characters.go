package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	// substring without repeat character
	sub := ""
	// sub's chars' index in the original string
	index := make(map[rune]int)
	for i, c := range s {
		idx, ok := index[c]
		if !ok {
			index[c] = i
			sub += string(c)
		} else { // when repeatitive char occured
			// sub's first char's index in the original string
			subFirstCharIdx := index[rune(sub[0])]
			// delete the repetitive char and the chars before it from the index map and the string sub
			for j := 0; j <= idx-subFirstCharIdx; j++ {
				delete(index, rune(sub[j]))
			}
			sub = sub[idx+1-subFirstCharIdx:len(sub)] + string(c)
			index[c] = i
		}
		if len(sub) > maxLength {
			maxLength = len(sub)
		}
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring("a"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("busvutpwmu"))
	fmt.Println(lengthOfLongestSubstring("tvqnkvovks"))
	fmt.Println(lengthOfLongestSubstring("cdd"))
}
