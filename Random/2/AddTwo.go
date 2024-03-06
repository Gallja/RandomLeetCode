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

func printListFromHead(head *ListNode) {
	if head == nil {
		fmt.Println("[ ]\nEmpty List...")
		return
	}

	fmt.Print("[ ", head.Val)

	curr := head.Next

	for curr != nil {
		fmt.Print(", ", curr.Val)
		curr = curr.Next
	}

	fmt.Println(" ]")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	curr := l1
	curr2 := l2

	listRis := linkedList{nil, 0}
	riporto := false
	rest := 0

	for curr != nil {
		second := 0

		if curr2 != nil {
			second = curr2.Val
		}

		if curr.Val+second > 9 {
			rest = curr.Val / second
			second = (curr.Val + second) % 10
			riporto = true
		}

		if riporto {
			listRis.add(second)
			riporto = false
		} else {
			listRis.add(curr.Val + second + rest)
			rest = 0
		}

		curr = curr.Next
		curr2 = curr2.Next
	}

	return listRis.head
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

	printListFromHead(addTwoNumbers(list1.head, list2.head))
}
