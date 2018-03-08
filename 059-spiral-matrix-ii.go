package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(4))
	fmt.Println(generateMatrix(5))
}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i, _ := range matrix {
		matrix[i] = make([]int, n)
	}
	generateMatrixIter(n, 0, 1, &matrix)
	if n%2 == 1 {
		matrix[n/2][n/2] = n * n
	}
	return matrix
}

func generateMatrixIter(n, layer int, element int, matrix *[][]int) {
	if layer > (n+1)/2 {
		return
	}

	sideLength := n - 2*layer - 1

	for i := 0; i < sideLength; i++ {
		(*matrix)[layer][layer+i] = element
		element++
	}

	for i := 0; i < sideLength; i++ {
		(*matrix)[layer+i][layer+sideLength] = element
		element++
	}

	for i := sideLength; i > 0; i-- {
		(*matrix)[layer+sideLength][layer+i] = element
		element++
	}

	for i := sideLength; i > 0; i-- {
		(*matrix)[layer+i][layer] = element
		element++
	}

	generateMatrixIter(n, layer+1, element, matrix)
}
