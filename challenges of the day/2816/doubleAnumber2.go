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
	listRet := LinkedList{nil, 0}

	if head.Val >= 5 {

		return listRet.head
	}

	curr := head
	curr2 := head.Next

	listRet.add(curr.Val * 2)
	pointerPrec := listRet.getByIndex(0)
	index := 1

	for curr2 != nil {
		if curr2.Val*2 >= 10 {
			firstNum := (curr2.Val * 2) % 10
			secondNum := (curr2.Val * 2) / 10

			// fmt.Println("primo:", firstNum)
			// fmt.Println("secondo:", secondNum)
			// modifica della linked list in base ai risultati di firstNum e secondNum:
			listRet.add(firstNum)
			index--
			pointerPrec = listRet.getByIndex(index)
			pointerPrec.Val += secondNum
			index += 2
		} else {
			listRet.add(curr2.Val * 2)
			pointerPrec = listRet.getByIndex(index)
			index++
		}

		curr2 = curr2.Next
	}

	listRet.printList()

	return listRet.head
}

func (list LinkedList) getByValue(value int) *ListNode {
	current := list.head

	for current.Next != nil {
		if current.Val == value {
			return current
		}

		current = current.Next
	}

	return nil
}

func (list LinkedList) getByIndex(index int) *ListNode {
	current := list.head

	for current != nil && index > 0 {
		current = current.Next
		index--
	}

	return current
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

	list.add(4)
	list.add(5)
	list.add(7)

	list.printList()

	doubleIt(list.head)

	list2 := LinkedList{nil, 0}

	list2.add(5)
	list2.add(4)
	list2.add(2)

	doubleIt(list2.head)
}
