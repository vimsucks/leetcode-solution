package main

import "fmt"

func judgeCircle(moves string) bool {
	x := 0
	y := 0
	for _, c := range moves {
		switch c {
		case 'U':
			y++
		case 'D':
			y--
		case 'R':
			x++
		case 'L':
			x--
		}
	}
	return x == 0 && y == 0
}

func main() {
	fmt.Println(judgeCircle("UD"))
}
