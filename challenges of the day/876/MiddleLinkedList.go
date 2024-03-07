package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type linkedList struct {
	head *ListNode
	size int
}

func (list *linkedList) getSize() int {
	return list.size
}

func getSizeFromHead(head *ListNode) int {
	if head == nil {
		return 0
	}

	cnt := 1
	curr := head.Next

	for curr != nil {
		cnt++
		curr = curr.Next
	}

	return cnt
}

func (list *linkedList) add(value int) {
	newNode := &ListNode{value, nil}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.Next != nil {
			current = current.Next
		}

		current.Next = newNode
	}

	list.size++
}

func (list *linkedList) printList() {
	if list.size == 0 {
		fmt.Println("[ ]\nEmpty List...")
		return
	}

	fmt.Print("[ ", list.head.Val)

	current := list.head.Next

	for current != nil {
		fmt.Print(", ", current.Val)
		current = current.Next
	}

	fmt.Println(" ]")
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	size := getSizeFromHead(head)
	middle := 0

	if size%2 == 0 {
		middle = size / 2
	} else {
		middle = (size / 2) + 1
	}

	cnt := 1
	curr := head

	for cnt < middle {
		curr = curr.Next
		cnt++
	}

	return curr
}

func main() {
	list := linkedList{nil, 0}

	list.add(1)
	list.add(2)
	list.add(3)
	list.add(4)
	list.add(5)

	list.printList()

	fmt.Println("Dimensione:", getSizeFromHead(list.head))

	fmt.Println(middleNode(list.head))
}
