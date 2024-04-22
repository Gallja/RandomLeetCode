package main

import (
	"fmt"
)

type node struct {
	next  *node
	value string
}

type queue struct {
	head *node
	last *node
}

type graph struct {
}

func (q *queue) enqueue(value string) {
	newNode := &node{nil, value}

	if q.head == nil && q.last == nil {
		q.head = newNode
		q.last = newNode
		return
	}

	q.last.next = newNode
	q.last = newNode
}

func (q *queue) dequeue() *node {
	if q.head == nil && q.last == nil {
		return nil
	}

	node := q.head
	if q.head == q.last {
		q.head = nil
		q.last = nil
	}

	q.head = node.next
	return node
}

func (q queue) isEmpty() bool {
	return q.head == nil
}

func main() {
	var queue queue
	slice1 := "0 0 0 0"
	slice2 := "1 0 1 1"

	fmt.Println(queue.isEmpty())
	fmt.Println()

	queue.enqueue(slice1)
	queue.enqueue(slice2)
	fmt.Println(queue.isEmpty())
}
