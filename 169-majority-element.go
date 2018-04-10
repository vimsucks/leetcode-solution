package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(majorityElement([]int{1}))
	fmt.Println(majorityElement([]int{6, 5, 5}))
	fmt.Println(majorityElement([]int{3, 2, 3}))
}

func majorityElement(nums []int) int {
	sort.Ints(nums)
	n := (len(nums) + 1) / 2 // 向上取整
	num := nums[0]
	count := 0
	for _, v := range nums {
		if v == num {
			count++
		} else {
			num = v
			count = 1
		}

		if count >= n {
			break
		}
	}
	return num
}
