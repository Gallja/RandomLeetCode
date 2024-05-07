package main

import (
	"fmt"
)

type LinkedList struct {
	head *ListNode
	size int
}

type ListNode struct {
	next  *ListNode
	value int
}

func (list *LinkedList) add(value int) {
	newNode := &ListNode{nil, value}

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

func (list *LinkedList) printList() {
	if list.size == 0 {
		fmt.Println("[ ]")
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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
	return nil
}

func main() {
	fmt.Println("prova")
}
