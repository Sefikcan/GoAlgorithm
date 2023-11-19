package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// This function helped to inorder navigation using the stack.
// Firstly perform to apply all left side tree, then apply right side of tree
// All operation continues until the stack is empty.
func inOrderTraversal(root *TreeNode) []int {
	result := []int{} // result array
	stack := []*TreeNode{}

	current := root // beginning node
	for current != nil || len(stack) > 0 {
		// continue until reach the bottom of the left subtree on the left side
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// remove the last added node from the stack
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// visit current node
		result = append(result, current.Val)

		// switch to right subtree
		current = current.Right
	}

	return result
}

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(inOrderTraversal(root))
}
