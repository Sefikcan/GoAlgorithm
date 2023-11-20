package main

// define of the stack
type CustomStack struct {
	stack   []int
	maxSize int
}

// instantiate the stack object
func Constructor(maxSize int) CustomStack {
	return CustomStack{
		stack:   make([]int, 0, maxSize),
		maxSize: maxSize,
	}
}

// implement the stack operation
func (this *CustomStack) Push(x int) {
	if len(this.stack) < this.maxSize {
		this.stack = append(this.stack, x)
	}
}

// pop:
func (this *CustomStack) Pop() int {
	if len(this.stack) == 0 {
		return -1
	}

	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return val
}

// increment: increment the value "val" to bottom k elements of the stack
func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < k && i < len(this.stack); i++ {
		this.stack[i] += val
	}
}
