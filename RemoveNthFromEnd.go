package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// create two selectors
	selector1, selector2 := head, head

	// walk first selector "n" steps
	for i := 0; i < n; i++ {
		selector1 = selector1.Next
	}

	// at this point if selector1 is nil, it means we need to remove the head node
	if selector1 == nil {
		return head.Next
	}

	// traverse selector1 until end of the list, while also traverse selector2 at same
	for selector1.Next != nil {
		selector1 = selector1.Next
		selector2 = selector2.Next
	}

	// at this point selector2 is at one step behind the node that needs to be removed
	selector2.Next = selector2.Next.Next

	return head
}
