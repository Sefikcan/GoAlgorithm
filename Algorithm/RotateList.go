package main

// compute length of listNode
// adjust k = k % len(listNode)
// detect new Tail; tail = len(listNode) - k
// identify new head
// set new head and tail

// perform rotations by "k" times from right of list
func rotateRight(head *ListNode, k int) *ListNode {
	// sanity checks
	if head == nil || head.Next == nil {
		// nothing to rotate
		return head
	}

	// compute the length of the linked list
	length := 1
	currentNode := head

	for currentNode.Next != nil {
		currentNode = currentNode.Next
		length++
	}

	// adjust value of "k"
	k = k % length
	if k == 0 {
		return head
	}

	// detect the new tail and new head for the rotating the list
	newTail := head
	for i := 1; i < length-k; i++ {
		newTail = newTail.Next
	}

	// detect new head
	newHead := newTail.Next

	//adjust the new tail/tail and break the original connection
	currentNode.Next = head // connect the original tail to original head
	newTail.Next = nil      // set the new tail's next pointer to nil
	head = newHead          // update the head of the rotated list

	return head
}
