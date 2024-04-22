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

func bfs(g graph, target string) (map[string]int, int) {
	start := "0 0 0 0"

	queue := queue{nil}
	distances := make(map[string]int)
	distances[start] = 0

	queue.enqueue(start)
	val := 0

	for !queue.isEmpty() {
		u := queue.dequeue()

		for _, v := range g.map_[u] {
			if _, reached := distances[v]; !reached {
				distances[v] = distances[u] + 1

				if v == target {
					val = distances[v]
				}

				queue.enqueue(v)
			}
		}

	}

	return distances, val
}

func main() {
	graph := newGraph()

	graph.map_["0 0 0 0"] = []string{"1 0 0 0", "0 1 0 0"}
	graph.map_["1 0 0 0"] = []string{"2 0 0 0", "1 1 0 0"}
	graph.map_["0 1 0 0"] = []string{"0 2 0 0", "0 1 1 0"}

	fmt.Println(bfs(graph, "0 2 0 0"))
}
