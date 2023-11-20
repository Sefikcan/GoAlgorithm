package main

import "fmt"

type SameTreeNode struct {
	Val   int
	Left  *SameTreeNode
	Right *SameTreeNode
}

func isSameTree(p *SameTreeNode, q *SameTreeNode) bool {
	// check if both trees are nil
	if p == nil && q == nil {
		return true
	}

	// check if only one of them is nil
	if p == nil || q == nil {
		return false
	}

	// verify the values
	if p.Val != q.Val {
		return false
	}

	// recursive and check if left and right subtrees are same
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	p1 := &SameTreeNode{Val: 1, Left: &SameTreeNode{Val: 2}, Right: &SameTreeNode{Val: 3}}
	q1 := &SameTreeNode{Val: 1, Left: &SameTreeNode{Val: 2}, Right: &SameTreeNode{Val: 3}}
	fmt.Println(isSameTree(p1, q1)) // Output: true
}
