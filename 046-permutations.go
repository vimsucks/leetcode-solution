package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	var permutations [][]int
	for i, v := range nums {
		permuteIter([]int{v}, remove(nums, i), &permutations)
	}
	return permutations
}

func permuteIter(permutation []int, nums []int, permutations *[][]int) {
	if  len(nums) == 0 {
		*permutations = append(*permutations, permutation)
	} else {
		for i, v := range nums {
			permuteIter(append(permutation, v), remove(nums, i), permutations)
		}
	}
}

func remove(slice []int, s int) []int {
	x := make([]int, len(slice))
	copy(x, slice)
	// append 如果容量不足会创建新的 slice，如果 cap 足够则会在原 slice 直接修改而不是创建新的 slice
	x = append(x[:s], slice[s+1:]...)
	return x
}
