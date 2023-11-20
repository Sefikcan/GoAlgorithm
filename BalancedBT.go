package main

import (
	"fmt"
	"math"
)

// The difference between the depth of both subtrees should not be greater than 1
type balancedTreeNode struct {
	Val   int
	Left  *balancedTreeNode
	Right *balancedTreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isBalanced(root *balancedTreeNode) bool {
	// sanity check
	if root == nil {
		return true
	}

	// calculate heights of left and right subtrees
	leftDepth := calcDepth(root.Left)
	rightDepth := calcDepth(root.Right)

	if math.Abs(float64(leftDepth-rightDepth)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func calcDepth(root *balancedTreeNode) int {
	// sanity check
	if root == nil {
		return 0
	}

	// get leftDepth and rightDepth
	leftDepth := calcDepth(root.Left)
	rightDepth := calcDepth(root.Right)

	return max(leftDepth, rightDepth) + 1
}

func main() {
	root1 := &balancedTreeNode{Val: 3, Left: &balancedTreeNode{Val: 9}, Right: &balancedTreeNode{Val: 20, Left: &balancedTreeNode{Val: 15}, Right: &balancedTreeNode{Val: 7}}}
	fmt.Println("Result:", isBalanced(root1)) // true
}
