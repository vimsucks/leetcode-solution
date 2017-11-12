package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersIter(l1, l2, 0)
}

func addTwoNumbersIter(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		if carry != 0 {
			return &ListNode{Val: carry}
		}
		return nil
	}
	l := &ListNode{Val: 0}
	if l1 != nil && l2 != nil {
		l.Val += l1.Val + l2.Val + carry
		carry = l.Val / 10
		l.Val %= 10
		l.Next = addTwoNumbersIter(l1.Next, l2.Next, carry)
	} else {
		if l1 == nil {
			l.Val += l2.Val + carry
			carry = l.Val / 10
			l.Val %= 10
			l.Next = addTwoNumbersIter(nil, l2.Next, carry)
		} else {
			l.Val += l1.Val + carry
			carry = l.Val / 10
			l.Val %= 10
			l.Next = addTwoNumbersIter(l1.Next, nil, carry)
		}
	}
	return l
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{Val: 3}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{Val: 4}}}
	l := addTwoNumbers(l1, l2)
	fmt.Println(l)
	fmt.Println(l.Next)
	fmt.Println(l.Next.Next)
	fmt.Println(l.Next.Next.Next)
}
