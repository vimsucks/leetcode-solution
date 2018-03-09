package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"123", "1", "12"}))
	fmt.Println(longestCommonPrefix([]string{"", ""}))
	fmt.Println(longestCommonPrefix([]string{}))
	fmt.Println(longestCommonPrefix([]string{"123", "1", "1"}))
	fmt.Println(longestCommonPrefix([]string{"a", "a", "b"}))
	fmt.Println(longestCommonPrefix([]string{"a"}))
}

func longestCommonPrefix(strs []string) string {
	// sort by str length
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}

	var commonPrefix string
	for i := 0; i < len(strs[0]); i++ {
		equal := true
		prefix := strs[0][:i+1]
		for j := 1; j < len(strs); j++ {
			if prefix != strs[j][:i+1] {
				equal = false
				break
			}
		}
		if equal {
			commonPrefix = prefix
		}
	}
	return commonPrefix
}
