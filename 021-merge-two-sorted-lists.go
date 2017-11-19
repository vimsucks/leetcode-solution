package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p := &ListNode{}
	pCopy := p
	for l1 != nil || l2 != nil {
		if l1 == nil {
			p.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		} else if l2 == nil {
			p.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			if l1.Val <= l2.Val {
				p.Next = &ListNode{Val: l1.Val}
				l1 = l1.Next
			} else {
				p.Next = &ListNode{Val: l2.Val}
				l2 = l2.Next
			}
		}
		p = p.Next
		println(p.Val)
	}
	return pCopy.Next
}

func main() {
	//l1 := &ListNode{1, &ListNode{4, &ListNode{7, &ListNode{Val: 10}}}}
	l2 := &ListNode{2, &ListNode{3, &ListNode{8, &ListNode{Val: 11}}}}
	mergeTwoLists(nil, l2)
}
