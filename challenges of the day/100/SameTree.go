package main

import (
	"fmt"
)

type node struct {
	next  *node
	value int
}

type linkedList struct {
	head *node
	size int
}

// Given the roots of two binary trees p and q, write a function to check if they are the same or not.
func isSameTree(p *node, q *node) bool {
	if getSizeByRoot(p) != getSizeByRoot(q) {
		return false
	}

	current := p
	current2 := q

	for current != nil && current2 != nil {
		if current.value != current2.value {
			return false
		}

		current = current.next
		current2 = current2.next
	}

	return true
}

func getSizeByRoot(root *node) int {
	size := 0
	current := root

	for current != nil {
		size++
		current = current.next
	}

	return size
}

func (list *linkedList) addValue(value int) {
	newNode := &node{nil, value}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}

	list.size++
}

func (list linkedList) getByIndex(index int) *node {
	current := list.head

	for current.next != nil && index > 0 {
		current = current.next
		index--
	}

	return current
}

func (list linkedList) printList() {
	if list.size == 0 {
		fmt.Println("[ ]\nEmpty list.")
		return
	}

	fmt.Print("[ ", list.head.value)

	current := list.head.next

	for current != nil {
		fmt.Print(", ", current.value)
		current = current.next
	}

	fmt.Print(" ]")
	fmt.Println()
}

func main() {
	tree1 := linkedList{nil, 0}
	tree2 := linkedList{nil, 0}

	tree1.addValue(1)
	tree1.addValue(2)
	tree1.addValue(3)

	tree2.addValue(1)
	tree2.addValue(2)
	tree2.addValue(3)
	// tree2.addValue(4)

	tree1.printList()
	tree2.printList()

	fmt.Println(isSameTree(tree1.head, tree2.head))
}
