package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t1 := &TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}
	t2 := &TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}
	fmt.Println(isSameTree(t1, t2))
	t3 := &TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 1}}
	t4 := &TreeNode{1, &TreeNode{Val: 1}, &TreeNode{Val: 2}}
	fmt.Println(isSameTree(t3, t4))
	t5 := &TreeNode{1, &TreeNode{Val: 2}, nil}
	t6 := &TreeNode{1, nil, &TreeNode{Val: 2}}
	fmt.Println(isSameTree(t5, t6))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameTreeIter(p, q)
}

func isSameTreeIter(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	if isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) {
		return true
	} else {
		return false
	}
}
