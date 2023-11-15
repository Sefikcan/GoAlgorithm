package main

import "fmt"

// define stack node
type stackNode struct {
	value int
	next  *stackNode
}

// define stack
type stack struct {
	head *stackNode
}

// stack operations : pop(), push() and peek()

// push(): to enter a value to top of the stack / head of the list
func (s *stack) push(x int) {
	// if stack is empty, make new node as head (top of the stack)
	newNode := &stackNode{value: x, next: s.head}

	// otherwise,
	s.head = newNode
}

// pop(): ejecting top of the stack / pulling from head of the list
func (s *stack) pop() int {
	if s.isEmpty() {
		return -1
	}
	val := s.head.value
	// assign new head for the stacl
	s.head = s.head.next
	return val
}

func (s *stack) isEmpty() bool {
	return s.head == nil
}

// peek(): take a look or view top of the stack
func (s *stack) peek() int {
	if s.isEmpty() {
		return -1
	}

	return s.head.value
}

// instantiate the stack object
func constructor() stack {
	return stack{head: nil}
}

func main() {
	fmt.Println("calling main..")

	st := constructor()

	// stack ops
	st.push(1)
	st.push(2)
	st.push(3)
	st.push(4)
	st.push(5)

	top := st.peek()

	fmt.Printf("top of the stack is now %d\n", top)

	top = st.pop()
	fmt.Printf("top of the stack is now %d\n", top)

	top = st.peek()
	fmt.Printf("top of the stack is now %d\n", top)
}
