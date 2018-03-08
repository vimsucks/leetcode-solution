package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	removeElement(nums, 3)
	fmt.Println(len(nums))
}

func removeElement(nums []int, val int) int {
	length := len(nums)
	i := 0
	for i < length {
		if nums[i] == val {
			nums = append(nums[0:i], nums[i+1:]...)
			length--
		} else {
			i++
		}
	}
	fmt.Println(nums)
	return len(nums)
}
