package main

func deleteNode(node *ListNode) {
	// sanity check
	if node == nil || node.Next == nil {
		return
	}

	// copy & skip
	node.Val = node.Next.Val   // copy
	node.Next = node.Next.Next // skip
}
