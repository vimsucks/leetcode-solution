package main

import "fmt"

func generateParenthesis(n int) []string {
	var output []string
	generateParenthesisIter("", 0, 0, n, &output, []byte{})
	return output
}

/*
   temp 是一个简单的栈，仅用于存放 '(' 和 ')'
   当放入 '(' 时，不做特殊处理；而放入 ')' 时，若当前栈顶元素为 '('，则删除栈顶元素，并不放入 ')'
   随后若 temp 的长度为 0，则该括号串是合法的
*/
func generateParenthesisIter(parens string, left int, right int, n int, output *[]string, temp []byte) {
	if len(parens) == 0 {
		generateParenthesisIter("(", 1, 0, n, output, append(temp, '('))
		return
	}
	if len(parens) == 2*n {
		if len(temp) == 0 {
			*output = append(*output, parens)
		}
		return
	}
	if left < n {
		generateParenthesisIter(parens+"(", left+1, right, n, output, append(temp, '('))
	}
	if right < n && len(temp) > 0 && temp[len(temp)-1] == '(' {
		generateParenthesisIter(parens+")", left, right+1, n, output, temp[:len(temp)-1])
	}
}

func main() {
	output := generateParenthesis(5)
	for _, v := range output {
		fmt.Println(v)
	}
}
