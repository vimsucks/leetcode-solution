package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil { // both nil
		return nil
	}
	t := &TreeNode{Val: 0}
	// add up Val
	if t1 != nil {
		t.Val += t1.Val
	}
	if t2 != nil {
		t.Val += t2.Val
	}

	// either t1 or t2 is nil
	if t1 == nil || t2 == nil {
		if t1 == nil {
			t.Left = mergeTrees(nil, t2.Left)
			t.Right = mergeTrees(nil, t2.Right)
		} else if t2 == nil {
			t.Left = mergeTrees(t1.Left, nil)
			t.Right = mergeTrees(t1.Right, nil)
		}
	} else { // both t1 and t2 are not null
		t.Left = mergeTrees(t1.Left, t2.Left)
		t.Right = mergeTrees(t1.Right, t2.Right)
	}
	return t
}

func main() {
	t1 := TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 2}}
	t2 := TreeNode{Val: 2, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}}
	t := mergeTrees(&t1, &t2)
	fmt.Print(t.Val)
	fmt.Print(t.Left.Val)
	fmt.Print(t.Left.Left.Val)
	fmt.Print(t.Left.Right.Val)
	fmt.Print(t.Right.Val)
	fmt.Print(t.Right.Right.Val)
}

/*
Input:
t1 & t2:
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
Output:
t:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7
*/
