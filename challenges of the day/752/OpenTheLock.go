package main

import (
	"fmt"
)

type node struct {
	next  *node
	value []int
}

type queue struct {
	head *node
	last *node
}

func (q *queue) enqueue(value []int) {
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
	fmt.Println("prova")

}
