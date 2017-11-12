package main

import "fmt"

func judgeCircle(moves string) bool {
	hor := 0
	ver := 0
	for _, c := range moves {
		switch c {
		case 'U':
			ver++
		case 'D':
			ver--
		case 'R':
			hor++
		case 'L':
			hor--
		}
	}
	return hor == 0 && ver == 0
}

func main() {
	fmt.Println(judgeCircle("UD"))
}
