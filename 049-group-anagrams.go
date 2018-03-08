package main

import (
	"strings"
	"sort"
	"fmt"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupAnagrams(strs []string) (anagrams [][]string) {
	anagramsMap := make(map[string][]string)
	for _, s := range strs {
		sortedS := sortString(s)
		anagramsMap[sortedS] = append(anagramsMap[sortedS], s)
	}
	for _, v := range anagramsMap {
		anagrams = append(anagrams, v)
	}
	return
}
