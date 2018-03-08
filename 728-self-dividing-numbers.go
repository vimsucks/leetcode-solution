package main

import "fmt"

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}

func selfDividingNumbers(left int, right int) []int {
	results := []int{}
	for ; left <= right; left++ {
		selfDividable := true
		operand, temp := left, left
		for temp != 0 {
			digit := temp % 10
			temp /= 10
			if digit == 0 || operand%digit != 0 {
				selfDividable = false
				break
			}
		}
		if selfDividable {
			results = append(results, operand)
		}
	}
	return results
}
