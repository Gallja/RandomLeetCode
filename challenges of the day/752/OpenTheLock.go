package main

import (
	"fmt"
	"strconv"
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

// bfs with graph
func bfs(deadends []string, g graph, target string) (map[string]int, int) {
	start := "0000"

	queue := queue{nil}
	distances := make(map[string]int)
	distances[start] = 0

	queue.enqueue(start)
	val := 0

	for !queue.isEmpty() {
		u := queue.dequeue()

		for _, v := range g.map_[u] {
			if _, reached := distances[v]; !reached {
				if !contains(deadends, v) {
					distances[v] = distances[u] + 1

					if v == target {
						val = distances[v]
					}

					queue.enqueue(v)
				}
			}
		}

	}

	return distances, val
}

func contains(slice []string, val string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == val {
			return true
		}
	}

	return false
}

// bfs with abstract graph
func openLock(deadends []string, target string) int {
	start := "0000"

	queue := queue{nil}
	distances := make(map[string]int)
	distances[start] = 0

	queue.enqueue(start)
	val := 0

	for !queue.isEmpty() {
		u := queue.dequeue()

		adjacent := searchAdjacent2(u)

		for i := 0; i < len(adjacent); i++ {
			if _, reached := distances[adjacent[i]]; !reached {
				if !contains(deadends, adjacent[i]) {
					distances[adjacent[i]] = distances[u] + 1

					if adjacent[i] == target {
						val = distances[adjacent[i]]
					}

					queue.enqueue(adjacent[i])
				}
			}
		}

	}

	return val
}

/*
// search adjacent brute force
func searchAdjacent(start string) []string {
	slice := stringToIntArray(start)
	ris := []string{}

	first1 := slice[0]
	tmp := slice[0]
	first1 = checkRis(first1 + 1)
	slice[0] = first1
	str := intToStringArray(slice)
	slice[0] = tmp
	ris = append(ris, str)

	first2 := slice[0]
	first2 = checkRis(first2 - 1)
	slice[0] = first2
	str = intToStringArray(slice)
	slice[0] = tmp
	ris = append(ris, str)

	sec1 := slice[1]
	tmp = slice[1]
	sec1 = checkRis(sec1 + 1)
	slice[1] = sec1
	str = intToStringArray(slice)
	slice[1] = tmp
	ris = append(ris, str)

	sec2 := slice[1]
	sec2 = checkRis(sec2 - 1)
	slice[1] = sec2
	str = intToStringArray(slice)
	slice[1] = tmp
	ris = append(ris, str)

	third1 := slice[2]
	tmp = slice[2]
	third1 = checkRis(third1 + 1)
	slice[2] = third1
	str = intToStringArray(slice)
	slice[2] = tmp
	ris = append(ris, str)

	third2 := slice[2]
	third2 = checkRis(third2 - 1)
	slice[2] = third2
	str = intToStringArray(slice)
	slice[2] = tmp
	ris = append(ris, str)

	fourth1 := slice[3]
	tmp = slice[3]
	fourth1 = checkRis(fourth1 + 1)
	slice[3] = fourth1
	str = intToStringArray(slice)
	slice[3] = tmp
	ris = append(ris, str)

	fourth2 := slice[3]
	fourth2 = checkRis(fourth2 - 1)
	slice[3] = fourth2
	str = intToStringArray(slice)
	slice[3] = tmp
	ris = append(ris, str)

	return ris
}
	*/

// search adjacent with iteration
func searchAdjacent2(start string) []string {
	slice := stringToIntArray(start)
	ris := []string{}

	for i := 0; i < len(slice); i++ {
		tmp := slice[i]
		for j := -1; j <= 1; j += 2 {
			slice[i] = checkRis(tmp + j)
			str := intToStringArray(slice)
			ris = append(ris, str)
			slice[i] = tmp
		}
	}

	return ris
}

func stringToIntArray(s string) []int {
	var result []int
	for _, c := range s {
		if num, err := strconv.Atoi(string(c)); err == nil {
			result = append(result, num)
		}
	}
	return result
}

func intToStringArray(slice []int) string {
	ris := ""

	for i := 0; i < len(slice); i++ {
		ris += strconv.Itoa(slice[i])
	}

	return ris
}

func checkRis(num int) int {
	if num == -1 {
		return 9
	}

	if num == 10 {
		return 0
	}

	return num
}

func main() {
	graph := newGraph()

	graph.map_["0000"] = []string{"1000", "0100"}
	graph.map_["1000"] = []string{"2000", "1100"}
	graph.map_["0100"] = []string{"0200", "0110"}
	graph.map_["0200"] = []string{"0300", "0000"}
	graph.map_["0110"] = []string{"0000", "1000"}

	deadends := []string{"0201", "0101", "0102", "1212", "2002"}

	fmt.Println(bfs(deadends, graph, "0300"))
	fmt.Println(openLock(deadends, "0202"))
}
