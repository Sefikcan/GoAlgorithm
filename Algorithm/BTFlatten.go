package main

import "fmt"

type flattenTreeNode struct {
	Val   int
	Left  *flattenTreeNode
	Right *flattenTreeNode
}

func flatten(root *flattenTreeNode) {
	// sanity checks
	if root == nil {
		return
	}

	// traverse the left subtree
	flatten(root.Left)

	// traverse the right subtree
	flatten(root.Right)

	// save the right subtree
	rTree := root.Right

	// make the left subtree as new right subtree
	root.Right = root.Left

	// clear the left subtree
	root.Left = nil

	// traverse until end of the new right subtree
	curr := root
	for curr.Right != nil {
		curr = curr.Right
	}

	// attach the saved right tree now
	curr.Right = rTree
}

func main() {
	root := &flattenTreeNode{Val: 1, Left: &flattenTreeNode{Val: 2, Left: &flattenTreeNode{Val: 3}, Right: &flattenTreeNode{Val: 4}}, Right: &flattenTreeNode{Val: 5, Right: &flattenTreeNode{Val: 6}}}
	flatten(root)
	current := root
	for current != nil {
		fmt.Print(current.Val)
		if current.Right != nil {
			fmt.Print(" -> ")
		}
		current = current.Right
	}
	fmt.Println()
}
