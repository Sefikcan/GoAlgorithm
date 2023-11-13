package main

type listNode struct {
	Val  int
	Next *listNode
}

func deleteMiddle(head *listNode) *listNode {
	// Empty or only one node in the list check
	if head == nil || head.Next == nil {
		return nil
	}

	// to fetch the middle node, set two pointers running at different pace
	slow, fast := head, head.Next.Next

	// traverse over the linked list
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// skip the slow.Next
	slow.Next = slow.Next.Next

	return head
}
