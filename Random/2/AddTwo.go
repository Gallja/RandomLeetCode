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
	head := new(ListNode) // head "dummy" (fittizia)
	curr := head
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next
		carry = sum / 10
	}

	return head.Next
}

func main() {
	list1 := linkedList{nil, 0}

	list1.add(9)
	list1.add(9)
	list1.add(9)
	list1.add(9)
	list1.add(9)
	list1.add(9)
	list1.add(9)

	list2 := linkedList{nil, 0}

	list2.add(9)
	list2.add(9)
	list2.add(9)
	list2.add(9)

	list1.printList()
	list2.printList()

	printListFromHead(addTwoNumbers(list1.head, list2.head))
}
