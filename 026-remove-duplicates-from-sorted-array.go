package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{2, 2, 3}))
	fmt.Println(removeDuplicates([]int{}))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	uniques := 0
	lastOne := nums[0] + 1
	length := len(nums)
	for i := 0; i < length; {
		if nums[i] != lastOne {
			uniques++
			lastOne = nums[i]
			i++
		} else {
			nums = append(nums[:i], nums[i+1:]...)
			length--
		}
	}
	return uniques
}
