package main

import (
	"fmt"
)

type queueNode struct {
	next  *queueNode
	value string
}

type queue struct {
	head *queueNode
}

type graph struct {
	map_ map[string][]string
}

func (q *queue) enqueue(value string) {
	if q.head == nil {
		q.head = &queueNode{nil, value}
		return
	}

	node := q.head

	for node.next != nil {
		node = node.next
	}

	newNode := queueNode{nil, value}
	node.next = &newNode
}

func (q *queue) dequeue() string {
	head := q.head
	q.head = q.head.next

	return head.value
}

func (q queue) isEmpty() bool {
	return q.head == nil
}

func newGraph() graph {
	return graph{make(map[string][]string)}
}

func bfs(g graph, start string) map[string]int {
	queue := queue{nil}
	distances := make(map[string]int)
	distances[start] = 0

	queue.enqueue(start)

	for !queue.isEmpty() {
		u := queue.dequeue()

		for _, v := range g.map_[u] {
			if _, reached := distances[v]; !reached {
				distances[v] = distances[u] + 1
				queue.enqueue(v)
			}
		}
	}

	return distances
}

func main() {
	fmt.Println()
}
