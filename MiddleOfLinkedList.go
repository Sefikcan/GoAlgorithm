package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Find middle of node
// We use two node, fast and slow
// Fast taking 2 steps and slow taking 1 step
// When faster reaches the end of the list, the slow will be at the middle node
func middleNode(head *ListNode) *ListNode {
	// two pointers marked at head
	slow, fast := head, head

	// perform traversal until end of the list
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result := middleNode(head)
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
}
