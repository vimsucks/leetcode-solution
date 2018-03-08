package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{Val: 4}}}}
	list = swapPairs(list)
	fmt.Println(list.Val)
	fmt.Println(list.Next.Val)
	fmt.Println(list.Next.Next.Val)
	fmt.Println(list.Next.Next.Next.Val)
	list = &ListNode{Val: 1}
	list = swapPairs(list)
	fmt.Println(list.Val)
	list = &ListNode{1, &ListNode{2, &ListNode{Val: 3}}}
	list = swapPairs(list)
	fmt.Println(list.Val)
	fmt.Println(list.Next.Val)
	fmt.Println(list.Next.Next.Val)
}

func swapPairs(head *ListNode) (newHead *ListNode) {
	if head == nil {
		return
	}
	if head.Next == nil {
		return head
	}
	newHead = head.Next
	var p1, p2, p3 *ListNode
	p1 = head
	for p1.Next != nil && p1.Next.Next != nil {
		p2 = p1.Next
		p3 = p1.Next.Next
		p1.Next = p3.Next
		p2.Next = p1
		p1 = p3
	}
	if p1.Next == nil {
		p2.Next.Next = p3
	} else {
		p2 = p1.Next
		p1.Next = p2.Next
		p2.Next = p1
	}
	return
}
