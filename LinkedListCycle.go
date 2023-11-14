package main

import "fmt"

// We use floyd's algorithm
// First we set slow and fast pointer to head node
// Then moves slowly 1 step, fast moves 2 steps with each step
// if there is no loop, the fast pointer will terminate at some point and the algorithm terminates.
// if there is loop, fast and slow meet same node

func hasCycle(head *ListNode) bool {
	// sanity checks for existance of list
	if head == nil {
		return false
	}

	// by Floyd's algorithm, we have two pointers
	// fast and slow that traverse with different pace
	slow, fast := head, head

	// traverse until end of the list
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}

	return false
}

func main() {
	head1 := &ListNode{Val: 3}
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 0}
	node3 := &ListNode{Val: -4}

	head1.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node1 // Cycle

	result1 := hasCycle(head1)
	fmt.Println("Example 1:", result1)
}
