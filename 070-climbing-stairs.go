package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(10))
	fmt.Println(climbStairs(35))
	fmt.Println(climbStairs(44))
}

func climbStairs(n int) int {
	ways := 0
	one := n % 2
	two := n / 2
	for two >= 0 {
		ways += C(one+two, one)
		one += 2
		two -= 1
	}
	return ways
}

// 排列组合
func C(x, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x - i
		result /= i + 1
	}
	return result
}
