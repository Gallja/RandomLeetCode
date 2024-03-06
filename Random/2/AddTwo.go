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

func (list *linkedList) add(value int) {
	newNode := &ListNode{value, nil}

	if list.head == nil {
		list.head = newNode
		fmt.Println(list.head)
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}

func main() {
	list1 := linkedList{nil, 0}

	list1.add(2)
	list1.add(4)
	list1.add(3)

	list2 := linkedList{nil, 0}

	list2.add(5)
	list2.add(6)
	list2.add(4)

	list1.printList()
	list2.printList()
}
