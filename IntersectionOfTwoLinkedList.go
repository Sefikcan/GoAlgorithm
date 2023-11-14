package main

import "fmt"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// compute the length of headA and headB
	lenA, lenB := 0, 0
	currentA, currentB := headA, headB
	for currentA != nil {
		lenA++
		currentA = currentA.Next
	}

	for currentB != nil {
		lenB++
		currentB = currentB.Next
	}

	// equalize starting point
	currentA, currentB = headA, headB
	for lenA > lenB {
		currentA = currentA.Next
		lenA--
	}

	for lenB > lenA {
		currentB = currentB.Next
		lenB--
	}

	// check intersection point
	for currentA != currentB {
		currentA = currentA.Next
		currentB = currentB.Next
	}

	return currentA
}

func main() {
	listA1 := &ListNode{Val: 4}
	listA1.Next = &ListNode{Val: 1}
	listA1.Next.Next = &ListNode{Val: 8}
	listA1.Next.Next.Next = &ListNode{Val: 4}
	listA1.Next.Next.Next.Next = &ListNode{Val: 5}

	listB1 := &ListNode{Val: 5}
	listB1.Next = &ListNode{Val: 6}
	listB1.Next.Next = &ListNode{Val: 1}
	listB1.Next.Next.Next = &ListNode{Val: 8}
	listB1.Next.Next.Next.Next = &ListNode{Val: 4}
	listB1.Next.Next.Next.Next.Next = &ListNode{Val: 5}

	listA1.Next.Next.Next.Next = listB1.Next.Next.Next

	result1 := getIntersectionNode(listA1, listB1)
	fmt.Println("Example 1:", result1.Val)
}
