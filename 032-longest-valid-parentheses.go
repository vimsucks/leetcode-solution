/*
	TODO: TO BE OPTIMIZED
*/

package main

import (
	"fmt"
)

func main() {
	longestValidParentheses("(()")
	longestValidParentheses("()")
	longestValidParentheses("())")
	longestValidParentheses("()(()")
	longestValidParentheses(")()(((())))(")
	longestValidParentheses("(()(((()")
}

func longestValidParentheses(s string) int {
	biggest := 0
	longestValidParenthesesIter(s, &biggest)
	fmt.Println(biggest)
	fmt.Println("==========================")
	return biggest
}

func longestValidParenthesesIter(s string, biggest *int) {
	for len(s) > 0 && s[0] == ')' {
		s = s[1:]
	}
	for len(s) > 0 && s[len(s)-1] == '(' {
		s = s[:len(s)-1]
	}

	var stack []byte
	i := 0
	for ; i < len(s); i++ {
		p := s[i]

		if len(stack) == 0 {
			if i > 0 && i > *biggest {
				*biggest = i
			}

			if p == '(' {
				push(&stack, p)
			} else {
				if len(s) > i {
					longestValidParenthesesIter(s[i+1:], biggest)
				}
				return
			}

		} else {
			if top(stack) == p {
				push(&stack, p)
			} else if top(stack) == '(' && p == ')' {
				pop(&stack)
			} else {
				break
			}
		}
	}

	if len(stack) == 0 && i == len(s) && i > *biggest {
		*biggest = i
	}

	if len(s) > 0 {
		longestValidParenthesesIter(s[1:], biggest)
	}
	return
}

func pop(s *[]byte) {
	*s = (*s)[:len(*s)-1]
}

func push(s *[]byte, p byte) {
	*s = append(*s, p)
}

func top(s []byte) byte {
	return s[len(s)-1]
}
