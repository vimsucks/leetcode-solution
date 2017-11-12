package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	p := head
	for p != nil {
		length++
		p = p.Next
	}
	idx := length - n // the index of the nth node from the end of list
	if idx == 0 {
		return head.Next
	}
	p = head
	for ; idx != 1; idx-- {
		p = p.Next
	}
	p.Next = p.Next.Next
	return head
}

func main() {
	l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{Val: 5}}}}}
	removeNthFromEnd(l, 2)
	fmt.Println(l.Val)
	fmt.Println(l.Next.Val)
	fmt.Println(l.Next.Next.Val)
	fmt.Println(l.Next.Next.Next.Val)
	fmt.Println(l.Next.Next.Next.Next)
	l = &ListNode{1, &ListNode{Val: 2}}
	removeNthFromEnd(l, 1)
}
