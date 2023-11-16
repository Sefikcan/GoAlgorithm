package main

import "fmt"

// Usage:
// Airport boarding system
// Job scheduling
// Railway ticketing system

// define the QNode
type queueNode struct {
	val  int
	next *queueNode
}

// define queue
type myQueue struct {
	front *queueNode
	rear  *queueNode
}

// define constructor
func QueueConstructor() myQueue {
	return myQueue{
		front: nil,
		rear:  nil,
	}
}

// enqueue: will be to rear end of the Queue
func (this *myQueue) enQueue(v int) {
	// create a new node of type queueMode
	newNode := &queueNode{
		val:  v,
		next: nil,
	}

	// if Queue is empty: make front/rear as new node
	if this.front == nil {
		this.front = newNode
		this.rear = newNode
	} else {
		this.rear.next = newNode
		this.rear = newNode
	}
}

// dequeue: remove node/element from front of the Queue
func (this *myQueue) deQueue() int {
	if this.isEmpty() {
		return -1
	}

	val := this.front.val
	//remove the node from the front
	this.front = this.front.next
	if this.front == nil {
		this.rear = nil
	}

	return val
}

func (this *myQueue) listQueue() {
	curr := this.front

	// traverse over the queue
	for curr != nil {
		fmt.Printf("..%d..", curr.val)
		curr = curr.next
	}
}

func (this *myQueue) Front() int {
	if this.isEmpty() {
		return -1
	}
	return this.front.val
}

// check if Queue is empty
func (this *myQueue) isEmpty() bool {
	if this.front == nil {
		return true
	}

	return false
}

func main() {
	q := QueueConstructor()

	q.enQueue(1)
	q.enQueue(2)
	q.enQueue(3)
	q.enQueue(4)
	q.enQueue(5)

	q.listQueue()

	front := q.Front()
	fmt.Printf("\n front of the queue is %d \n", front) // front is 1

	front = q.deQueue()
	fmt.Printf("\n front of the queue is %d \n", front) // front is 1

	front = q.Front()
	fmt.Printf("\n front of the queue is %d \n", front) // front is 2
}
