package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rowChars := [][]rune{}
	for i := 0; i < numRows; i++ {
		rowChars = append(rowChars, []rune{})
	}
	numPerGroup := numRows*2 - 2
	i := 0
	for _, c := range s {
		if i == numPerGroup {
			i = 0
		}
		if i < numRows {
			rowChars[i] = append(rowChars[i], c)
		} else {
			j := 2*numRows - 2 - i
			rowChars[j] = append(rowChars[j], c)
		}
		i++
	}
	output := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(rowChars[i]); j++ {
			output += string(rowChars[i][j])
		}
	}
	return output
}

func main() {
	println(convert("PAYPALISHIRING", 1))
	println(convert("PAYPALISHIRING", 2))
	println(convert("PAYPALISHIRING", 3))
}
