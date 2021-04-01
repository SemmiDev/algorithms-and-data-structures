package main

import "fmt"

type Queue []interface{}

type QueueOperations interface {
	IsEmpty() bool
	size() int
	enqueue(value interface{})
	dequeue()
}

func NewQueue() QueueOperations {
	return &Queue{}
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) size() int {
	return len(*q)
}

func (q *Queue) enqueue(value interface{}) {
	*q = append(*q, value)
}

func (q *Queue) dequeue() {
	*q = (*q)[1:]
}

func main() {
	queue := NewQueue()

	queue.enqueue(1)
	queue.enqueue(2)

	fmt.Print(queue)
	queue.dequeue()
	fmt.Print(queue)

	queue.enqueue(3)
	queue.enqueue(4)
	queue.enqueue(5)

	fmt.Print(queue)
	queue.dequeue()
	fmt.Print(queue)
}
