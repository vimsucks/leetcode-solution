package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t := &TreeNode{3, &TreeNode{Val: 9}, &TreeNode{20, &TreeNode{Val: 15}, &TreeNode{Val: 7}}}
	fmt.Println(maxDepth(t))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depth int
	maxDepthIter(root, 1, &depth)
	return depth
}

func maxDepthIter(root *TreeNode, current int, depth *int) {
	if root.Left == nil && root.Right == nil {
		if current > *depth {
			*depth = current
		}
	} else {
		if root.Left != nil {
			maxDepthIter(root.Left, current+1, depth)
		}
		if root.Right != nil {
			maxDepthIter(root.Right, current+1, depth)
		}
	}
}
