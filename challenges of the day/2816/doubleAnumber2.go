package main

import (
	"fmt"
)

type LinkedList struct {
	head *ListNode
	size int
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *LinkedList) add(Val int) {
	newNode := &ListNode{Val, nil}

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

func (list *LinkedList) printList() {
	if list.size == 0 {
		fmt.Println("[ ]")
		return
	}

	fmt.Print("[ ", list.head.Val)

	current := list.head.Next

	for current != nil {
		fmt.Print(", ", current.Val)
		current = current.Next
	}

	fmt.Print(" ]")
	fmt.Println()
}

func doubleIt(head *ListNode) *ListNode {

	return nil
}

func intToSlice(n int) []int {
	var result []int

	for n > 0 {
		digit := n % 10
		result = append([]int{digit}, result...)
		n /= 10
	}

	return result
}

func main() {
	list := LinkedList{nil, 0}

	list.add(9)
	list.add(1)
	list.add(9)
	list.add(5)

	list.printList()

	fmt.Println(doubleIt(list.head))
}
