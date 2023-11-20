package main

import "fmt"

type treeNodePostOrder struct {
	Val   int
	Left  *treeNodePostOrder
	Right *treeNodePostOrder
}

func postOrderTraversal(root *treeNodePostOrder) []int {
	res := []int{}

	if root != nil {
		// traverse the left subtree
		res = append(res, postOrderTraversal(root.Left)...)

		// traverse the right subtree
		res = append(res, postOrderTraversal(root.Right)...)

		// visit the root
		res = append(res, root.Val)
	}

	return res
}

func main() {
	root1 := &treeNodePostOrder{Val: 1, Right: &treeNodePostOrder{Val: 2, Left: &treeNodePostOrder{Val: 3}}}
	fmt.Println(postOrderTraversal(root1)) // Output: [3 2 1]
}
