package main

import "fmt"

func main() {
	m1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(m1)
	rotate(m1)
	fmt.Println(m1)
	fmt.Println("=========")
	m2 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	fmt.Println(m2)
	rotate(m2)
	fmt.Println(m2)
}

func rotate(matrix [][]int) {
	n := len(matrix[0])
	var temp int
	for i := 0; i < n/2; i++ {
		for j := 0; j < n-2*i-1; j++ {
			temp, matrix[i+j][n-i-1] = matrix[i+j][n-i-1], matrix[i][i+j]
			temp, matrix[n-i-1][n-i-j-1] = matrix[n-i-1][n-i-j-1], temp
			temp, matrix[n-i-j-1][i] = matrix[n-i-j-1][i], temp
			matrix[i][i+j] = temp
		}
	}
}
