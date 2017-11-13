package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	p := head

	// first iteration, calc the length
	for p != nil {
		length++
		p = p.Next
	}

	idx := length - n // the index of the nth node from the end of list
	if idx == 0 {
		return head.Next
	}
	p = head
	for i := 0; i < idx-1; i++ {
		p = p.Next
	}
	p.Next = p.Next.Next
	return head
}

func main() {
	l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{Val: 5}}}}}
	removeNthFromEnd(l, 2)
	l = &ListNode{1, &ListNode{Val: 2}}
	removeNthFromEnd(l, 1)
}
