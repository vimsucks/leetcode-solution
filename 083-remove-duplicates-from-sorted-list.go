package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{Val: 3}}}}}
	deleteDuplicates(list)
	fmt.Println(list.Val)
	fmt.Println(list.Next.Val)
	fmt.Println(list.Next.Next.Val)
	fmt.Println(list.Next.Next.Next.Val)
}

func deleteDuplicates(head *ListNode) *ListNode {
	ptr := head
	for ptr != nil {
		for ptr.Next != nil && ptr.Val == ptr.Next.Val {
			ptr.Next = ptr.Next.Next
		}
		ptr = ptr.Next
	}
	return head
}
