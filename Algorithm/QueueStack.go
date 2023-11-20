package main

import "fmt"

// define stack node
type StackNode struct {
	value int
	next  *StackNode
}

// define stack
type Stack struct {
	head *StackNode
}

func (s *Stack) Push(x int) {
	newNode := &StackNode{value: x, next: s.head}

	s.head = newNode
}

func (s *Stack) Pop() int {
	if s.isEmpty() {
		return -1
	}

	val := s.head.value
	s.head = s.head.next

	return val
}

func (s *Stack) Peek() int {
	if s.isEmpty() {
		return -1
	}

	return s.head.value
}

func (s *Stack) isEmpty() bool {
	return s.head == nil
}

type MyQueue struct {
	stack1 *Stack
	stack2 *Stack
}

func Constructors() MyQueue {
	return MyQueue{
		stack1: &Stack{},
		stack2: &Stack{},
	}
}

func (this *MyQueue) Empty() bool {
	return this.stack1.isEmpty() && this.stack2.isEmpty()
}

func (this *MyQueue) Pop() int {
	if this.stack1.isEmpty() && this.stack2.isEmpty() {
		return -1
	}

	if this.stack2.isEmpty() {
		for !this.stack1.isEmpty() {
			this.stack2.Push(this.stack1.Pop())
		}
	}

	return this.stack2.Pop()
}

func (this *MyQueue) Peek() int {
	if this.stack1.isEmpty() && this.stack2.isEmpty() {
		return -1
	}

	if this.stack2.isEmpty() {
		for !this.stack1.isEmpty() {
			this.stack2.Push(this.stack1.Pop())
		}
	}

	return this.stack2.Peek()
}

func (this *MyQueue) Push(x int) {
	this.stack1.Push(x)
}

func main() {
	myQueue := Constructors()
	myQueue.Push(1)
	myQueue.Push(2)
	fmt.Println(myQueue.Peek())  // Output: 1
	fmt.Println(myQueue.Pop())   // Output: 1
	fmt.Println(myQueue.Empty()) // Output: false
}
