package main

import "fmt"

type MinStack struct {
	stack    []int
	minStack []int
}

func MinStackConstructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}

	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	if top == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -1
	}

	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) == 0 {
		return -1
	}

	return this.minStack[len(this.minStack)-1]
}

func main() {
	obj := MinStackConstructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println("Top:", obj.Top())    // Output: -3
	fmt.Println("Min:", obj.GetMin()) // Output: -3
	obj.Pop()
	fmt.Println("Top:", obj.Top())    // Output: 0
	fmt.Println("Min:", obj.GetMin()) // Output: -2
}
