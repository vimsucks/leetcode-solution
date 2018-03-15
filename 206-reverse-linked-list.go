package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := &ListNode{1, &ListNode{2, &ListNode{Val: 3}}}
	nl := reverseList(l)
	fmt.Println(nl.Val)
	fmt.Println(nl.Next.Val)
	fmt.Println(nl.Next.Next.Val)

	l = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{Val: 4}}}}
	nl = reverseList(l)
	fmt.Println(nl.Val)
	fmt.Println(nl.Next.Val)
	fmt.Println(nl.Next.Next.Val)
	fmt.Println(nl.Next.Next.Next.Val)

	l = &ListNode{1, &ListNode{Val: 2}}
	nl = reverseList(l)
	fmt.Println(nl.Val)
	fmt.Println(nl.Next.Val)

	l = &ListNode{Val: 1}
	nl = reverseList(l)
	fmt.Println(nl.Val)

	l = &ListNode{}
	nl = reverseList(l)
	fmt.Println(nl.Val)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p1, p2 := head, head.Next
	var lastP2 *ListNode = nil
	for p2 != nil && p2.Next != nil {
		c := p2.Next
		p1.Next = lastP2
		p2.Next = p1
		lastP2 = p2
		p1 = c
		p2 = c.Next
	}

	if p2 == nil {
		p1.Next = lastP2
		return p1
	} else {
		p1.Next = lastP2
		p2.Next = p1
		return p2
	}
}
