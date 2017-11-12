package main

import "fmt"

// 方法一
func hammingDistance01(x int, y int) int {
	distance := 0
	xored := x ^ y
	for xored != 0 {
		distance++
		xored = xored & (xored - 1)
	}
	return distance
}

// 方法二
func hammingDistance02(x int, y int) int {
	distance := 0
	xored := x ^ y
	for xored != 0 {
		if xored&1 == 1 {
			distance++
		}
		xored = xored >> 1
	}
	return distance
}

// 方法三
func hammingDistance03(x int, y int) int {
	distance := 0
	xored := x ^ y
	flag := 1
	for flag != 0 {
		if xored&flag != 0 {
			distance++
		}
		flag = flag << 1
	}
	return distance
}

// 方法四
func hammingDistance04(x int, y int) int {
	distance := 0
	xored := x ^ y
	for xored != 0 {
		if (xored>>1)<<1 != xored {
			distance++
		}
		xored = xored >> 1
	}
	return distance
}

func main() {
	x := 1
	y1 := 4
	y2 := -4
	format := "x=%d, y=%d, output=%d\n"
	fmt.Println("方法一：")
	fmt.Printf(format, x, y1, hammingDistance01(x, y1))
	fmt.Printf(format, x, y2, hammingDistance01(x, y2))
	fmt.Println("方法二：")
	fmt.Printf(format, x, y1, hammingDistance02(x, y1))
	fmt.Println("方法三：")
	fmt.Printf(format, x, y1, hammingDistance03(x, y1))
	fmt.Printf(format, x, y2, hammingDistance03(x, y2))
	fmt.Println("方法四：")
	fmt.Printf(format, x, y1, hammingDistance04(x, y1))
}

/*
Output:
方法一：
x=1, y=4, output=2
x=1, y=-4, output=63
方法二：
x=1, y=4, output=2
方法三：
x=1, y=4, output=2
x=1, y=-4, output=63
方法四：
x=1, y=4, output=2
*/
