package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{1, 2, 3, 4}))
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
}

func subsets(nums []int) [][]int {
	sort.Ints(nums)
	var sets [][]int
	sets = append(sets, []int{})
	for i, v := range nums {
		fmt.Println(nums[i+1:], []int{v})
		subsetsIter(nums[i+1:], []int{v}, &sets)
	}
	return sets
}

func subsetsIter(nums []int, set []int, sets *[][]int) {
	*sets = append(*sets, set)

	if len(nums) == 0 {
		return
	}

	for i, v := range nums {
		s := make([]int, len(set))
		copy(s, set)
		subsetsIter(nums[i+1:], append(s, v), sets)
	}
}
