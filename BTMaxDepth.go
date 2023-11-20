package main

import (
	"fmt"
	"math"
)

type maxDepthTreeNode struct {
	Val   int
	Left  *maxDepthTreeNode
	Right *maxDepthTreeNode
}

func maxDepth(root *maxDepthTreeNode) int {
	if root == nil {
		return 0
	}

	// recursively calculate max dept of left and right tree
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	// return max depth
	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
}

func main() {
	root1 := &maxDepthTreeNode{
		Val: 3,
		Left: &maxDepthTreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &maxDepthTreeNode{
			Val: 20,
			Left: &maxDepthTreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &maxDepthTreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	// Example 2: [1,null,2]
	root2 := &maxDepthTreeNode{
		Val:  1,
		Left: nil,
		Right: &maxDepthTreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	// Calculate and print the maximum depth for each example
	fmt.Println("Example 1 Output:", maxDepth(root1)) // Output: 3
	fmt.Println("Example 2 Output:", maxDepth(root2)) // Output: 2
}
