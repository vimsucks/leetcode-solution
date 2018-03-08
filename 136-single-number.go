package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber([]int{3, 4, 7, 1, 4, 7, 3}))
}

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = result ^ nums[i]
	}
	return result
}
