// TODO: TLE
package main

import "fmt"

func main() {
	var board [][]byte
	board = [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'X', 'X'},
	}
	printArray(board)
	solve(board)
	printArray(board)

	board = [][]byte{
		{'O', 'X', 'X', 'O', 'X'},
		{'X', 'O', 'O', 'X', 'O'},
		{'X', 'O', 'X', 'O', 'X'},
		{'O', 'X', 'O', 'O', 'O'},
		{'X', 'X', 'O', 'X', 'O'},
	}
	printArray(board)
	solve(board)
	printArray(board)

	board = [][]byte{
		{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'O', 'O'},
		{'X', 'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X', 'O', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X', 'X', 'O', 'X', 'X', 'X', 'O'},
		{'O', 'X', 'X', 'X', 'O', 'X', 'O', 'X', 'O', 'X'},
		{'X', 'X', 'O', 'X', 'X', 'O', 'O', 'X', 'X', 'X'},
		{'O', 'X', 'X', 'O', 'O', 'X', 'O', 'X', 'X', 'O'},
		{'O', 'X', 'X', 'X', 'X', 'X', 'O', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X', 'X', 'O', 'X', 'X', 'O', 'O'},
		{'X', 'X', 'X', 'O', 'O', 'X', 'O', 'X', 'X', 'O'}}
	printArray(board)
	solve(board)
	printArray(board)

}

func printArray(a [][]byte) {
	for _, bs := range a {
		fmt.Println(string(bs))
	}
	fmt.Println("===========================")
}

const (
	FRIENDLY = 'O'
	ENEMY    = 'X'
)

func solve(board [][]byte) {
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == FRIENDLY {
				if !isNotSurrounded(board, i, j, visited) {
					board[i][j] = ENEMY
				}
			}
		}
	}
}

func isNotSurrounded(board [][]byte, i, j int, visited [][]bool) bool {
	if i == 0 || i == len(board)-1 {
		return true
	}
	if j == 0 || j == len(board[0])-1 {
		return true
	}

	visited[i][j] = true
	notSurrounded := false

	if visited[i-1][j] == false && board[i-1][j] == FRIENDLY {
		notSurrounded = notSurrounded || isNotSurrounded(board, i-1, j, visited)
	}
	if notSurrounded {
		visited[i][j] = false
		return true
	}

	if visited[i+1][j] == false && board[i+1][j] == FRIENDLY {
		notSurrounded = notSurrounded || isNotSurrounded(board, i+1, j, visited)
	}
	if notSurrounded {
		visited[i][j] = false
		return true
	}

	if visited[i][j-1] == false && board[i][j-1] == FRIENDLY {
		notSurrounded = notSurrounded || isNotSurrounded(board, i, j-1, visited)
	}
	if notSurrounded {
		visited[i][j] = false
		return true
	}

	if visited[i][j+1] == false && board[i][j+1] == FRIENDLY {
		notSurrounded = notSurrounded || isNotSurrounded(board, i, j+1, visited)
	}

	visited[i][j] = false
	return notSurrounded
}
