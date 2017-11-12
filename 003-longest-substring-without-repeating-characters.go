package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	// sub 用来记录循环过程中没有出现过重复字符的 substring
	sub := ""
	// index 用来记录 sub 中的字符在原字符串中的下标
	index := make(map[rune]int)
	for i, c := range s {
		idx, ok := index[c]
		if !ok {
			index[c] = i
			sub += string(c)
		} else { // 当重复字符出现时
			// sub 字符串第一个字符在原字符串中的下标
			subFirstCharIdx := index[rune(sub[0])]
			// 在 index 字典中删除 sub 中出现重复的字符以前的所有字符，并截取 sub 字符串，去除出现重复的字符和之前的字符
			for j := 0; j <= idx-subFirstCharIdx; j++ {
				delete(index, rune(sub[j]))
			}
			sub = sub[idx+1-subFirstCharIdx:len(sub)] + string(c)
			// 更改为新出现字符的下标
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
