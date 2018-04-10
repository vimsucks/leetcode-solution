package main

import "fmt"

func main() {
	case1 := []int{1, 2, 3, 1}
	fmt.Println(findPeakElement(case1))
	case2 := []int{3, 2, 1}
	fmt.Println(findPeakElement(case2))
	case3 := []int{1}
	fmt.Println(findPeakElement(case3))
	case4 := []int{1, 2}
	fmt.Println(findPeakElement(case4))
	case5 := []int{2, 1}
	fmt.Println(findPeakElement(case5))
}

func findPeakElement(nums []int) int {
	i := 0
	for i < len(nums)-1 {
		if i == 0 || nums[i] > nums[i-1] {
			if nums[i] > nums[i+1] {
				return i
			} else {
				i++
			}
		} else {
			if nums[i] > nums[i+1] {
				i += 2
			} else {
				i++
			}
		}
	}
	return i
}
