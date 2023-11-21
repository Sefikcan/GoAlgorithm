package main

import "fmt"

type DuplicateListNode struct {
	Val  int
	Next *DuplicateListNode
}

func deleteDuplicates(head *DuplicateListNode) *DuplicateListNode {
	current := head

	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

func main() {
	head1 := &DuplicateListNode{Val: 1, Next: &DuplicateListNode{Val: 1, Next: &DuplicateListNode{Val: 2}}}
	head1 = deleteDuplicates(head1)
	for head1 != nil {
		fmt.Printf("%d", head1.Val)
		head1 = head1.Next
		fmt.Println()
	}
}
